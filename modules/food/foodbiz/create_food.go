package foodbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/food/foodmodel"
	"lesson-5-goland/modules/restaurantfoods/restaurantfoodsmodel"
	"lesson-5-goland/pubsub"
)

type CreateFoodStore interface {
	Create(ctx context.Context, food *foodmodel.Food) error
}

type createFoodBiz struct {
	store  CreateFoodStore
	pubSub pubsub.Pubsub
}

func NewBizCreateFood(store CreateFoodStore, pubSub pubsub.Pubsub) *createFoodBiz {
	return &createFoodBiz{store: store, pubSub: pubSub}
}

func (biz *createFoodBiz) CreateFood(ctx context.Context, data *foodmodel.Food) error {
	data.Status = 1

	err := biz.store.Create(ctx, data)

	if err != nil {
		return common.ErrCannotCreateEntity(foodmodel.Food{}.TableName(), err)
	}

	biz.pubSub.Publish(ctx, common.TopicCreateRestaurantFoodsAfterCreateFood, pubsub.NewMessage(restaurantfoodsmodel.CreateRestaurantFoods{
		RestaurantId: data.RestaurantId,
		FoodId:       data.Id,
	}))

	return nil
}
