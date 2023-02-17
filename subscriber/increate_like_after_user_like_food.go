package subscriber

import (
	"context"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/food/foodstorage"
	"lesson-5-goland/pubsub"
)

type HasFoodId interface {
	GetFoodId() int
}

func RunIncreaseLikeCountAfterUserLikeFood(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase like count after user likes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := foodstorage.NewSqlStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasFoodId)
			id := likeData.GetFoodId()
			return store.InCreateFoodLikeCount(ctx, id)
		},
	}
}
