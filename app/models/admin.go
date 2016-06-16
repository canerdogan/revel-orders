package models

import (
	"fmt"
	"github.com/revel/revel"
)

type Admin struct {
	AdminId			int			`db:"admin_id" json:"admin_id"`
	Username		string		`db:"username" json:"username"`
	Password		string		`db:"password" json:"password"`
	HashedPassword	[]byte		`db:"hashed_password" json:"hashed_password"`
}

func (u *Admin) String() string {
	return fmt.Sprintf("User(%s)", u.Username)
}

func (u *Admin) Validate(v *revel.Validation) {
	v.Check(u.Username,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{4},
		revel.Match{userRegex},
	)

	v.Check(u.Password,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{5},
	)
}
