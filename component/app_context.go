package component

import (
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

type TimeJWT struct {
	TimeAccess  int
	TimeRefresh int
}

type appCtx struct {
	db      *mongo.Client
	secret  string
	redis   *redis.Client
	TimeJWT TimeJWT
}

func NewAppContext(db *mongo.Client, secret string, redis *redis.Client, timeJWT TimeJWT) *appCtx {
	return &appCtx{db, secret, redis, timeJWT}
}

func (app *appCtx) GetNewDataMongoDB() (db *mongo.Client) {
	return app.db
}

func (app *appCtx) GetSecret() (secret string) {
	return app.secret
}

func (app *appCtx) GetRedis() (redis *redis.Client) {
	return app.redis
}

type AppContext interface {
	GetNewDataMongoDB() (db *mongo.Client)
	GetSecret() (secret string)
	GetRedis() (redis *redis.Client)
	GetTimeJWT() TimeJWT
}

func (ctx *appCtx) GetTimeJWT() TimeJWT {
	return ctx.TimeJWT
}
