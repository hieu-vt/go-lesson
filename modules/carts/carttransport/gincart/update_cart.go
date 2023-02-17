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

func UpdateCart(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		foodId, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var body cartmodel.UpdateCart

		if err := c.ShouldBind(&body); err != nil {
			panic(err)
		}

		store := cartstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := cartbiz.NewBizUpdateCart(store)

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		if err := biz.UpdateCart(c, requester.GetUserId(), int(foodId.GetLocalID()), &body); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
