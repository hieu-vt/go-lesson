package restaurantbiz

import (
	"context"
	"lesson-5-goland/modules/restaurant/restaurantmodel"
)

type CreateRestaurantStore interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := data.ValidateRestaurantData(); err != nil {
		return err
	}

	err := biz.store.Create(ctx, data)

	return err
}
