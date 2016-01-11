package models

import (
	"fmt"
	"github.com/canerdogan/revel-orders/Godeps/_workspace/src/github.com/go-gorp/gorp"
	"github.com/canerdogan/revel-orders/Godeps/_workspace/src/github.com/revel/revel"
)

type Requests struct {
	RequestsId   int
	UserId       int
	Alias        string
	RequestType  string
	RequestCount int
	IsActive     bool
	RequestTime  int64

	// Transient
	User *User
}

func (requests Requests) Validate(v *revel.Validation) {
	v.Required(requests.User)
	v.Required(requests.RequestType)
	v.Required(requests.RequestCount)
}

func (r *Requests) PreInsert(_ gorp.SqlExecutor) error {
	r.UserId = r.User.UserId
	return nil
}

func (r *Requests) PostGet(exe gorp.SqlExecutor) error {
	var (
		obj interface{}
		err error
	)

	obj, err = exe.Get(User{}, r.UserId)
	if err != nil {
		return fmt.Errorf("Error loading a requested user (%d): %s", r.UserId, err)
	}
	r.User = obj.(*User)

	return nil
}
