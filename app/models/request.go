package models

import (
	"fmt"
	"github.com/go-gorp/gorp"
	"github.com/revel/revel"
)

type Requests struct {
	RequestsId   int		`db:"requests_id" json:"requests_id"`
	UserId       int		`db:"user_id" json:"user_id"`
	Alias        string		`db:"alias" json:"alias"`
	RequestType  string		`db:"request_type" json:"request_type"`
	RequestCount int		`db:"request_count" json:"request_count"`
	IsActive     bool		`db:"is_active" json:"is_active"`
	RequestTime  int64		`db:"request_time" json:"request_time"`

	// Transient
	User *User				`db:"user" json:"user"`
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
