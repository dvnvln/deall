package handler

import (
	"context"

	"github.com/dvnvln/deallscrud/model"
)

//go:generate mockgen -build_flags=-mod=mod -source=dependencies.go -package=todo -destination=dependencies.mock.test.go

type userController interface {
	GetUser(ctx context.Context, userID string) ([]model.User, error)
	ListUser(ctx context.Context) ([]model.User, error)
	CreateUser(ctx context.Context, payload model.CreateUserReqBody) (model.UserRes, error)
	UpdateUser(ctx context.Context, userID string, payload model.CreateUserReqBody) (model.UserRes, error)
	DeleteUser(ctx context.Context, userID string) (model.UserRes, error)
	UserLogin(ctx context.Context, payload model.UserReqBody) (model.AuthSuccess, error)
}
