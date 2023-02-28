package handlers

import (
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"lesson-5-goland/common"
	"lesson-5-goland/middleware"
	"lesson-5-goland/modules/restaurant/restauranttransport/ginrestaurent"
	ginlikerestaurant "lesson-5-goland/modules/restaurantlike/transporter/gin"
	"lesson-5-goland/modules/user/userstorage"
	"lesson-5-goland/modules/user/usertransport/ginuser"
	"net/http"
)

func MainRoute(router *gin.Engine, sc goservice.ServiceContext) {
	dbConn := sc.MustGet(common.DBMain).(*gorm.DB)
	userStore := userstorage.NewSqlStore(dbConn)

	v1 := router.Group("/v1")
	{
		v1.GET("/admin",
			middleware.RequiredAuth(sc, userStore),
			//middleware.RequiredRoles(sc, "admin", "mod"),
			func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"data": 1})
			})

		v1.POST("/register", ginuser.Register(sc))
		v1.POST("/login", ginuser.Login(sc))
		v1.GET("/profile", middleware.RequiredAuth(sc, userStore), ginuser.GetProfile(sc))
		//
		restaurants := v1.Group("/restaurants")
		{
			restaurants.POST("", ginrestaurent.CreateRestaurant(sc))
			restaurants.GET("", ginrestaurent.ListRestaurant(sc))
			restaurants.GET("/:id", ginrestaurent.GetRestaurant(sc))
			restaurants.PATCH("/:id", ginrestaurent.UpdateRestaurant(sc))
			restaurants.DELETE("/:id", ginrestaurent.DeleteRestaurant(sc))
			restaurants.POST("/:id/like", middleware.RequiredAuth(sc, userStore), ginlikerestaurant.UserLikeRestaurant(sc))
			restaurants.DELETE("/:id/unlike", middleware.RequiredAuth(sc, userStore), ginlikerestaurant.UserUnLikeRestaurant(sc))
		}
	}
}
