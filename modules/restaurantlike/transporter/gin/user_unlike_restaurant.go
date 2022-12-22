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

func UserUnLikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uidRestaurant, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(err)
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var data restaurantlikemodel.RestaurantCreateLike

		data.RestaurantId = int(uidRestaurant.GetLocalID())
		data.UserId = requester.GetUserId()

		store := restaurantlikestorage.NewSqlStore(appCtx.GetMainDBConnection())
		//deCreateLikeRestaurant := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantlikebiz.NewUnlikeRestaurantStore(store, appCtx.GetPubsub())

		if err := biz.UserUnlikeRestaurant(c, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]int{"ok": 1}))
	}
}
