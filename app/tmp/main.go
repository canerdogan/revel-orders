// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	_ "github.com/canerdogan/emre/Godeps/_workspace/src/github.com/mattn/go-sqlite3"
	controllers0 "github.com/canerdogan/emre/Godeps/_workspace/src/github.com/revel/modules/static/app/controllers"
	_ "github.com/canerdogan/emre/Godeps/_workspace/src/github.com/revel/modules/testrunner/app"
	controllers1 "github.com/canerdogan/emre/Godeps/_workspace/src/github.com/revel/modules/testrunner/app/controllers"
	"github.com/canerdogan/emre/Godeps/_workspace/src/github.com/revel/revel"
	"github.com/canerdogan/emre/Godeps/_workspace/src/github.com/revel/revel/testing"
	websocket "github.com/canerdogan/emre/Godeps/_workspace/src/golang.org/x/net/websocket"
	_ "github.com/canerdogan/emre/app"
	_ "github.com/canerdogan/emre/app/chatroom"
	controllers "github.com/canerdogan/emre/app/controllers"
	tests "github.com/canerdogan/emre/tests"
	"reflect"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")

	revel.RegisterController((*controllers.WebSocket)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "RoomSocket",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "ws", Type: reflect.TypeOf((**websocket.Conn)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
		})

	revel.RegisterController((*controllers.Application)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{},
				RenderArgNames: map[int][]string{
					31: []string{
						"requests",
					},
				},
			},
			&revel.MethodType{
				Name: "Remove",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name: "Api",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "requestType", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "requestCount", Type: reflect.TypeOf((*int)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
		})

	revel.RegisterController((*controllers.GorpController)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name:           "Begin",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name:           "Commit",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name:           "Rollback",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
		})

	revel.RegisterController((*controllers0.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
		})

	revel.RegisterController((*controllers1.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{},
				RenderArgNames: map[int][]string{
					70: []string{
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{
					107: []string{},
				},
			},
			&revel.MethodType{
				Name:           "List",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
		})

	revel.DefaultValidationKeys = map[string]map[int]string{
		"github.com/canerdogan/emre/app/models.(*User).Validate": {
			27: "user.Alias",
			32: "user.Name",
		},
		"github.com/canerdogan/emre/app/models.Requests.Validate": {
			23: "requests.User",
			24: "requests.RequestType",
			25: "requests.RequestCount",
		},
	}
	testing.TestSuites = []interface{}{
		(*tests.ApplicationTest)(nil),
	}

	revel.Run(*port)
}
