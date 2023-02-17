package ginuser

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/component/hasher"
	"lesson-5-goland/component/tokenprovider/jwt"
	"lesson-5-goland/modules/user/userbiz"
	"lesson-5-goland/modules/user/usermodel"
	"lesson-5-goland/modules/user/userstorage"
	"net/http"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body usermodel.UserLogin
		reddit := appCtx.GetReddit()
		if err := c.ShouldBind(&body); err != nil {
			panic(err)
		}

		store := userstorage.NewSqlStore(appCtx.GetMainDBConnection())
		tokenProvider := jwt.NewTokenJwt(appCtx.SecretKey())
		md5Hasher := hasher.NewMd5Hash()
		biz := userbiz.NewLoginBiz(store, md5Hasher, tokenProvider, 60*60*24*30, reddit)
		token, err := biz.Login(c, &body)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(token))
	}
}
