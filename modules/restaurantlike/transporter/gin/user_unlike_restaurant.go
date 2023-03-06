package ginlikerestaurant

import (
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	restaurantlikebiz "lesson-5-goland/modules/restaurantlike/business"
	restaurantlikemodel "lesson-5-goland/modules/restaurantlike/model"
	restaurantlikestorage "lesson-5-goland/modules/restaurantlike/storage"
	"lesson-5-goland/plugin/pubsub"
	"net/http"
)

func UserUnLikeRestaurant(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uidRestaurant, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(err)
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var data restaurantlikemodel.RestaurantCreateLike

		data.RestaurantId = int(uidRestaurant.GetLocalID())
		data.UserId = requester.GetUserId()

		store := restaurantlikestorage.NewSqlStore(common.GetMainDb(sc))
		//deCreateLikeRestaurant := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
		pb := sc.MustGet(common.PluginNATS).(pubsub.NatsPubSub)
		biz := restaurantlikebiz.NewUnlikeRestaurantStore(store, pb)

		if err := biz.UserUnlikeRestaurant(c, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]int{"ok": 1}))
	}
}
