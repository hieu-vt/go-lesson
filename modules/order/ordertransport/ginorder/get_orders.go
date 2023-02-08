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
		userId, err := common.FromBase58(c.Param("userId"))

		if err != nil {
			panic(err)
		}

		store := orderstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := orderbiz.NewGetOrderBiz(store)

		data, err := biz.GetOrders(c, int(userId.GetLocalID()))

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
