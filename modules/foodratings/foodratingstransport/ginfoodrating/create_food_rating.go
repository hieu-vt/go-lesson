package ginfoodrating

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/foodratings/foodratingsmodel"
	"lesson-5-goland/modules/foodratings/foodratingstorage"
	"net/http"
)

func CreateFoodRating(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body foodratingsmodel.CreateFoodRatings
		if err := c.ShouldBind(&body); err != nil {
			panic(err)
		}

		foodUid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(err)
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := foodratingstorage.NewSqlStore(appCtx.GetMainDBConnection())
		newFoodRating := foodratingsmodel.FoodRatings{
			SqlModel: common.SqlModel{
				Status: 1,
			},
			UserId:  requester.GetUserId(),
			FoodId:  int(foodUid.GetLocalID()),
			Point:   body.Point,
			Comment: body.Comment,
		}
		//biz

		newFoodRating.Mask()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(newFoodRating.FakeId))
	}
}
