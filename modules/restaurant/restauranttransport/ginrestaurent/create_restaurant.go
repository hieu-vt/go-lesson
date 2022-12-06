package ginrestaurent

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"lesson-5-goland/modules/restaurant/restaurantbiz"
	"lesson-5-goland/modules/restaurant/restaurantmodel"
	"lesson-5-goland/modules/restaurant/restaurantstorage"
	"net/http"
)

func CreateRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, map[string]string{
				"error": err.Error(),
			})

			return
		}

		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		err := biz.CreateRestaurant(c, &data)

		if err != nil {
			c.JSON(401, map[string]string{
				"error": err.Error(),
			})

		}

		c.JSON(http.StatusOK, data)
	}

}
