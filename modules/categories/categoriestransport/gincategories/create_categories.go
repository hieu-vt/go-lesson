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

func CreateCategories(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var category categoriesmodel.CreateCategories

		if err := c.ShouldBind(&category); err != nil {
			panic(err)
		}

		categoryData := categoriesmodel.Categories{
			SqlModel: common.SqlModel{
				Status: 1,
			},
			Icon:        category.Icon,
			Name:        category.Name,
			Description: category.Description,
		}

		store := categoriesstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := categoriesbiz.NewCreateCategoryBiz(store)

		if err := biz.CreateCategory(c, &categoryData); err != nil {
			panic(err)
		}

		categoryData.Mask()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(categoryData.FakeId))
	}
}
