package ginfoodrating

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/foodratings/foodratingsbiz"
	"lesson-5-goland/modules/foodratings/foodratingstorage"
	"net/http"
)

func ListFoodRating(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}
		paging.FullFill()

		foodUid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(err)
		}

		store := foodratingstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := foodratingsbiz.NewListFoodRatingBiz(store)

		results, errListRating := biz.ListFoodRating(c, int(foodUid.GetLocalID()), &paging)

		if errListRating != nil {
			panic(errListRating)
		}

		for i := range results {
			results[i].Mask()

			if paging.Limit <= len(results) {
				paging.NextCursor = results[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(results, paging, nil))
	}
}
