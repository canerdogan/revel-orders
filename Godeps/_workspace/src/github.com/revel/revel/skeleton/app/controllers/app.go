package controllers

import "github.com/canerdogan/emre/Godeps/_workspace/src/github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}