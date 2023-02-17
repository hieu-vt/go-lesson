package foodlikebusiness

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/foodlike/foodlikemodel"
	"lesson-5-goland/pubsub"
)

type unlikeFoodStore interface {
	DeleteLikeFood(ctx context.Context, data *foodlikemodel.FoodLikes) error
	FindLikeFood(ctx context.Context, data *foodlikemodel.FoodLikes) (bool, error)
}

//type DeCreateLikeRestaurantStore interface {
//	DeCreateLikeCount(ctx context.Context, restaurantId int) error
//}

type unlikeFoodBiz struct {
	store unlikeFoodStore
	//deCreateStore DeCreateLikeRestaurantStore
	pubsub pubsub.Pubsub
}

func NewUnlikeFoodStore(
	store unlikeFoodStore,
	//deCreateStore DeCreateLikeRestaurantStore,
	pubsub pubsub.Pubsub,
) *unlikeFoodBiz {
	return &unlikeFoodBiz{
		store: store,
		//deCreateStore: deCreateStore,
		pubsub: pubsub,
	}
}

func (biz *unlikeFoodBiz) UserUnlikeRestaurant(ctx context.Context, data *foodlikemodel.FoodLikes) error {
	isExist, _ := biz.store.FindLikeFood(ctx, data)

	if !isExist {
		return foodlikemodel.ErrLikeFoodExist(nil)
	}

	if err := biz.store.DeleteLikeFood(ctx, data); err != nil {
		return foodlikemodel.ErrCannotUnlikeFood(err)
	}

	//go func() {
	//	defer common.AppRecover()
	//	biz.deCreateStore.DeCreateLikeCount(ctx, data.RestaurantId)
	//}()

	// side effect
	//job := asyncjob.NewJob(func(ctx context.Context) error {
	//	return biz.deCreateStore.DeCreateLikeCount(ctx, data.RestaurantId)
	//})
	//
	//asyncjob.NewGroup(true, job).Run(ctx)
	biz.pubsub.Publish(ctx, common.TopicUserDislikeFood, pubsub.NewMessage(data))

	return nil
}
