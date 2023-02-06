package ginorderdetail

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/orderdetails/orderdetailbiz"
	"lesson-5-goland/modules/orderdetails/orderdetailmodel"
	"lesson-5-goland/modules/orderdetails/orderdetailstorage"
	"net/http"
)

func CreateOrderDetail(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var orderDetail orderdetailmodel.CreateOrderDetail

		if err := c.ShouldBind(&orderDetail); err != nil {
			panic(err)
		}

		store := orderdetailstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := orderdetailbiz.NewOrderDetailBiz(store)

		if err := biz.CreateOrderDetail(c, &orderDetail); err != nil {
			panic(err)
		}

		orderDetail.Mask()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(orderDetail.FakeId))
	}
}
