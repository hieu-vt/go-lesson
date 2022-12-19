package restaurantlikebiz

import (
	"context"
	"lesson-5-goland/common"
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

	go func() {
		defer common.AppRecover()
		biz.deCreateStore.DeCreateLikeCount(ctx, data.RestaurantId)
	}()

	return nil
}
