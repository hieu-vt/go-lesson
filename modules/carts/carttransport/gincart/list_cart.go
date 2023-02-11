package gincart

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/carts/cartbiz"
	"lesson-5-goland/modules/carts/cartstorage"
	"net/http"
)

func ListCart(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}

		paging.FullFill()

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := cartstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := cartbiz.NewListCartBiz(store)

		result, err := biz.ListCart(c, int(requester.GetUserId()), paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask()

			if paging.Limit <= len(result) {
				paging.NextCursor = result[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
