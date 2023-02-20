package handlers

import (
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"lesson-5-goland/common"
	"lesson-5-goland/middleware"
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
		//restaurants := v1.Group("/restaurants")
		//{
		//	restaurants.POST("", restaurantgin.CreateRestaurantHandler(sc))
		//	restaurants.GET("", restaurantgin.ListRestaurant(sc))
		//	restaurants.GET("/:restaurant_id", restaurantgin.GetRestaurantHandler(sc))
		//	restaurants.PUT("/:restaurant_id", restaurantgin.UpdateRestaurantHandler(sc))
		//	restaurants.DELETE("/:restaurant_id", restaurantgin.DeleteRestaurantHandler(sc))
		//}
	}
}
