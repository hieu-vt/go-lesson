package ginfood

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/food/foodbiz"
	"lesson-5-goland/modules/food/foodstorage"
	"net/http"
)

func GetFoods(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}

		paging.FullFill()

		store := foodstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := foodbiz.NewGetFoodBiz(store)

		result, err := biz.GetFood(c, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask()

			if paging.Limit <= len(result) {
				paging.NextCursor = result[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
