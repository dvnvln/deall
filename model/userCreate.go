package model

import (
	"errors"
	"net/http"
)

type CreateUserReqBody struct {
	Username string
	Password string
	RoleName string
}

func (u *CreateUserReqBody) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *CreateUserReqBody) Bind(r *http.Request) error {
	return nil
}

func (u *CreateUserReqBody) Validate() error {
	if len(u.Username) < 1 {
		return errors.New("Username must not empty")
	}
	if len(u.Password) < 1 {
		return errors.New("Password must not empty")
	}
	if len(u.RoleName) < 1 {
		return errors.New("RoleName must not empty")
	}
	return nil
}
