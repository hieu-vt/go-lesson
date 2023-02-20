package ginuser

import (
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"lesson-5-goland/common"
	"lesson-5-goland/component/hasher"
	"lesson-5-goland/modules/user/userbiz"
	"lesson-5-goland/modules/user/usermodel"
	"lesson-5-goland/modules/user/userstorage"
	"lesson-5-goland/plugin/jwtprovider"
	"net/http"
)

func Login(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body usermodel.UserLogin
		//reddit := appCtx.GetReddit()
		if err := c.ShouldBind(&body); err != nil {
			panic(err)
		}
		store := userstorage.NewSqlStore(common.GetMainDb(sc))
		tokenProvider := sc.MustGet(common.JwtProvider).(jwtprovider.Provider)
		md5Hasher := hasher.NewMd5Hash()
		biz := userbiz.NewLoginBiz(store, md5Hasher, tokenProvider, 60*60*24*30, nil)
		token, err := biz.Login(c, &body)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(token))
	}
}
