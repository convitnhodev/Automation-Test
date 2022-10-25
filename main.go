package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/modules/user/userTransport"
)

func main() {
	db := common.InitMongoDB()

	fmt.Println(db)

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
	fmt.Print("Hello")
}

func runService(db *mongo.Client) error {
	r := gin.Default()
	appCtx := component.NewAppContext(db)

	user := r.Group("/user")
	{
		user.POST("/register", userTransport.UserRegister(appCtx))
	}

	return r.Run(":8080")
}
