package users

import (
	"strings"

	"github.com/fdiaz7/bookstore_users-api/utils/errors"
)

const (
	StatusActive = "active"
)

//User domain
type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

//Validate email
func (u *User) Validate() *errors.RestErr {
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	u.Password = strings.TrimSpace(u.Password)
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	if u.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
