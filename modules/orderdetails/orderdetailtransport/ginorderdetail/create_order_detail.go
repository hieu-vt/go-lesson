package ginorderdetail

import (
	"encoding/json"
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

		orderId, err := common.FromBase58(orderDetail.OrderId)

		if err != nil {
			panic(err)
		}

		jFoodOrigin, fOriginErr := json.Marshal(orderDetail.FoodOrigin)

		if fOriginErr != nil {
			panic(fOriginErr)
		}

		orderDetailCreated := orderdetailmodel.OrderDetail{
			OrderId:    int(orderId.GetLocalID()),
			FoodOrigin: string(jFoodOrigin),
			Price:      orderDetail.Price,
			Quantity:   orderDetail.Quantity,
			Discount:   orderDetail.Discount,
		}

		if err := biz.CreateOrderDetail(c, &orderDetailCreated); err != nil {
			panic(err)
		}

		orderDetailCreated.Mask()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(orderDetailCreated.FakeId))
	}
}
