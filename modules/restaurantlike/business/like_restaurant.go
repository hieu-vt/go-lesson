package restaurantlikebiz

import (
	"context"
	"lesson-5-goland/component/asyncjob"
	restaurantlikemodel "lesson-5-goland/modules/restaurantlike/model"
	"log"
)

type LikeRestaurantStore interface {
	CreateLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantCreateLike) error
	FindLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantCreateLike) (bool, error)
}

type InCreateLikeRestaurantStore interface {
	InCreateLikeCount(ctx context.Context, restaurantId int) error
}

type likeRestaurantBiz struct {
	store LikeRestaurantStore

	restaurantStore InCreateLikeRestaurantStore
}

func NewLikeRestaurantStore(store LikeRestaurantStore, restaurantStore InCreateLikeRestaurantStore) *likeRestaurantBiz {
	return &likeRestaurantBiz{store: store, restaurantStore: restaurantStore}
}

func (biz *likeRestaurantBiz) UserLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantCreateLike) error {
	isExist, errFind := biz.store.FindLikeRestaurant(ctx, data)

	if errFind != nil {
		log.Println("cannot find like before", errFind)
	}

	if isExist {
		return restaurantlikemodel.ErrLikeRestaurantExist(errFind)
	}

	if err := biz.store.CreateLikeRestaurant(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
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
	job := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.restaurantStore.InCreateLikeCount(ctx, data.RestaurantId)
	})

	_ = asyncjob.NewGroup(true, job).Run(ctx)

	return nil
}
