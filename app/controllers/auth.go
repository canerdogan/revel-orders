package controllers

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/revel/revel"
	"github.com/canerdogan/revel-orders/app/models"
	"github.com/canerdogan/revel-orders/app/routes"
	"time"
)

const (
	DATE_FORMAT     = "Jan _2, 2006 15:04:05"
	SQL_DATE_FORMAT = "2006-01-02 15:04:05"
)

type Auth struct {
	GorpController
	*revel.Controller
}

func (c Auth) AddUser() revel.Result {
	if user := c.connected(); user != nil {
		c.RenderArgs["user"] = user
	}
	return nil
}

func (c Auth) connected() *models.Admin {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.Admin)
	}
	if username, ok := c.Session["user"]; ok {
		return c.getUser(username)
	}
	return nil
}

func (c Auth) getUserById(id int) *models.User {
	h, err := c.Txn.Get(models.User{}, id)
	if err != nil {
		panic(err)
	}
	if h == nil {
		return nil
	}
	return h.(*models.User)
}

func (c Auth) getUser(username string) *models.Admin {
	users, err := c.Txn.Select(models.Admin{}, `select * from Admin where username = ?`, username)
	if err != nil {
		panic(err)
	}
	if len(users) == 0 {
		return nil
	}
	return users[0].(*models.Admin)
}

func (c Auth) Index() revel.Result {
	if c.connected() != nil {
		return c.Redirect(routes.Auth.Dashboard())
	}
	c.Flash.Error("Please log in first")
	return c.Render()
}

func (c Auth) Login(username, password string, remember bool) revel.Result {
	user := c.getUser(username)
	if user != nil {
		err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
		if err == nil {
			c.Session["user"] = username
			if remember {
				c.Session.SetDefaultExpiration()
			} else {
				c.Session.SetNoExpiration()
			}
			c.Flash.Success("Welcome, " + username)
			return c.Redirect(routes.Auth.Dashboard())
		}
	}
	c.Flash.Out["username"] = username
	c.Flash.Error("Login failed")
	return c.Redirect(routes.Auth.Index())
}

func (c Auth) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.Auth.Index())
}

func (c Auth) Dashboard() revel.Result {
	if c.connected() == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.Auth.Index())
	}

	c.FlashParams()
	day1 := c.Params.Get("day1")
	day2 := c.Params.Get("day2")

	var requests []*models.RequestCount
	var requestsitems []*models.RequestItems
	var requestsplace []*models.RequestPlace

	if day2 == "" || day1 == "" {
		requests = loadRequests(c.Txn.Select(models.RequestCount{}, "SELECT user_id, sum(request_count) as total_request_count, count(requests_id) as total FROM Requests GROUP BY user_id"))

		requestsitems = loadRequestsItems(c.Txn.Select(models.RequestItems{}, "SELECT request_type, sum(request_count) as total FROM Requests GROUP BY request_type"))

		requestsplace = loadRequestPlace(c.Txn.Select(models.RequestPlace{}, "SELECT request_place, sum(request_count) as total_request_count, count(requests_id) as total FROM Requests GROUP BY request_place"))
	} else {
		layout := "01/02/2006 15:04:05"
		t1, _ := time.Parse(layout, day1 + " 00:00:00")
		t2, _ := time.Parse(layout, day2 + " 23:59:59")

		requests = loadRequests(c.Txn.Select(models.RequestCount{}, "SELECT user_id, sum(request_count) as total_request_count, count(requests_id) as total FROM Requests WHERE request_time_str >= ? AND request_time_str <= ? GROUP BY user_id", t1.Format(SQL_DATE_FORMAT), t2.Format(SQL_DATE_FORMAT)))

		requestsitems = loadRequestsItems(c.Txn.Select(models.RequestItems{}, "SELECT request_type, sum(request_count) as total FROM Requests WHERE request_time_str >= ? AND request_time_str <= ? GROUP BY request_type", t1.Format(SQL_DATE_FORMAT), t2.Format(SQL_DATE_FORMAT)))

		requestsplace = loadRequestPlace(c.Txn.Select(models.RequestPlace{}, "SELECT request_place, sum(request_count) as total_request_count, count(requests_id) as total FROM Requests WHERE request_time_str >= ? AND request_time_str <= ? GROUP BY request_place", t1.Format(SQL_DATE_FORMAT), t2.Format(SQL_DATE_FORMAT)))
	}

	return c.Render(requests, requestsitems, requestsplace, day1, day2)
}

func loadRequests(results []interface{}, err error) []*models.RequestCount {
	if err != nil {
		panic(err)
	}
	var requests []*models.RequestCount
	for _, r := range results {
		requests = append(requests, r.(*models.RequestCount))
	}
	return requests
}

func loadRequestsItems(results []interface{}, err error) []*models.RequestItems {
	if err != nil {
		panic(err)
	}
	var requestsitems []*models.RequestItems
	for _, r := range results {
		requestsitems = append(requestsitems, r.(*models.RequestItems))
	}
	return requestsitems
}

func loadRequestPlace(results []interface{}, err error) []*models.RequestPlace {
	if err != nil {
		panic(err)
	}
	var requests []*models.RequestPlace
	for _, r := range results {
		requests = append(requests, r.(*models.RequestPlace))
	}
	return requests
}


func (c Auth) UserList() revel.Result {
	if c.connected() == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.Auth.Index())
	}
	users, err := c.Txn.Select(models.User{}, `select * from User`)

	if err != nil {
		panic(err)
	}

	return c.Render(users)
}

func (c Auth) UserAdd(name string, alias string) revel.Result {
	if c.connected() == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.Auth.Index())
	}
	user := models.User{
		Alias:	alias,
		Name:	name,
	}

	user.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Flash.Out["name"] = name
		c.Flash.Out["alias"] = alias
		c.Flash.Error("Error validation")
		return c.Redirect(routes.Auth.UserList())
	} else {
		if err := c.Txn.Insert(&user); err != nil {
			c.Flash.Out["name"] = name
			c.Flash.Out["alias"] = alias
			c.Flash.Error("Error adding user")
			return c.Redirect(routes.Auth.UserList())
		} else {
			c.Flash.Success("User added")
			return c.Redirect(routes.Auth.UserList())
		}
	}
}

func (c Auth) UserRemove(id int) revel.Result {
	if c.connected() == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.Auth.Index())
	}
	existingUser := c.getUserById(id)
	_, err := c.Txn.Delete(existingUser)

	if err != nil {
		panic(err)
	}

	c.Flash.Success("User removed")
	return c.Redirect(routes.Auth.UserList())
}

func (c Auth) User(id int) revel.Result {
	if c.connected() == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.Auth.Index())
	}
	existingUser := c.getUserById(id)

	c.FlashParams()
	day1 := c.Params.Get("day1")
	day2 := c.Params.Get("day2")

	var requests []*models.RequestPlace
	var requestsitems []*models.RequestItems

	if day2 == "" || day1 == "" {
		requests = loadRequestPlace(c.Txn.Select(models.RequestPlace{}, "SELECT request_place, sum(request_count) as total_request_count, count(requests_id) as total FROM Requests WHERE user_id = ? GROUP BY request_place", existingUser.UserId))

		requestsitems = loadRequestsItems(c.Txn.Select(models.RequestItems{}, "SELECT request_type, sum(request_count) as total FROM Requests WHERE user_id = ? GROUP BY request_type", existingUser.UserId))
	} else {
		layout := "01/02/2006 15:04:05"
		t1, _ := time.Parse(layout, day1 + " 00:00:00")
		t2, _ := time.Parse(layout, day2 + " 23:59:59")

		requests = loadRequestPlace(c.Txn.Select(models.RequestPlace{}, "SELECT request_place, sum(request_count) as total_request_count, count(requests_id) as total FROM Requests WHERE user_id = ? AND request_time_str >= ? AND request_time_str <= ? GROUP BY request_place", existingUser.UserId, t1.Format(SQL_DATE_FORMAT), t2.Format(SQL_DATE_FORMAT)))

		requestsitems = loadRequestsItems(c.Txn.Select(models.RequestItems{}, "SELECT request_type, sum(request_count) as total FROM Requests WHERE user_id = ? AND request_time_str >= ? AND request_time_str <= ? GROUP BY request_type", existingUser.UserId, t1.Format(SQL_DATE_FORMAT), t2.Format(SQL_DATE_FORMAT)))
	}

	return c.Render(requests, requestsitems, existingUser, day1, day2)
}
