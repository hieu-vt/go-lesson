package restaurantstorage

import (
	"context"
	"lesson-5-goland/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) UpdateData(ctx context.Context, condition map[string]interface{}, body *restaurantmodel.RestaurantUpdate) error {
	db := s.db

	if err := db.Table(restaurantmodel.RestaurantUpdate{}.TableName()).Where(condition).Updates(&body).Error; err != nil {
		return err
	}

	return nil
}