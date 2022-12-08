package restaurantbiz

import (
	"context"
	"errors"
	"lesson-5-goland/modules/restaurant/restaurantmodel"
)

type DeleteRestaurantStore interface {
	FindByCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
	DeleteRestaurantWithCondition(ctx context.Context, id int) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{
		store: store,
	}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {

	data, err := biz.store.FindByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if data.Status <= 0 {
		return errors.New("restaurant not found")
	}

	if err := biz.store.DeleteRestaurantWithCondition(ctx, id); err != nil {
		return err
	}

	return nil
}
