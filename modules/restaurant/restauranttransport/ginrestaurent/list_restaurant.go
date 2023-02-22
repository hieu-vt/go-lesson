package ginrestaurent

import (
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/restaurant/restaurantbiz"
	"lesson-5-goland/modules/restaurant/restaurantmodel"
	"lesson-5-goland/modules/restaurant/restaurantrepository"
	"lesson-5-goland/modules/restaurant/restaurantstorage"
	restaurantlikestorage "lesson-5-goland/modules/restaurantlike/storage"
	"net/http"
)

func ListRestaurant(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
			return
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.FullFill()
		db := common.GetMainDb(sc)
		userService := sc.MustGet(common.UserApi).(restaurantrepository.UserService)
		store := restaurantstorage.NewSqlStore(db)
		likeStore := restaurantlikestorage.NewSqlStore(db)
		repository := restaurantrepository.NewListRepository(store, likeStore, userService)
		biz := restaurantbiz.NewListRestaurant(repository)
		result, err := biz.ListRestaurant(c, filter, paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)

			if paging.Limit <= len(result) {
				paging.NextCursor = result[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
