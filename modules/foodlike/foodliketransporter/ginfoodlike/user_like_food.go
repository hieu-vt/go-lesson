package ginlikefood

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/foodlike/foodlikebusiness"
	"lesson-5-goland/modules/foodlike/foodlikemodel"
	"lesson-5-goland/modules/foodlike/foodlikestorage"
	"net/http"
)

func UserLikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uidFood, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(err)
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var data foodlikemodel.FoodLikes

		data.FoodId = int(uidFood.GetLocalID())
		data.UserId = requester.GetUserId()

		store := foodlikestorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := foodlikebusiness.NewLikeFoodStore(store, appCtx.GetPubsub())

		if err := biz.UserLikeRestaurant(c, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]int{"ok": 1}))
	}
}
