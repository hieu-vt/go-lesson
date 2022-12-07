package main

import (
	"github.com/gin-gonic/gin"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/restaurant/restauranttransport/ginrestaurent"
	"log"
	"net/http"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	//os.Setenv("DBConnectionStr", "root:ead8686ba57479778a76e@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local")
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
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// CRUD
	appCtx := component.NewAppContext(db)
	restaurants := r.Group("/restaurants")
	{
		// create Restaurant
		restaurants.POST("", ginrestaurent.CreateRestaurant(appCtx))

		// Get By Id
		restaurants.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(401, map[string]string{
					"error": err.Error(),
				})

				return
			}

			var data Restaurant

			if err := db.Where("id = ?", id).First(&data).Error; err != nil {
				c.JSON(401, map[string]string{
					"error": err.Error(),
				})

				return
			}

			c.JSON(http.StatusOK, data)
		})

		// Get restaurants
		restaurants.GET("/", ginrestaurent.ListRestaurant(appCtx))

		// Update Restaurant
		restaurants.PATCH("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(401, map[string]string{
					"error": err.Error(),
				})

				return
			}

			var body RestaurantUpdate

			if err := c.ShouldBind(&body); err != nil {
				c.JSON(401, map[string]string{
					"error": err.Error(),
				})

				return
			}

			if err := db.Table(RestaurantUpdate{}.TableName()).Where("id = ?", id).Updates(&body).Error; err != nil {
				c.JSON(401, map[string]string{
					"error": err.Error(),
				})

				return
			}

			c.JSON(http.StatusOK, body)
		})

		// Delete Restaurant
		restaurants.DELETE("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(401, map[string]string{
					"error": err.Error(),
				})

				return
			}

			if err := db.Table(Restaurant{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
				c.JSON(401, map[string]string{
					"error": err.Error(),
				})

				return
			}

			c.JSON(http.StatusOK, map[string]int{
				"ok": 1,
			})
		})
	}

	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
