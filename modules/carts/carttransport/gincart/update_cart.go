package gincart

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/carts/cartmodel"
	"net/http"
)

func UpdateCart(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		cartId, err := common.FromBase58(c.Param("cartId"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var body cartmodel.UpdateCart

		if err := c.ShouldBind(&body); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(cartId))
	}
}
