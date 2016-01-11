package app

import (
	"fmt"
	"github.com/canerdogan/emre/Godeps/_workspace/src/github.com/revel/revel"
)

func init() {
	revel.OnAppStart(func() {
		fmt.Println("Go to /@tests to run the tests.")
	})
}
