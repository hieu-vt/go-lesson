package restaurantbiz

import (
	"context"
	"errors"
	"lesson-5-goland/modules/restaurant/restaurantmodel"
)

type UpdateRestaurantStore interface {
	FindByCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
	UpdateData(ctx context.Context, id int, body *restaurantmodel.RestaurantUpdate) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurant(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{
		store: store,
	}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, id int, body *restaurantmodel.RestaurantUpdate) error {
	data, err := biz.store.FindByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if data.Status <= 0 {
		return errors.New("restaurant not active")
	}

	if err := biz.store.UpdateData(ctx, id, body); err != nil {
		return err
	}

	return nil
}
