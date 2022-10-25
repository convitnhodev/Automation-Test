package main

import (
	"backend_autotest/modules/command/commandTransport"
	"backend_autotest/modules/node/nodeTransport"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"

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

}

func runService(db *mongo.Client) error {
	r := gin.Default()
	appCtx := component.NewAppContext(db)

	user := r.Group("/user")
	{
		user.POST("/register", userTransport.UserRegister(appCtx))
	}

	node := r.Group("/node")
	{
		node.POST("/register", nodeTransport.NodeRegister(appCtx))
		node.DELETE("/delete", nodeTransport.NodeDelete(appCtx))
		node.GET("/list", nodeTransport.NodeList(appCtx))
		node.POST("/result", nodeTransport.NodePostResult(appCtx))
		node.GET("/result", nodeTransport.NodeGetResult(appCtx))
	}

	command := node.Group("/command")
	{
		command.POST("/new", commandTransport.NewNodeCommand(appCtx))
		command.GET("/get", commandTransport.GetAndDeleteCommand(appCtx))

	}

	return r.Run(":8080")
}
