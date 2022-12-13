package restaurantbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/restaurant/restaurantmodel"
	"log"
)

type ListRestaurantLikeStore interface {
	GetRestaurantLike(ctx context.Context, ids []int) (map[int]int, error)
}

type ListRestaurantStore interface {
	ListRestaurantWithCondition(ctx context.Context, condition map[string]interface{},
		filter restaurantmodel.Filter,
		paging common.Paging,
		moreOptions ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	store     ListRestaurantStore
	likeStore ListRestaurantLikeStore
}

func NewListRestaurant(store ListRestaurantStore, likeStore ListRestaurantLikeStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store, likeStore: likeStore}
}

func (biz *listRestaurantBiz) ListRestaurant(ctx context.Context,
	filter restaurantmodel.Filter,
	paging common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListRestaurantWithCondition(ctx, nil, filter, paging)

	if err != nil {
		return nil, err
	}

	ids := make([]int, len(result))

	for i := range result {
		ids[i] = result[i].Id
	}

	rLikeIds, err := biz.likeStore.GetRestaurantLike(ctx, ids)

	if err != nil {
		log.Println("cannot get likes restaurant")
	}

	if v := rLikeIds; v != nil {
		for i, item := range result {
			result[i].LikeCount = rLikeIds[item.Id]
		}
	}

	return result, nil
}
