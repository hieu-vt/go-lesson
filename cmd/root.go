package cmd

import (
	"fmt"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"lesson-5-goland/cmd/handlers"
	"lesson-5-goland/common"
	"lesson-5-goland/middleware"
	"lesson-5-goland/modules/user/repository/grpcrepository"
	"lesson-5-goland/modules/user/userstorage"
	"lesson-5-goland/plugin/appredis"
	"lesson-5-goland/plugin/jwtprovider/jwt"
	"lesson-5-goland/plugin/pubsub/nats"
	"lesson-5-goland/plugin/remoteapi/appgrpc"
	"lesson-5-goland/plugin/remoteapi/restfull"
	sdkgorm2 "lesson-5-goland/plugin/sdkgorm"
	user "lesson-5-goland/proto/userproto"
	"net/http"
	"os"
)

func newService() goservice.Service {
	service := goservice.New(
		goservice.WithName("food-delivery"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(sdkgorm2.NewGormDB("main", common.DBMain)),
		goservice.WithInitRunnable(jwt.NewTokenJwtProvider(common.JwtProvider)),
		goservice.WithInitRunnable(restfull.NewUserApi(common.UserApi)),
		goservice.WithInitRunnable(nats.NewNatsPubSub(common.PluginNATS)),
		goservice.WithInitRunnable(appredis.NewAppRedis("main-redis", common.PluginAppRedis)),
		goservice.WithInitRunnable(appgrpc.NewGrpcServer(common.PluginGRPCServer)),
		goservice.WithInitRunnable(appgrpc.NewUserClient(common.PluginGrpcUserClient)),
	)

	return service
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start an food delivery service",
	Run: func(cmd *cobra.Command, args []string) {
		service := newService()
		serviceLogger := service.Logger("service")

		if err := service.Init(); err != nil {
			serviceLogger.Fatalln(err)
		}

		service.MustGet(common.PluginGRPCServer).(interface {
			SetGrpcHandler(grpcHandler func(*grpc.Server))
		}).SetGrpcHandler(func(gs *grpc.Server) {
			db := service.MustGet(common.DBMain).(*gorm.DB)
			user.RegisterUserServiceServer(gs, grpcrepository.NewGrpcUserRepository(userstorage.NewSqlStore(db)))
		})

		service.HTTPServer().AddHandler(func(engine *gin.Engine) {
			engine.Use(middleware.Recover())

			engine.GET("/ping", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"data": "pong"})
			})

			handlers.MainRoute(engine, service)
			handlers.UserServiceRoute(engine, service)
		})

		if err := service.Start(); err != nil {
			serviceLogger.Fatalln(err)
		}
	},
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)

	rootCmd.AddCommand(StartSubscribeUserLikeRestaurantCmd)
	rootCmd.AddCommand(StartSubscribeUserDislikeRestaurantCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}
