package controllers

import (
	"golang.org/x/crypto/bcrypt"
	"database/sql"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	r "github.com/revel/revel"
	"github.com/canerdogan/revel-orders/app/models"
)

var (
	Dbm *gorp.DbMap
)

func InitDB() {
    connectionString := getConnectionString()
    if db, err := sql.Open("mysql", connectionString); err != nil {
        r.ERROR.Fatal(err)
    } else {
        Dbm = &gorp.DbMap{
            Db: db,
            Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
    }
    // Defines the table for use by GORP
    // This is a function we will create soon.
	setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
		for col, size := range colSizes {
			t.ColMap(col).MaxSize = size
		}
	}

	t := Dbm.AddTable(models.Admin{}).SetKeys(true, "admin_id")
	setColumnSizes(t, map[string]int{
		"username": 20,
		"password":  100,
	})

	t = Dbm.AddTable(models.User{}).SetKeys(true, "user_id")
	setColumnSizes(t, map[string]int{
		"alias": 20,
		"name":  100,
	})

	t = Dbm.AddTable(models.Requests{}).SetKeys(true, "requests_id")
	t.ColMap("user").Transient = true
	t.ColMap("request_time").Transient = true
	// t.ColMap("RequestTime").Transient = true
	setColumnSizes(t, map[string]int{
		"alias":        20,
		"request_type":  50,
		"request_count": 50,
	})

    if err := Dbm.CreateTablesIfNotExists(); err != nil {
        r.ERROR.Fatal(err)
    }

	demoUser := &models.User{1, "Demo User", "demo"}
	err := Dbm.SelectOne(&demoUser, "SELECT * FROM User WHERE user_id=?", demoUser.UserId)
	if err != nil {
		if err = Dbm.Insert(demoUser); err != nil {
			panic(err)
		}
	}

	bcryptPassword, _ := bcrypt.GenerateFromPassword(
			[]byte("admin"), bcrypt.DefaultCost)
	adminUser := &models.Admin{1, "admin", "admin", bcryptPassword}
	err = Dbm.SelectOne(&adminUser, "SELECT * FROM Admin WHERE admin_id=?", adminUser.AdminId)
	if err != nil {
		if err = Dbm.Insert(adminUser); err != nil {
			panic(err)
		}
	}
}

// func InitDB() {
// 	db.Init()
// 	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.SqliteDialect{}}
//
// 	setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
// 		for col, size := range colSizes {
// 			t.ColMap(col).MaxSize = size
// 		}
// 	}
//
// 	t := Dbm.AddTable(models.User{}).SetKeys(true, "UserId")
// 	setColumnSizes(t, map[string]int{
// 		"Alias": 20,
// 		"Name":  100,
// 	})
//
// 	t = Dbm.AddTable(models.Requests{}).SetKeys(true, "RequestsId")
// 	t.ColMap("User").Transient = true
// 	// t.ColMap("RequestTime").Transient = true
// 	setColumnSizes(t, map[string]int{
// 		"Alias":        20,
// 		"RequestType":  50,
// 		"RequestCount": 50,
// 	})
//
// 	Dbm.TraceOn("[gorp]", r.INFO)
// 	Dbm.CreateTables()
//
// 	demoUser := &models.User{1, "Demo User", "demo"}
// 	err := Dbm.SelectOne(&demoUser, "SELECT * FROM User WHERE UserId=?", demoUser.UserId)
// 	if err != nil {
// 		if err = Dbm.Insert(demoUser); err != nil {
// 			panic(err)
// 		}
// 	}
// }

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
