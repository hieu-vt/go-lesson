package restaurantbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/restaurant/restaurantmodel"
)

type ListRestaurantStore interface {
	ListRestaurantWithCondition(ctx context.Context, condition map[string]interface{},
		filter restaurantmodel.Filter,
		paging common.Paging,
		moreOptions ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurant(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListRestaurant(ctx context.Context,
	filter restaurantmodel.Filter,
	paging common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	return biz.store.ListRestaurantWithCondition(ctx, nil, filter, paging)
}
