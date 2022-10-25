package commandBiz

import (
	"backend_autotest/common"
	"backend_autotest/modules/command/commandModel"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type FindCommandStore interface {
	GetCommandAndDelete(ctx context.Context, conditions interface{}) (*commandModel.Command, error)
}

type findCommandBiz struct {
	store FindCommandStore
}

func NewFindCommandBiz(store FindCommandStore) *findCommandBiz {
	return &findCommandBiz{store}
}

func (biz *findCommandBiz) FindCommandAndDelete(ctx context.Context, data *commandModel.Node) (*commandModel.Command, error) {
	result, err := biz.store.GetCommandAndDelete(ctx, bson.M{"node_id": data.NodeId})
	if err != nil {
		return nil, common.ErrCannotDeleteEntity("Node Command", err)
	}
	return result, nil
}
