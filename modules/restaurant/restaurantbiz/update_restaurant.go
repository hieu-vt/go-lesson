package restaurantbiz

import (
	"context"
	"lesson-5-goland/modules/restaurant/restaurantmodel"
)

type UpdateRestaurantStore interface {
	UpdateData(ctx context.Context, condition map[string]interface{}, body *restaurantmodel.RestaurantUpdate) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurant(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{
		store: store,
	}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, id interface{}, body *restaurantmodel.RestaurantUpdate) error {
	if err := biz.store.UpdateData(ctx, map[string]interface{}{"id": id}, body); err != nil {
		return err
	}

	return nil
}
