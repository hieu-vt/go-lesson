package restaurantfoodsstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/restaurantfoods/restaurantfoodsmodel"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantfoodsmodel.RestaurantFoods) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
