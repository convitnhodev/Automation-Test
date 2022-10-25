package userModel

import (
	"errors"
	"strings"
)

type User struct {
	UserName string `bson:"username,omitempty"`
	Password string `bson:"password"`
	FullName string `bson:"fullname"`
	Company  string `bson:"company,omitempty"`
}

func (user *User) Validate() error {

	//check validate of email
	user.UserName = strings.TrimSpace(user.UserName)
	user.FullName = strings.TrimSpace(user.FullName)
	user.Company = strings.TrimSpace(user.Company)
	user.Password = strings.TrimSpace(user.Password)

	if user.UserName == "" {
		return errors.New("username name can not be blank")
	}

	if user.Password == "" {
		return errors.New("password name can not be blank")
	}

	if user.FullName == "" {
		return errors.New("fullname name can not be blank")
	}

	return nil
}
