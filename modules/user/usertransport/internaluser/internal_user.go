package internaluser

import (
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/user/userstorage"
	"net/http"
)

func GetUserByIds(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var param struct {
			Ids []int `json:"ids"`
		}

		if err := c.ShouldBind(&param); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := sc.MustGet(common.DBMain).(*gorm.DB)
		store := userstorage.NewSqlStore(db)

		result, err := store.GetUsers(c.Request.Context(), param.Ids)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
