package controller

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/dvnvln/deall/model"
	"github.com/dvnvln/deall/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *Controller) CreateUser(ctx context.Context, payload model.CreateUserReqBody) (model.UserRes, error) {
	res := model.UserRes{
		IsSuccess: false,
		Time:      time.Now(),
	}
	pwd, err := utils.HashAndSalt(payload.Password)
	if err != nil {
		return res, err
	}
	user := model.User{
		ID:       primitive.NewObjectID(),
		Username: payload.Username,
		Password: pwd,
		Role: model.Details{
			RoleName: payload.RoleName,
		},
		CreatedAt: time.Now(),
	}

	err = c.repo.Add(ctx, user)
	if err != nil {
		return res, err
	}
	res.IsSuccess = true
	return res, nil
}

func (c *Controller) GetUser(ctx context.Context, userID string) ([]model.User, error) {
	user, err := c.repo.GetByUserID(ctx, userID)
	return user, err
}

func (c *Controller) ListUser(ctx context.Context) ([]model.User, error) {
	users, err := c.repo.Get(ctx)
	return users, err
}

func (c *Controller) UpdateUser(ctx context.Context, userID string, payload model.CreateUserReqBody) (model.UserRes, error) {
	res := model.UserRes{
		IsSuccess: false,
		Time:      time.Now(),
	}
	pwd, err := utils.HashAndSalt(payload.Password)
	user := model.User{
		Username: payload.Username,
		Password: pwd,
		Role: model.Details{
			RoleName: payload.RoleName,
		},
		UpdatedAt: time.Now(),
	}
	log.Printf(payload.Username)
	err = c.repo.Update(ctx, userID, user)
	if err != nil {
		return res, err
	}
	res.IsSuccess = true
	return res, nil
}

func (c *Controller) DeleteUser(ctx context.Context, userID string) (model.UserRes, error) {
	err := c.repo.Delete(ctx, userID)
	if err != nil {
		return model.UserRes{}, err
	}
	res := model.UserRes{
		IsSuccess: true,
		Time:      time.Now(),
	}
	return res, nil
}

func (c *Controller) UserLogin(ctx context.Context, payload model.UserReqBody) (model.AuthSuccess, error) {
	log.Print("controller here..")
	log.Print("Calling repo ..")
	user, err := c.repo.Login(ctx, payload)
	if err != nil {
		return model.AuthSuccess{}, err
	}

	isSamePass, err := utils.ComparePasswords(user.Password, payload.Password)
	if err != nil || !isSamePass {
		res := model.AuthSuccess{
			IsSuccess: false,
			Token:     "",
		}
		return res, err
	}
	token, err := utils.GetToken(os.Getenv("SECRET")).MakeToken(user.ID.Hex(), user.Role.RoleName)
	resAuth := model.AuthSuccess{
		IsSuccess: err == nil,
		Token:     token,
	}
	log.Print("controller done here..")
	return resAuth, err
}
