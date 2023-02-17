package ginorder

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/order/orderbiz"
	"lesson-5-goland/modules/order/ordermodel"
	"lesson-5-goland/modules/order/orderstorage"
	"net/http"
)

func CreateOrder(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataOrder ordermodel.CreateOrder

		if err := c.ShouldBind(&dataOrder); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		dataOrder.UserId = requester.GetUserId()

		store := orderstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := orderbiz.NewCreateOrderBiz(store)
		if err := biz.CreateOrder(c, &dataOrder); err != nil {
			panic(err)
		}

		dataOrder.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(dataOrder.FakeId.String()))
	}
}
