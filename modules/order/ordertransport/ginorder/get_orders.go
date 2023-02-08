package ginorder

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/order/orderbiz"
	"lesson-5-goland/modules/order/orderstorage"
	"net/http"
)

func GetOrders(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := orderstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := orderbiz.NewGetOrderBiz(store)

		data, err := biz.GetOrders(c, int(requester.GetUserId()))

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
