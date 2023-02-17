package ginfoodrating

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"net/http"
)

func ListFoodRating(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}

		//requester := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(http.StatusOK, common.NewSuccessResponse(nil, paging, nil))
	}
}
