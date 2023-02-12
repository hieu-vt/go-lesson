package subscriber

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/restaurantfoods/restaurantfoodsmodel"
	"lesson-5-goland/modules/restaurantfoods/restaurantfoodsstorage"
	"lesson-5-goland/pubsub"
)

func CreateRestaurantFoodsAfterCreateFood(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Create RestaurantFoods after create food",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantfoodsstorage.NewSqlStore(appCtx.GetMainDBConnection())
			data := message.Data().(restaurantfoodsmodel.CreateRestaurantFoods)

			return store.Create(ctx, &restaurantfoodsmodel.RestaurantFoods{
				SqlModel: common.SqlModel{
					Status: 1,
				},
				RestaurantId: data.RestaurantId,
				FoodId:       data.FoodId,
			})
		},
	}
}
