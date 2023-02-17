package ginfoodrating

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/foodratings/foodratingsmodel"
	"lesson-5-goland/modules/foodratings/foodratingstorage"
	"net/http"
)

func UpdateRating(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateData foodratingsmodel.CreateFoodRatings
		if err := c.ShouldBind(&updateData); err != nil {
			panic(err)
		}
		rUid := c.Param("id")

		store := foodratingstorage.NewSqlStore(appCtx.GetMainDBConnection())
		// biz

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(true))

	}
}
