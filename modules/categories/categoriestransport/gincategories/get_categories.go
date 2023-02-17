package gincategories

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/categories/categoriesbiz"
	"lesson-5-goland/modules/categories/categoriesstorage"
	"net/http"
)

func GetCategories(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := categoriesstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := categoriesbiz.NewGetCategoriesBiz(store)

		result, err := biz.GetCategories(c)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask()
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
