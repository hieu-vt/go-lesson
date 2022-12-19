package restaurantlikebiz

import (
	"context"
	"lesson-5-goland/component/asyncjob"
	restaurantlikemodel "lesson-5-goland/modules/restaurantlike/model"
)

type UnlikeRestaurantStore interface {
	DeleteLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantCreateLike) error
	FindLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantCreateLike) (bool, error)
}

type DeCreateLikeRestaurantStore interface {
	DeCreateLikeCount(ctx context.Context, restaurantId int) error
}

type unlikeRestaurantBiz struct {
	store         UnlikeRestaurantStore
	deCreateStore DeCreateLikeRestaurantStore
}

func NewUnlikeRestaurantStore(store UnlikeRestaurantStore, deCreateStore DeCreateLikeRestaurantStore) *unlikeRestaurantBiz {
	return &unlikeRestaurantBiz{store: store, deCreateStore: deCreateStore}
}

func (biz *unlikeRestaurantBiz) UserUnlikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantCreateLike) error {
	isExist, _ := biz.store.FindLikeRestaurant(ctx, data)

	if !isExist {
		return restaurantlikemodel.ErrLikeRestaurantDidLikeThisRestaurant(nil)
	}

	if err := biz.store.DeleteLikeRestaurant(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannotUnlikeRestaurant(err)
	}

	//go func() {
	//	defer common.AppRecover()
	//	biz.deCreateStore.DeCreateLikeCount(ctx, data.RestaurantId)
	//}()

	// side effect
	job := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.deCreateStore.DeCreateLikeCount(ctx, data.RestaurantId)
	})

	asyncjob.NewGroup(true, job).Run(ctx)

	return nil
}
