package restaurantbiz

import (
	"context"
	"lesson-5-goland/modules/restaurant/restaurantmodel"
)

type GetRestaurantStore interface {
	FindByCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
}

type getRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetRestaurantBiz(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz *getRestaurantBiz) GetRestaurantById(ctx context.Context, id interface{}) (*restaurantmodel.Restaurant, error) {
	result, err := biz.store.FindByCondition(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
