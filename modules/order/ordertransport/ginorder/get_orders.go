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

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.FullFill()

		store := orderstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := orderbiz.NewGetOrderBiz(store)

		result, err := biz.GetOrders(c, int(requester.GetUserId()), paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask()

			if paging.Limit <= len(result) {
				paging.NextCursor = result[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
