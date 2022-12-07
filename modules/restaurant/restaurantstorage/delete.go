package restaurantstorage

import (
	"context"
	"lesson-5-goland/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) DeleteRestaurantWithCondition(ctx context.Context, condition map[string]interface{}) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where(condition).Delete(nil).Error; err != nil {
		return err
	}

	return nil
}
