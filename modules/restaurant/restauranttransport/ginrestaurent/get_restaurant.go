package ginrestaurent

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/restaurant/restaurantbiz"
	"lesson-5-goland/modules/restaurant/restaurantstorage"
	"net/http"
	"strconv"
)

func GetRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)

		result, err := biz.GetRestaurantById(c, id)

		if err != nil {
			c.JSON(401, err)

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
