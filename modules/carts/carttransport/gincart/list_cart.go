package gincart

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"net/http"
)

func ListCart(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}

		paging.FullFill()

		c.JSON(http.StatusOK, common.NewSuccessResponse(nil, paging, nil))
	}
}
