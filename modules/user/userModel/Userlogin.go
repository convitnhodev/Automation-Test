package userModel

import (
	"backend_autotest/common"
	"errors"
)

type UserLogin struct {
	UserName string `json:"user_name" bson:"user_name,omitempty"`
	Password string `json:"password"bson:"password"`
}

func (u UserLogin) TableName() string {
	return "users"
}

type PASSWORD string

func (pass PASSWORD) Validate() error {
	if len(pass) < 6 {
		return common.ErrInvalidPassword(errors.New("password must be at least 6 characters"))
	}

	return nil

}
