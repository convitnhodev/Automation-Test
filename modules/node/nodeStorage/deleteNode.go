package nodeStorage

import (
	"backend_autotest/common"
	"context"
)

func (db *mongoStore) DeleteNode(ctx context.Context, conditions interface{}) error {
	collection := db.db.Database("AutomationTest").Collection("Node")

	_, err := collection.DeleteOne(ctx, conditions)
	if err != nil {

		return common.ErrDB(err)
	}

	return nil
	// delete oke
}
