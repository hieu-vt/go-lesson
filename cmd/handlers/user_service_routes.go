package handlers

import (
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"lesson-5-goland/modules/user/usertransport/internaluser"
)

func UserServiceRoute(router *gin.Engine, sc goservice.ServiceContext) {
	internal := router.Group("/internal")
	{
		internal.POST("/get-users-by-ids", internaluser.GetUserByIds(sc))
	}
}
