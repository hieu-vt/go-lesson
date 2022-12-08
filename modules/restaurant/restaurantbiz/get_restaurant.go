package restaurantbiz

import (
	"context"
	"errors"
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

	if result.Status == 0 {
		return nil, errors.New("data deleted")
	}

	return result, nil
}
