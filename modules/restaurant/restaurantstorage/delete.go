package restaurantstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) DeleteRestaurantWithCondition(ctx context.Context, id int) error {
	db := s.db
	status := 0

	if err := db.Where("id = ?", id).Updates(&restaurantmodel.RestaurantDelete{Status: &status}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
