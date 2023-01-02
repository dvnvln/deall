package model

import (
	"errors"
	"net/http"
)

type UserReqBody struct {
	Username string
	Password string
}

func (u *UserReqBody) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (u *UserReqBody) Bind(r *http.Request) error {
	return nil
}
func (u *UserReqBody) Validate() error {
	if len(u.Username) < 1 {
		return errors.New("Username must not empty")
	}
	if len(u.Password) < 1 {
		return errors.New("Password must not empty")
	}
	return nil
}
