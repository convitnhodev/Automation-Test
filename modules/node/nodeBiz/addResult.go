package nodeBiz

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/modules/node/nodeModel"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type AddResultStore interface {
	PostResult(ctx context.Context, data *nodeModel.Node) (*mongo.InsertOneResult, error)
}

type addResultBiz struct {
	store AddResultStore
}

func NewAddResultBiz(store AddResultStore) *addResultBiz {
	return &addResultBiz{store}
}

func (biz *addResultBiz) AddNewResult(ctx context.Context, data *nodeModel.Node) (*mongo.InsertOneResult, error) {

	result, err := biz.store.PostResult(ctx, data)
	if err != nil {
		component.InfoLogger.Println("Can not Create Node")
		return nil, common.ErrCannotCreateEntity("Node Registerd", err)
	}

	return result, nil
}
