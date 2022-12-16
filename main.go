package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lesson-5-goland/component"
	"lesson-5-goland/component/uploadprovider"
	"lesson-5-goland/middleware"
	"lesson-5-goland/modules/restaurant/restauranttransport/ginrestaurent"
	ginlikerestaurant "lesson-5-goland/modules/restaurantlike/transporter/gin"
	"lesson-5-goland/modules/upload/uploadtransport/ginupload"
	"lesson-5-goland/modules/user/usertransport/ginuser"
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
	S3BucketName := os.Getenv("S3BucketNameStr")
	S3Region := os.Getenv("S3RegionStr")
	S3ApiKey := os.Getenv("S3ApiKeyStr")
	S3Secret := os.Getenv("S3SecretStr")
	S3Domain := os.Getenv("S3DomainStr")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	s3Provider := uploadprovider.NewS3Provider(S3BucketName, S3Region, S3ApiKey, S3Secret, S3Domain)

	if err != nil {
		log.Fatalln(err)
	}

	if error := runService(db, s3Provider); error != nil {
		log.Fatalln(error)
	}
}

func runService(db *gorm.DB, provider uploadprovider.UploadProvider) error {
	appCtx := component.NewAppContext(db, provider)
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// CRUD
	v1 := r.Group("/v1")

	// authorized
	v1.POST("/register", ginuser.Register(appCtx))

	// upload
	v1.POST("/upload", ginupload.UploadFile(appCtx))

	// restaurant
	restaurants := v1.Group("/restaurants")
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

	// like Restaurant
	likeRestaurants := v1.Group("/like")
	{
		likeRestaurants.POST("", ginlikerestaurant.LikeOrUnlikeRestaurant(appCtx))
	}

	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
