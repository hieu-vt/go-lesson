package subscriber

import (
	"context"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/restaurant/restaurantstorage"
	"lesson-5-goland/pubsub"
)

func DeCreateLikeCountAfterUserDislikeRestaurant(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "DeCreate like count after user dislike restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
			data := message.Data().(HasRestaurantId)

			return store.DeCreateLikeCount(ctx, data.GetRestaurantId())
		},
	}
}
