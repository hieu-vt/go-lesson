package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lesson-5-goland/component"
	"lesson-5-goland/middleware"
	"lesson-5-goland/modules/restaurant/restauranttransport/ginrestaurent"
	"log"
	"net/http"
	"os"
)

type Restaurant struct {
	Id   int    `json:"id,omitempty" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr" gorm:"column:addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func main() {
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if error := runService(db); error != nil {
		log.Fatalln(error)
	}
}

func runService(db *gorm.DB) error {
	appCtx := component.NewAppContext(db)
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// CRUD
	restaurants := r.Group("/restaurants")
	{
		// create Restaurant
		restaurants.POST("", ginrestaurent.CreateRestaurant(appCtx))

		// Get By id
		restaurants.GET("/:id", ginrestaurent.GetRestaurant(appCtx))

		// Get restaurants
		restaurants.GET("/", ginrestaurent.ListRestaurant(appCtx))

		// Update Restaurant
		restaurants.PATCH("/:id", ginrestaurent.UpdateRestaurant(appCtx))

		// Delete Restaurant
		restaurants.DELETE("/:id", ginrestaurent.DeleteRestaurant(appCtx))
	}

	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
