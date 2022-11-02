package userBiz

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/component/tokenprovider"
	"backend_autotest/modules/user/userModel"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions interface{}) (*userModel.User, error)
}

type Hasher interface {
	Hash(data string) string
}

type setTime struct {
	timeAccess int
}

func NewSetTime(timeAccess int) *setTime {
	return &setTime{
		timeAccess: timeAccess,
	}
}

type TokenConfig interface {
	GetAtExp() int
}

func (timeSet *setTime) GetAtExp() int {
	return timeSet.timeAccess
}

type loginBusiness struct {
	appCtx        component.AppContext
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	tkCfg         TokenConfig
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher, tkCfg TokenConfig) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		tkCfg:         tkCfg,
	}
}

func (biz *loginBusiness) Login(ctx context.Context, data *userModel.UserLogin) (*tokenprovider.Account, error) {
	user, err := biz.storeUser.FindUser(ctx, bson.M{"user_name": data.UserName})
	if err != nil {
		if err.Error() != common.RecordNotFound {
			return nil, common.ErrDB(err)
		}
		return nil, common.ErrInvalidLogin(err)
	}

	passHash := biz.hasher.Hash(data.Password)

	if user.Password != passHash {
		return nil, common.ErrInvalidLogin(err)
	}

	payload := tokenprovider.TokenPayload{
		UserName: user.UserName,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.tkCfg.GetAtExp())
	if err != nil {
		return nil, common.GenerateJWTFail(err)
	}

	account := tokenprovider.Account{
		AccessToken: accessToken,
	}

	return &account, nil

}
