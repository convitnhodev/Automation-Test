package nodeStorage

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/modules/node/nodeModel"
	"context"
	"errors"
)

func (db *mongoStore) FindNode(ctx context.Context, conditions map[string]interface{}) (*nodeModel.Node, error) {
	collection := db.db.Database("AutomationTest").Collection("Node")

	var data *nodeModel.Node

	if err := collection.FindOne(ctx, conditions).Decode(data); err != nil {
		component.ErrorLogger.Println("Can't Insert to DB, somtthing DB is error")
		return nil, common.ErrDB(err)
	}

	if data == nil {
		return nil, errors.New("record not found")
	}

	return data, nil
}
