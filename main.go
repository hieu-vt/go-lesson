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
	"lesson-5-goland/pubsub/pubsublocal"
	"lesson-5-goland/subscriber"
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
	secretKey := os.Getenv("SecretKeyStr")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	s3Provider := uploadprovider.NewS3Provider(S3BucketName, S3Region, S3ApiKey, S3Secret, S3Domain)

	if err != nil {
		log.Fatalln(err)
	}

	if error := runService(db, s3Provider, secretKey); error != nil {
		log.Fatalln(error)
	}
}

func runService(db *gorm.DB, provider uploadprovider.UploadProvider, secretKey string) error {
	appCtx := component.NewAppContext(db, provider, secretKey, pubsublocal.NewPubSub())
	r := gin.Default()
	engine := subscriber.NewEngine(appCtx)

	engine.Start()

	//subscriber.IncreaseLikeCountAfterUserLikeRestaurant(appCtx, context.Background())
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
	v1.POST("/login", ginuser.Login(appCtx))
	v1.GET("/profile", middleware.RequiredAuth(appCtx), ginuser.GetProfile(appCtx))

	// upload
	v1.POST("/upload", ginupload.UploadFile(appCtx))

	// restaurant
	restaurants := v1.Group("/restaurants", middleware.RequiredAuth(appCtx))
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

		// like Restaurant
		restaurants.POST("/:id/like", ginlikerestaurant.UserLikeRestaurant(appCtx))

		// unlike Restaurant
		restaurants.DELETE("/:id/unlike", ginlikerestaurant.UserUnLikeRestaurant(appCtx))
	}

	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
