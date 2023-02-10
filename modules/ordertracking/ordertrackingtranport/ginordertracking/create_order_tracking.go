package ginordertracking

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/ordertracking/ordertrackingbiz"
	"lesson-5-goland/modules/ordertracking/ordertrackingmodel"
	"lesson-5-goland/modules/ordertracking/ordertrackingstorage"
	"net/http"
)

func CreateOrderTracking(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data ordertrackingmodel.CreateOrderTrackingParams

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := ordertrackingstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := ordertrackingbiz.NewOrderTrackingBiz(store)

		orderId, err := common.FromBase58(data.OrderId)

		if err != nil {
			panic(err)
		}

		createOrder := ordertrackingmodel.OrderTracking{
			OrderId: int(orderId.GetLocalID()),
			State:   data.State,
		}

		if err := biz.CreateOrderTracking(c, &createOrder); err != nil {
			panic(err)
		}

		createOrder.Mask()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(createOrder.FakeId))
	}
}
