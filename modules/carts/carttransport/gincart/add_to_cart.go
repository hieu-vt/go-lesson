package gincart

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/carts/cartbiz"
	"lesson-5-goland/modules/carts/cartmodel"
	"lesson-5-goland/modules/carts/cartstorage"
	"net/http"
)

func AddToCart(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body cartmodel.CreateCart

		if err := c.ShouldBind(&body); err != nil {
			panic(err)
		}

		foodId, err := common.FromBase58(body.FoodId)

		if err != nil {
			panic(err)
		}

		store := cartstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := cartbiz.NewCreateCartBiz(store)

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var result cartmodel.Cart

		result = cartmodel.Cart{
			SqlModel: common.SqlModel{
				Status: 1,
			},
			UserId:   int(requester.GetUserId()),
			FoodId:   int(foodId.GetLocalID()),
			Quantity: body.Quantity,
		}

		if err := biz.CreateCart(c, &result); err != nil {
			panic(err)
		}

		result.Mask()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result.FakeId))
	}
}
