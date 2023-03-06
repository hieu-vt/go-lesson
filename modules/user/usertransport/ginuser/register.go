package ginuser

import (
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component/hasher"
	"lesson-5-goland/modules/user/userbiz"
	"lesson-5-goland/modules/user/usermodel"
	"lesson-5-goland/modules/user/userstorage"
	"net/http"
)

func Register(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSqlStore(common.GetMainDb(sc))
		hash := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBiz(store, hash)

		user, err := biz.Register(c, &data)

		if err != nil {
			panic(err)
		}

		user.GenUID(common.DbTypeUser)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(user.FakeId.String()))
	}
}
