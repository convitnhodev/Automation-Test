package main

import (
	"backend_autotest/component"
	"backend_autotest/modules/user/userTransport"
	"context"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//dsn := "root:Thaothaothao223051@@tcp(localhost:3306)/autotest?charset=utf8mb4&parseTime=True&loc=Local"
	//db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//
	//fmt.Println(db)

	// // listen to port
	// http.ListenAndServe(":5050", nil)

	db := connectMongoDB()
	app := component.NewAppContext(db)

	http.HandleFunc("/user/register", userTransport.UserRegister(app))

	http.ListenAndServe(":5050", nil)

}

func connectMongoDB() *mongo.Client {
	ctx := context.TODO()
	opts := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		component.ErrorLogger.Println("Can not connect to MongoDB")
		panic(err)
	}
	component.InfoLogger.Println("Connect DB successly")
	defer client.Disconnect(ctx)
	fmt.Printf("%T\n", client)

	dbNames, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	fmt.Println(dbNames)

	return client

}
