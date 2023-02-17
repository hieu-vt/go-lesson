package ginfood

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/food/foodbiz"
	"lesson-5-goland/modules/food/foodmodel"
	"lesson-5-goland/modules/food/foodstorage"
	"net/http"
)

func CreateFood(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body foodmodel.CreateFood
		c.ShouldBind(&body)

		uuidRestaurantID, err := common.FromBase58(body.RestaurantId)

		if err != nil {
			panic(err)
		}

		store := foodstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := foodbiz.NewBizCreateFood(store, appCtx.GetPubsub())
		var food foodmodel.Food

		food = foodmodel.Food{
			Name:         body.Name,
			Description:  body.Description,
			RestaurantId: int(uuidRestaurantID.GetLocalID()),
			Price:        body.Price,
			Images:       body.Images}

		if err := biz.CreateFood(
			c,
			&food,
		); err != nil {
			panic(err)
		}

		food.Mask()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(food.FakeId))
	}
}
