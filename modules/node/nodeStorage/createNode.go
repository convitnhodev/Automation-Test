package nodeStorage

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/modules/node/nodeModel"
	"context"
)

func (db *mongoStore) CreateStore(ctx context.Context, data *nodeModel.Node) error {
	collection := db.db.Database("AutomationTest").Collection("Node")

	if _, err := collection.InsertOne(ctx, data); err != nil {
		component.ErrorLogger.Println("Can't Insert to DB, something DB is error")
		return common.ErrDB(err)
	}

	return nil
}
