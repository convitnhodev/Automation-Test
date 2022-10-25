package nodeBiz

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/modules/node/nodeModel"
	"context"
)

type AddSeftNodeStore interface {
	FindNode(ctx context.Context, conditions map[string]interface{}) (*nodeModel.Node, error)
	CreateStore(ctx context.Context, data *nodeModel.Node) error
}

type addSeftNodeBiz struct {
	store AddSeftNodeStore
}

func NewAddSeftNodeBiz(store AddSeftNodeStore) *addSeftNodeBiz {
	return &addSeftNodeBiz{store}
}

func (biz *addSeftNodeBiz) AddNewSeftNode(ctx context.Context, data *nodeModel.Node) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if node, err := biz.store.FindNode(ctx, map[string]interface{}{"serial_number": data.SerialNumber}); node != nil {
		component.InfoLogger.Println("Node is registed before")
		return common.ErrEntityExisted(" Node Register", err)
	}

	if err := biz.store.CreateStore(ctx, data); err != nil {
		component.InfoLogger.Println("Can not Create Node")
		return common.ErrCannotCreateEntity("Node Registerd", err)
	}

	return nil
}
