package commandTransport

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/modules/command/commandBiz"
	"backend_autotest/modules/command/commandModel"
	"backend_autotest/modules/command/commandStorage"
	"github.com/gin-gonic/gin"
)

func GetAndDeleteCommand(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var data commandModel.Node

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := commandStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := commandBiz.NewFindCommandBiz(store)

		result, err := biz.FindCommandAndDelete(c.Request.Context(), &data)
		if err != nil {
			c.JSON(400, err)
		}

		c.JSON(200, common.SimpleSuccessResponse(result))
	}
}
