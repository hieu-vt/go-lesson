package gincategories

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/categories/categoriesbiz"
	"lesson-5-goland/modules/categories/categoriesmodel"
	"lesson-5-goland/modules/categories/categoriesstorage"
	"net/http"
)

func UpdateCategory(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryId, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(err)
		}

		var dataUpdate categoriesmodel.Categories
		if err := c.ShouldBind(&dataUpdate); err != nil {
			panic(err)
		}

		store := categoriesstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := categoriesbiz.NewUpdateCategoryBiz(store)

		if err := biz.UpdateCategory(c, int(categoryId.GetLocalID()), &dataUpdate); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
