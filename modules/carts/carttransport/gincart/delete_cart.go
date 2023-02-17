package gincart

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/carts/cartbiz"
	"lesson-5-goland/modules/carts/cartstorage"
	"net/http"
)

func DeleteCart(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		foodId, err := common.FromBase58(c.Query("foodId"))

		if err != nil {
			panic(err)
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := cartstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := cartbiz.NewDeleteCartBiz(store)

		if err := biz.DeleteCarts(c, requester.GetUserId(), int(foodId.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse("ok"))
	}
}
