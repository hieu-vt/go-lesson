package ginordertracking

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/ordertracking/ordertrackingmodel"
	"net/http"
)

func CreateOrderTracking(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data ordertrackingmodel.OrderTracking

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		data.Mask()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId))
	}
}
