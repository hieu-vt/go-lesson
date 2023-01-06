package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lesson-5-goland/component"
	"lesson-5-goland/component/uploadprovider"
	"lesson-5-goland/middleware"
	"lesson-5-goland/modules/food/foodtransport/ginfood"
	"lesson-5-goland/modules/restaurant/restauranttransport/ginrestaurent"
	ginlikerestaurant "lesson-5-goland/modules/restaurantlike/transporter/gin"
	"lesson-5-goland/modules/upload/uploadtransport/ginupload"
	"lesson-5-goland/modules/user/usertransport/ginuser"
	"lesson-5-goland/pubsub/pubsublocal"
	"lesson-5-goland/reddit"
	"lesson-5-goland/skio"
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
	appCtx := component.NewAppContext(db, provider, secretKey, pubsublocal.NewPubSub(), reddit.NewRedditEngine())
	r := gin.Default()

	rtEngine := skio.NewEngine()
	if err := rtEngine.Run(appCtx, r); err != nil {
		log.Fatalln(err)
	}

	engine := subscriber.NewEngine(appCtx, rtEngine)

	engine.Start()

	//subscriber.IncreaseLikeCountAfterUserLikeRestaurant(appCtx, context.Background())
	r.Use(middleware.Recover(appCtx))
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
	}

	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

//func startSocketIOServer(engine *gin.Engine, appCtx component.AppContext) {
//	server := socketio.NewServer(&engineio.Options{
//		Transports: []transport.Transport{websocket.Default},
//	})
//
//	server.OnConnect("/", func(s socketio.Conn) error {
//		//s.SetContext("")
//		fmt.Println("connected:", s.ID(), " IP:", s.RemoteAddr())
//
//		//s.Join("Shipper")
//		//server.BroadcastToRoom("/", "Shipper", "test", "Hello 200lab")
//
//		return nil
//	})
//
//	server.OnError("/", func(s socketio.Conn, e error) {
//		fmt.Println("meet error:", e)
//	})
//
//	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
//		fmt.Println("closed", reason)
//		// Remove socket from socket engine (from app context)
//	})
//
//	//server.OnEvent("/", "authenticate", func(s socketio.Conn, token string) {
//	//
//	//	// Validate token
//	//	// If false: s.Close(), and return
//	//
//	//	// If true
//	//	// => UserId
//	//	// Fetch db find user by Id
//	//	// Here: s belongs to who? (user_id)
//	//	// We need a map[user_id][]socketio.Conn
//	//
//	//	db := appCtx.GetMainDBConnection()
//	//	store := userstorage.NewSQLStore(db)
//	//	//
//	//	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
//	//	//
//	//	payload, err := tokenProvider.Validate(token)
//	//
//	//	if err != nil {
//	//		s.Emit("authentication_failed", err.Error())
//	//		s.Close()
//	//		return
//	//	}
//	//	//
//	//	user, err := store.FindUser(context.Background(), map[string]interface{}{"id": payload.UserId})
//	//	//
//	//	if err != nil {
//	//		s.Emit("authentication_failed", err.Error())
//	//		s.Close()
//	//		return
//	//	}
//	//
//	//	if user.Status == 0 {
//	//		s.Emit("authentication_failed", errors.New("you has been banned/deleted"))
//	//		s.Close()
//	//		return
//	//	}
//	//
//	//	user.Mask(false)
//	//
//	//	s.Emit("your_profile", user)
//	//})
//
//	type Person struct {
//		Name string `json:"name"`
//		Age  int    `json:"age"`
//	}
//
//	server.OnEvent("/", "notice", func(s socketio.Conn, p Person) {
//		fmt.Println("server receive notice:", p.Name, p.Age)
//
//		p.Age = 33
//		s.Emit("notice", p)
//
//	})
//
//	server.OnEvent("/", "test", func(s socketio.Conn, msg string) {
//		fmt.Println("server receive test:", msg)
//		s.Emit("test", "Hello client")
//	})
//	//
//	//server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
//	//	s.SetContext(msg)
//	//	return "recv " + msg
//	//})
//	//
//	//server.OnEvent("/", "bye", func(s socketio.Conn) string {
//	//	last := s.Context().(string)
//	//	s.Emit("bye", last)
//	//	s.Close()
//	//	return last
//	//})
//	//
//	//server.OnEvent("/", "noteSumit", func(s socketio.Conn) string {
//	//	last := s.Context().(string)
//	//	s.Emit("bye", last)
//	//	s.Close()
//	//	return last
//	//})
//
//	go server.Serve()
//
//	engine.GET("/socket.io/*any", gin.WrapH(server))
//	engine.POST("/socket.io/*any", gin.WrapH(server))
//}
