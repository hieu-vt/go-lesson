package restaurantbiz

import "context"

type DeleteRestaurantStore interface {
	DeleteRestaurantWithCondition(ctx context.Context, condition map[string]interface{}) error
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
	err := biz.store.DeleteRestaurantWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	return nil
}
