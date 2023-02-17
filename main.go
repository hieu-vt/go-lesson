package main

import (
	"context"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.6.1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/component/uploadprovider/firebasestorage"
	"lesson-5-goland/component/uploadprovider/s3"
	"lesson-5-goland/middleware"
	"lesson-5-goland/modules/carts/carttransport/gincart"
	"lesson-5-goland/modules/categories/categoriestransport/gincategories"
	"lesson-5-goland/modules/food/foodtransport/ginfood"
	ginlikefood "lesson-5-goland/modules/foodlike/foodliketransporter/ginfoodlike"
	"lesson-5-goland/modules/order/ordertransport/ginorder"
	"lesson-5-goland/modules/orderdetails/orderdetailtransport/ginorderdetail"
	"lesson-5-goland/modules/ordertracking/ordertrackingtranport/ginordertracking"
	"lesson-5-goland/modules/restaurant/restauranttransport/ginrestaurent"
	ginlikerestaurant "lesson-5-goland/modules/restaurantlike/transporter/gin"
	"lesson-5-goland/modules/upload/uploadtransport/ginupload"
	"lesson-5-goland/modules/user/usertransport/ginuser"
	"lesson-5-goland/pubsub/pubsublocal"
	"lesson-5-goland/reddit"
	"lesson-5-goland/skio"
	"lesson-5-goland/subscriber"
	"net/http"
	"os"
)

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(common.TRACE_SERVICE),
			semconv.DeploymentEnvironmentKey.String(common.ENVIRONMENT),
		)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}

func main() {
	dsn := os.Getenv("DBConnectionStr")
	S3BucketName := os.Getenv("S3BucketNameStr")
	S3Region := os.Getenv("S3RegionStr")
	S3ApiKey := os.Getenv("S3ApiKeyStr")
	S3Secret := os.Getenv("S3SecretStr")
	S3Domain := os.Getenv("S3DomainStr")
	secretKey := os.Getenv("SecretKeyStr")
	projectFirebaseStorageID := os.Getenv("ProjectFirebaseStorageID")
	storageBucket := os.Getenv("StorageBucket")
	firebaseStorageDomain := os.Getenv("FirebaseStorageDomain")
	CredentialsFileName := os.Getenv("CredentialsFileName")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	s3Provider := s3.NewS3Provider(S3BucketName, S3Region, S3ApiKey, S3Secret, S3Domain)
	firebaseBucket := firebasestorage.NewFirebaseStorage(
		context.Background(),
		projectFirebaseStorageID,
		storageBucket,
		firebaseStorageDomain,
		CredentialsFileName,
	)

	if err != nil {
		log.Fatalln(err)
	}

	if error := runService(db, s3Provider, secretKey, firebaseBucket); error != nil {
		log.Fatalln(error)
	}
}

func runService(db *gorm.DB,
	provider s3.UploadProvider,
	secretKey string,
	firebaseBucket firebasestorage.UploadFirebaseStorageProvider,
) error {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	appCtx := component.NewAppContext(
		db,
		provider,
		secretKey,
		pubsublocal.NewPubSub(),
		reddit.NewRedditEngine(),
		firebaseBucket,
	)

	r := gin.Default()
	// bypass cors
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	log.WithField("text", "Hello").Info("My name is Hieu 1")
	log.Info("hello 1111")

	rtEngine := skio.NewEngine()
	if err := rtEngine.Run(appCtx, r); err != nil {
		log.Fatalln(err)
	}

	engine := subscriber.NewEngine(appCtx, rtEngine)

	engine.Start()

	_, err := tracerProvider("http://localhost:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}

	//subscriber.IncreaseLikeCountAfterUserLikeRestaurant(appCtx, context.Background())
	r.Use(middleware.Recover(appCtx))
	r.Use(otelgin.Middleware(common.TRACE_SERVICE))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// CRUD
	r.StaticFile("/demo/", "./demo.html")
	r.StaticFile("/demo/shipper", "./demoshipper.html")
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

	// food
	foods := v1.Group("/foods", middleware.RequiredAuth(appCtx))
	{
		foods.POST("", ginfood.CreateFood(appCtx))
		foods.GET("", ginfood.GetFoods(appCtx))
		// like Food
		foods.POST("/:id/like", ginlikefood.UserLikeRestaurant(appCtx))
		// unlike Food
		foods.DELETE("/:id/unlike", ginlikefood.UserUnLikeRestaurant(appCtx))
	}

	// Order
	orders := v1.Group("/orders", middleware.RequiredAuth(appCtx))
	{
		orders.POST("", ginorder.CreateOrder(appCtx))
		orders.POST("/detail", ginorderdetail.CreateOrderDetail(appCtx))
		orders.POST("/tracking", ginordertracking.CreateOrderTracking(appCtx))
		orders.GET("", ginorder.GetOrders(appCtx))
	}

	// cart
	carts := v1.Group("/carts", middleware.RequiredAuth(appCtx))
	{
		carts.POST("", gincart.AddToCart(appCtx))
		carts.PUT("/:id", gincart.UpdateCart(appCtx))
		carts.DELETE("", gincart.DeleteCart(appCtx))
		carts.GET("", gincart.ListCart(appCtx))
	}

	categories := v1.Group("/categories", middleware.RequiredAuth(appCtx))
	{
		categories.POST("", gincategories.CreateCategories(appCtx))
		categories.DELETE("/:id", gincategories.DeleteCategory(appCtx))
		categories.PUT("/:id", gincategories.UpdateCategory(appCtx))
		categories.GET("", gincategories.GetCategories(appCtx))
	}

	v1.GET("/encode-uid", func(c *gin.Context) {
		type reqData struct {
			DbType int `form:"type"`
			RealId int `form:"id"`
		}

		var d reqData
		c.ShouldBind(&d)

		c.JSON(http.StatusOK, gin.H{
			"id": common.NewUID(uint32(d.RealId), d.DbType, 1),
		})
	})

	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
