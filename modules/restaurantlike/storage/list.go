package restaurantlikestorage

import "context"

func (*sqlStore) GetRestaurantLike(ctx context.Context, ids []int) (map[int]int, error) {
	rLikeIds := make(map[int]int, len(ids))

	for i, item := range ids {
		rLikeIds[item] = i
	}

	return rLikeIds, nil
}
