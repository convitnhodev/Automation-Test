package commandStorage

import (
	"backend_autotest/common"
	"backend_autotest/modules/command/commandModel"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *mongoStore) GetCommandAndDelete(ctx context.Context, conditions interface{}) (*commandModel.Command, error) {
	var tmp bson.M
	collection := db.db.Database("AutomationTest").Collection("Each_Node_Command")
	err := collection.FindOneAndDelete(ctx, conditions).Decode(&tmp)
	if err != nil {
		return nil, common.ErrDB(err)
	}

	var result commandModel.Command
	bsonBytes, _ := bson.Marshal(tmp)
	bson.Unmarshal(bsonBytes, &result)
	return &result, nil

}
