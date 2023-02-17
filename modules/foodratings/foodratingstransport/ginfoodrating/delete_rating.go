package ginfoodrating

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/foodratings/foodratingsbiz"
	"lesson-5-goland/modules/foodratings/foodratingstorage"
	"net/http"
)

func DeleteRating(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		rUid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(err)
		}

		store := foodratingstorage.NewSqlStore(appCtx.GetMainDBConnection())
		// biz
		biz := foodratingsbiz.NewDeleteFoodRatingBiz(store)

		if err := biz.DeleteFoodRating(c, int(rUid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(true))

	}
}
