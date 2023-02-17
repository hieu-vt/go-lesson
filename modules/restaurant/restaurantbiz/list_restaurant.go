package restaurantbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/restaurant/restaurantmodel"
)

type ListRestaurantRepository interface {
	ListDataRestaurant(ctx context.Context,
		filter restaurantmodel.Filter,
		paging common.Paging) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	repository ListRestaurantRepository
}

func NewListRestaurant(repository ListRestaurantRepository) *listRestaurantBiz {
	return &listRestaurantBiz{repository: repository}
}

func (biz *listRestaurantBiz) ListRestaurant(ctx context.Context,
	filter restaurantmodel.Filter,
	paging common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.repository.ListDataRestaurant(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	return result, nil
}
