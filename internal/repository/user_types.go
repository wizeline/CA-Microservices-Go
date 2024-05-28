package repository

import (
	"time"

	"github.com/wizeline/CA-Microservices-Go/internal/util"
)

type UserCreateArgs struct {
	FirstName string
	LastName  string
	Birthday  time.Time
	Email     string
	Username  string
	Passwd    string
}

func (args UserCreateArgs) Validate() error {
	if args == (UserCreateArgs{}) {
		return &EntityEmptyErr{Name: "UserCreateArgs"}
	}
	if args.FirstName == "" {
		return &InvalidFieldErr{Name: "FirstName", Err: util.ErrEmptyValue}
	}
	if args.LastName == "" {
		return &InvalidFieldErr{Name: "LastName", Err: util.ErrEmptyValue}
	}
	if args.Birthday.IsZero() {
		return &InvalidFieldErr{Name: "Birthday", Err: util.ErrEmptyValue}
	}
	if err := util.ValidateEmail(args.Email); err != nil {
		return err
	}
	if args.Username == "" {
		return &InvalidFieldErr{Name: "Username", Err: util.ErrEmptyValue}
	}
	if args.Passwd == "" {
		return &InvalidFieldErr{Name: "Passwd", Err: util.ErrEmptyValue}
	}
	return nil
}
