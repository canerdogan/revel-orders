package models

import (
	"fmt"
	"github.com/canerdogan/revel-orders/Godeps/_workspace/src/github.com/revel/revel"
	"regexp"
)

type User struct {
	UserId int
	Name   string
	Alias  string
}

func (u *User) String() string {
	return fmt.Sprintf("User(%s)", u.Alias)
}

var userRegex = regexp.MustCompile("^\\w*$")

func (user *User) Validate(v *revel.Validation) {
	v.Check(user.Alias,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{4},
		revel.Match{userRegex},
	)

	v.Check(user.Name,
		revel.Required{},
		revel.MaxSize{100},
	)
}
