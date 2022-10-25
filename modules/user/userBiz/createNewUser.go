package userBiz

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/modules/user/userModel"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type CreateUserStore interface {
	CreateUser(ctx context.Context, data *userModel.User) error
	FindUser(ctx context.Context, conditions interface{}) (*userModel.User, error)
}

type createUserBiz struct {
	store CreateUserStore
}

func NewCreateUserBiz(store CreateUserStore) *createUserBiz {
	return &createUserBiz{store}
}

func (biz *createUserBiz) CreateNewUser(ctx context.Context, data *userModel.User) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if user, err := biz.store.FindUser(ctx, bson.M{"username": data.UserName}); user != nil {
		component.InfoLogger.Println("User is registed before")
		return common.ErrEntityExisted("User Register", err)
	}

	if err := biz.store.CreateUser(ctx, data); err != nil {
		component.InfoLogger.Println("Can not Create User")
		return common.ErrCannotCreateEntity("User Registerd", err)
	}

	return nil
}
