package models

import (
	"fmt"
	"github.com/go-gorp/gorp"
)

type RequestCount struct {
	UserId       		int			`db:"user_id" json:"user_id"`
	RequestType       	string		`db:"request_type" json:"request_type"`
	Total 				int			`db:"total" json:"total"`
	TotalRequestCount	int			`db:"total_request_count" json:"total_request_count"`

	// Transient
	User *User				`db:"user" json:"user"`
}

func (r *RequestCount) PreInsert(_ gorp.SqlExecutor) error {
	r.UserId = r.User.UserId
	return nil
}

func (r *RequestCount) PostGet(exe gorp.SqlExecutor) error {
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
