package ginfoodrating

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/foodratings/foodratingsbiz"
	"lesson-5-goland/modules/foodratings/foodratingsmodel"
	"lesson-5-goland/modules/foodratings/foodratingstorage"
	"net/http"
)

func UpdateRating(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateData foodratingsmodel.UpdateFoodRatings
		if err := c.ShouldBind(&updateData); err != nil {
			panic(err)
		}
		rUid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(err)
		}

		store := foodratingstorage.NewSqlStore(appCtx.GetMainDBConnection())
		// biz
		biz := foodratingsbiz.NewUpdateFoodRatingBiz(store)

		if err := biz.UpdateFoodRating(c, int(rUid.GetLocalID()), &updateData); err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(true))

	}
}
