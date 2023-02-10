package gincart

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/carts/cartmodel"
	"net/http"
)

func DeleteCart(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var listId cartmodel.DeleteCart

		if err := c.ShouldBind(&listId); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse("ok"))
	}
}
