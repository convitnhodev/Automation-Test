package component

import "go.mongodb.org/mongo-driver/mongo"

type appCtx struct {
	db *mongo.Client
}

func NewAppContext(db *mongo.Client) *appCtx {
	return &appCtx{db}
}

func (app *appCtx) GetNewDataMongoDB() (db *mongo.Client) {
	return app.db
}

type AppContext interface {
	GetNewDataMongoDB() (db *mongo.Client)
}
