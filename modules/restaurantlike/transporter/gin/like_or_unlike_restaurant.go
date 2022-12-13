package ginlikerestaurant

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	restaurantlikebiz "lesson-5-goland/modules/restaurantlike/business"
	restaurantlikemodel "lesson-5-goland/modules/restaurantlike/model"
	restaurantlikestorage "lesson-5-goland/modules/restaurantlike/storage"
	"net/http"
)

func LikeOrUnlikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		type BodyData struct {
			UserId       int    `json:"user_id"`
			RestaurantId string `json:"restaurant_id"`
		}

		var body BodyData

		if err := c.ShouldBind(&body); err != nil {
			panic(err)
		}

		uidRestaurant, err := common.FromBase58(body.RestaurantId)

		if err != nil {
			panic(err)
		}

		var data restaurantlikemodel.RestaurantCreateLike

		data.RestaurantId = int(uidRestaurant.GetLocalID())
		data.UserId = body.UserId

		store := restaurantlikestorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantlikebiz.NewLikeRestaurantStore(store)

		if err := biz.CreateLikeOrUnlikeRestaurant(c, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]int{"ok": 1}))
	}
}
