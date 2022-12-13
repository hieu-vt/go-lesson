package restaurantlikebiz

import (
	"context"
	"lesson-5-goland/common"
	restaurantlikemodel "lesson-5-goland/modules/restaurantlike/model"
	"log"
)

type LikeRestaurantStore interface {
	CreateLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantCreateLike) error
	FindLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantCreateLike) (bool, error)
	DeleteLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantCreateLike) error
}

type likeRestaurantBiz struct {
	store LikeRestaurantStore
}

func NewLikeRestaurantStore(store LikeRestaurantStore) *likeRestaurantBiz {
	return &likeRestaurantBiz{store: store}
}

func (biz *likeRestaurantBiz) CreateLikeOrUnlikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantCreateLike) error {
	isExist, errFind := biz.store.FindLikeRestaurant(ctx, data)

	if errFind != nil {
		log.Println("cannot find like before", errFind)
	}

	if !isExist {
		err := biz.store.CreateLikeRestaurant(ctx, data)

		if err != nil {
			return common.ErrInvalidRequest(err)
		}
	} else {
		err := biz.store.DeleteLikeRestaurant(ctx, data)

		if err != nil {
			return common.ErrInvalidRequest(err)
		}
	}

	return nil
}
