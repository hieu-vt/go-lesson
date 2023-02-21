package ginlikerestaurant

import (
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	restaurantlikebiz "lesson-5-goland/modules/restaurantlike/business"
	restaurantlikemodel "lesson-5-goland/modules/restaurantlike/model"
	restaurantlikestorage "lesson-5-goland/modules/restaurantlike/storage"
	"net/http"
)

func UserLikeRestaurant(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uidRestaurant, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(err)
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var data restaurantlikemodel.RestaurantCreateLike

		data.RestaurantId = int(uidRestaurant.GetLocalID())
		data.UserId = requester.GetUserId()

		db := common.GetMainDb(sc)

		store := restaurantlikestorage.NewSqlStore(db)
		//inCreateStore := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantlikebiz.NewLikeRestaurantStore(store, appCtx.GetPubsub())

		if err := biz.UserLikeRestaurant(c, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]int{"ok": 1}))
	}
}
