package restaurantrepository

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

type listRestaurantRepository struct {
	store     ListRestaurantStore
	likeStore ListRestaurantLikeStore
}

func NewListRepository(store ListRestaurantStore, likeStore ListRestaurantLikeStore) *listRestaurantRepository {
	return &listRestaurantRepository{
		store:     store,
		likeStore: likeStore,
	}
}

func (repository *listRestaurantRepository) ListDataRestaurant(ctx context.Context,
	filter restaurantmodel.Filter,
	paging common.Paging) ([]restaurantmodel.Restaurant, error) {

	result, err := repository.store.ListRestaurantWithCondition(ctx, nil, filter, paging)

	if err != nil {
		return nil, err
	}

	ids := make([]int, len(result))

	for i := range result {
		ids[i] = result[i].Id
	}

	rLikeIds, err := repository.likeStore.GetRestaurantLike(ctx, ids)

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
