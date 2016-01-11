package controllers

import (
	"github.com/canerdogan/emre/Godeps/_workspace/src/github.com/revel/revel"
	"github.com/canerdogan/emre/app/chatroom"
	"github.com/canerdogan/emre/app/models"
)

type Application struct {
	GorpController
	*revel.Controller
}

func (c Application) loadRequestsById(id int) *models.Requests {
	h, err := c.Txn.Get(models.Requests{}, id)
	if err != nil {
		panic(err)
	}
	if h == nil {
		return nil
	}
	return h.(*models.Requests)
}

func (c Application) Index() revel.Result {
	requests, err := c.Txn.Select(models.Requests{}, `SELECT * FROM Requests WHERE IsActive > 0`)
	if err != nil {
		return c.RenderText("Error trying to get records from DB.")
	}

	return c.Render(requests)
}

func (c Application) Remove(id int) revel.Result {
	request := c.loadRequestsById(id)
	if request == nil {
		return c.NotFound("Request %d does not exist", id)
	}

	request.IsActive = false
	_, err := c.Txn.Update(request)

	if err != nil {
		return c.NotFound("Request %d can't up to date", id)
	}

	return c.Redirect(Application.Index)
}

func (c Application) Api(username string, requestType string, requestCount int) revel.Result {
	var user models.User
	if err := c.Txn.SelectOne(&user, "select * from User where Alias=?", username); err != nil {
		return c.RenderText("We couldn't find User Alias")
	} else {
		request := models.Requests{
			Alias:        username,
			RequestType:  requestType,
			RequestCount: requestCount,
			IsActive:     true,
			User:         &user,
		}

		request.Validate(c.Validation)
		if c.Validation.HasErrors() {
			return c.RenderText("Error loading a requested user")
		} else {
			if err := c.Txn.Insert(&request); err != nil {
				return c.RenderText(
					"Error inserting record into database!")
			} else {
				chatroom.SendRequest(&request)
				return c.RenderJson(request)
			}
		}
	}
}
