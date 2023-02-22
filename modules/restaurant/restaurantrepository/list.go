package restaurantrepository

import (
	"context"
	"errors"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/restaurant/restaurantmodel"
)

type ListRestaurantLikeStore interface {
	GetRestaurantLike(ctx context.Context, ids []int) (map[int]int, error)
}

type UserService interface {
	GetUsers(ctx context.Context, ids []int) ([]common.SimpleUser, error)
}

type ListRestaurantStore interface {
	ListRestaurantWithCondition(ctx context.Context, condition map[string]interface{},
		filter restaurantmodel.Filter,
		paging common.Paging,
		moreOptions ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantRepository struct {
	store       ListRestaurantStore
	likeStore   ListRestaurantLikeStore
	userService UserService
}

func NewListRepository(store ListRestaurantStore, likeStore ListRestaurantLikeStore, userService UserService) *listRestaurantRepository {
	return &listRestaurantRepository{
		store:       store,
		likeStore:   likeStore,
		userService: userService,
	}
}

func (repository *listRestaurantRepository) ListDataRestaurant(
	ctx context.Context,
	filter restaurantmodel.Filter,
	paging common.Paging,
) ([]restaurantmodel.Restaurant, error) {

	result, err := repository.store.ListRestaurantWithCondition(ctx, nil, filter, paging)

	if err != nil {
		return nil, err
	}

	ids := make([]int, len(result))

	for i := range result {
		ids[i] = result[i].Id
	}

	users, err := repository.userService.GetUsers(ctx, ids)

	if err != nil {
		return nil, errors.New("cannot get users by ids")
	}

	cacheUsers := make(map[int]*common.SimpleUser)

	for i, item := range users {
		cacheUsers[item.Id] = &users[i]
	}

	for i := range result {
		result[i].Owner = cacheUsers[result[i].OwnerId]
	}

	return result, nil
}
