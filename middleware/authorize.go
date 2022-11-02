package middleware

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/component/tokenprovider/jwt"
	"backend_autotest/modules/user/userStorage"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	// split Bearer and token
	parts := strings.Split(s, " ")
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}
	return parts[1], nil
}

func RequireAuth(appCtx component.AppContext) func(c *gin.Context) {

	tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecret())

	return func(c *gin.Context) {

		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		db := appCtx.GetNewDataMongoDB()
		store := userStorage.NewMongoStore(db)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			c.JSON(401, "token is invalid")
			panic(err)
		}

		user, err := store.FindUser(c.Request.Context(), bson.M{"user_name": payload.UserName})
		if err != nil {
			panic(err)
		}

		c.Set(component.CurrentUser, user)
		c.Next()
	}
}
