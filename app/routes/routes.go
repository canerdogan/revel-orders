// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/canerdogan/revel-orders/Godeps/_workspace/src/github.com/revel/revel"

type tWebSocket struct{}

var WebSocket tWebSocket

func (_ tWebSocket) RoomSocket(
	ws interface{},
) string {
	args := make(map[string]string)

	revel.Unbind(args, "ws", ws)
	return revel.MainRouter.Reverse("WebSocket.RoomSocket", args).Url
}

type tApplication struct{}

var Application tApplication

func (_ tApplication) Index() string {
	args := make(map[string]string)

	return revel.MainRouter.Reverse("Application.Index", args).Url
}

func (_ tApplication) Remove(
	id int,
) string {
	args := make(map[string]string)

	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Application.Remove", args).Url
}

func (_ tApplication) Api(
	username string,
	requestType string,
	requestCount int,
) string {
	args := make(map[string]string)

	revel.Unbind(args, "username", username)
	revel.Unbind(args, "requestType", requestType)
	revel.Unbind(args, "requestCount", requestCount)
	return revel.MainRouter.Reverse("Application.Api", args).Url
}

type tGorpController struct{}

var GorpController tGorpController

func (_ tGorpController) Begin() string {
	args := make(map[string]string)

	return revel.MainRouter.Reverse("GorpController.Begin", args).Url
}

func (_ tGorpController) Commit() string {
	args := make(map[string]string)

	return revel.MainRouter.Reverse("GorpController.Commit", args).Url
}

func (_ tGorpController) Rollback() string {
	args := make(map[string]string)

	return revel.MainRouter.Reverse("GorpController.Rollback", args).Url
}

type tStatic struct{}

var Static tStatic

func (_ tStatic) Serve(
	prefix string,
	filepath string,
) string {
	args := make(map[string]string)

	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
	moduleName string,
	prefix string,
	filepath string,
) string {
	args := make(map[string]string)

	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}

type tTestRunner struct{}

var TestRunner tTestRunner

func (_ tTestRunner) Index() string {
	args := make(map[string]string)

	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
	suite string,
	test string,
) string {
	args := make(map[string]string)

	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List() string {
	args := make(map[string]string)

	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}
