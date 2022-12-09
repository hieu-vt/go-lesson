package restaurantstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) UpdateData(ctx context.Context, id int, body *restaurantmodel.RestaurantUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(&body).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
