package foodlikebusiness

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/foodlike/foodlikemodel"
	"lesson-5-goland/pubsub"
	"log"
)

type LikeFoodStore interface {
	CreateLikeFood(ctx context.Context, data *foodlikemodel.FoodLikes) error
	FindLikeFood(ctx context.Context, data *foodlikemodel.FoodLikes) (bool, error)
}

//type InCreateLikeRestaurantStore interface {
//	InCreateLikeCount(ctx context.Context, restaurantId int) error
//}

type likeFoodBiz struct {
	store  LikeFoodStore
	pubsub pubsub.Pubsub
}

func NewLikeFoodStore(
	store LikeFoodStore,
	pubsub pubsub.Pubsub,
) *likeFoodBiz {
	return &likeFoodBiz{
		store: store,
		//restaurantStore: restaurantStore,
		pubsub: pubsub,
	}
}

func (biz *likeFoodBiz) UserLikeRestaurant(ctx context.Context, data *foodlikemodel.FoodLikes) error {
	isExist, errFind := biz.store.FindLikeFood(ctx, data)

	if errFind != nil {
		log.Println("cannot find like before", errFind)
	}

	if isExist {
		return foodlikemodel.ErrLikeFoodExist(errFind)
	}

	if err := biz.store.CreateLikeFood(ctx, data); err != nil {
		return foodlikemodel.ErrCannotLikeFood(err)
	}

	//job := asyncjob.NewJob(func(ctx context.Context) error {
	//	return biz.restaurantStore.InCreateLikeCount(ctx, data.RestaurantId)
	//})
	//
	//group := asyncjob.NewGroup(true, job)
	//
	//go func() {
	//	defer common.AppRecover()
	//	group.Run(ctx)
	//}()

	// side effect
	//job := asyncjob.NewJob(func(ctx context.Context) error {
	//	return biz.restaurantStore.InCreateLikeCount(ctx, data.RestaurantId)
	//})
	//
	//_ = asyncjob.NewGroup(true, job).Run(ctx)

	biz.pubsub.Publish(ctx, common.TopicUserLikeFood, pubsub.NewMessage(data))

	return nil
}
