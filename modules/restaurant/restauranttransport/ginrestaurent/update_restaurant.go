package ginrestaurent

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/restaurant/restaurantbiz"
	"lesson-5-goland/modules/restaurant/restaurantmodel"
	"lesson-5-goland/modules/restaurant/restaurantstorage"
	"net/http"
	"strconv"
)

func UpdateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(401, map[string]string{
				"error": err.Error(),
			})

			return
		}

		var body restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&body); err != nil {
			c.JSON(401, map[string]string{
				"error": err.Error(),
			})

			return
		}

		store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurant(store)

		if err := biz.UpdateRestaurant(c, id, &body); err != nil {
			c.JSON(401, map[string]string{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]int{"ok": 1}))
	}
}