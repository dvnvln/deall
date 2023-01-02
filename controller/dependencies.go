package controller

import (
	"context"

	"github.com/dvnvln/deall/model"
)

//go:generate mockgen -build_flags=-mod=mod -source=dependencies.go -package=todo -destination=dependencies.mock.test.go

type userRepository interface {
	Get(ctx context.Context) ([]model.User, error)
	GetByUserID(ctx context.Context, userID string) ([]model.User, error)
	Add(ctx context.Context, user model.User) error
	Update(ctx context.Context, userID string, user model.User) error
	Delete(ctx context.Context, userID string) error
	Login(ctx context.Context, user model.UserReqBody) (model.User, error)
}
