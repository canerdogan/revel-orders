package controllers

import (
	"database/sql"
	"github.com/canerdogan/revel-orders/Godeps/_workspace/src/github.com/go-gorp/gorp"
	_ "github.com/canerdogan/revel-orders/Godeps/_workspace/src/github.com/mattn/go-sqlite3"
	"github.com/canerdogan/revel-orders/Godeps/_workspace/src/github.com/revel/modules/db/app"
	r "github.com/canerdogan/revel-orders/Godeps/_workspace/src/github.com/revel/revel"
	"github.com/canerdogan/revel-orders/app/models"
)

var (
	Dbm *gorp.DbMap
)

func InitDB() {
	db.Init()
	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.SqliteDialect{}}

	setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
		for col, size := range colSizes {
			t.ColMap(col).MaxSize = size
		}
	}

	t := Dbm.AddTable(models.User{}).SetKeys(true, "UserId")
	setColumnSizes(t, map[string]int{
		"Alias": 20,
		"Name":  100,
	})

	t = Dbm.AddTable(models.Requests{}).SetKeys(true, "RequestsId")
	t.ColMap("User").Transient = true
	// t.ColMap("RequestTime").Transient = true
	setColumnSizes(t, map[string]int{
		"Alias":        20,
		"RequestType":  50,
		"RequestCount": 50,
	})

	Dbm.TraceOn("[gorp]", r.INFO)
	Dbm.CreateTables()

	demoUser := &models.User{1, "Demo User", "demo"}
	err := Dbm.SelectOne(&demoUser, "SELECT * FROM User WHERE UserId=?", demoUser.UserId)
	if err != nil {
		if err = Dbm.Insert(demoUser); err != nil {
			panic(err)
		}
	}
}

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
