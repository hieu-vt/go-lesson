package ginfoodrating

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/foodratings/foodratingstorage"
	"net/http"
)

func DeleteRating(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		rUid := c.Param("id")

		store := foodratingstorage.NewSqlStore(appCtx.GetMainDBConnection())
		// biz

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(true))

	}
}
