package userStorage

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/modules/user/userModel"
	"context"
	"errors"
)

func (db *mongoStore) FindUser(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error) {
	collection := db.db.Database("AutomationTest").Collection("User")

	var data *userModel.User

	if err := collection.FindOne(ctx, conditions).Decode(data); err != nil {
		component.ErrorLogger.Println("Can't Insert to DB, something DB is error")
		return nil, common.ErrDB(err)
	}

	if data == nil {
		return nil, errors.New("record not found")
	}

	return data, nil
}
