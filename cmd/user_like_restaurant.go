package cmd

import (
	"context"
	goservice "github.com/200Lab-Education/go-sdk"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"lesson-5-goland/common"
	"lesson-5-goland/component/asyncjob"
	"lesson-5-goland/modules/restaurant/restaurantstorage"
	"lesson-5-goland/plugin/pubsub"
	"lesson-5-goland/plugin/pubsub/nats"
	"lesson-5-goland/plugin/sdkgorm"
)

var StartSubscribeUserLikeRestaurantCmd = &cobra.Command{
	Use:   "user-like-restaurant",
	Short: "User like restaurant",
	Run: func(cmd *cobra.Command, args []string) {
		service := goservice.New(
			goservice.WithInitRunnable(sdkgorm.NewGormDB("main", common.DBMain)),
			goservice.WithInitRunnable(nats.NewNatsPubSub(common.PluginNATS)),
		)

		if err := service.Init(); err != nil {
			log.Fatalln(err)
		}

		ps := service.MustGet(common.PluginNATS).(pubsub.NatsPubSub)

		ctx := context.Background()

		ch, _ := ps.Subscribe(ctx, common.TopicUserLikeRestaurant)

		for msg := range ch {
			db := service.MustGet(common.DBMain).(*gorm.DB)

			if restaurantId, ok := msg.Data()["restaurant_id"]; ok {
				job := asyncjob.NewJob(func(ctx context.Context) error {
					return restaurantstorage.NewSqlStore(db).InCreateLikeCount(ctx, int(restaurantId.(float64)))
				})

				if err := asyncjob.NewGroup(true, job).Run(ctx); err != nil {
					log.Println(err)
				}
			}
		}

	},
}
