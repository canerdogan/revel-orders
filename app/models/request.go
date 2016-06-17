package models

import (
	"fmt"
	"github.com/go-gorp/gorp"
	"github.com/revel/revel"
	"time"
)

type Requests struct {
	RequestsId   	int			`db:"requests_id" json:"requests_id"`
	UserId       	int			`db:"user_id" json:"user_id"`
	RequestTimeStr  string		`db:"request_time_str" json:"request_time_str"`
	Alias        	string		`db:"alias" json:"alias"`
	RequestType  	string		`db:"request_type" json:"request_type"`
	RequestCount 	int			`db:"request_count" json:"request_count"`
	IsActive     	bool		`db:"is_active" json:"is_active"`
	RequestPlace	string		`db:"request_place" json:"request_place "`
	// Transient
	RequestTime  time.Time	`db:"request_time" json:"request_time"`
	User *User				`db:"user" json:"user"`
}

const (
	DATE_FORMAT     = "Jan _2, 2006 15:04:05"
	SQL_DATE_FORMAT = "2006-01-02 15:04:05"
)

func (requests Requests) Validate(v *revel.Validation) {
	v.Required(requests.User)
	v.Required(requests.RequestType)
	v.Required(requests.RequestCount)
	v.Required(requests.RequestTime)
	v.Required(requests.RequestPlace)
}

func (r *Requests) PreInsert(_ gorp.SqlExecutor) error {
	r.UserId = r.User.UserId
	r.RequestTimeStr = r.RequestTime.Format(SQL_DATE_FORMAT)
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

	if r.RequestTime, err = time.Parse(SQL_DATE_FORMAT, r.RequestTimeStr); err != nil {
		return fmt.Errorf("Error parsing request date '%s':", r.RequestTimeStr, err)
	}

	return nil
}
