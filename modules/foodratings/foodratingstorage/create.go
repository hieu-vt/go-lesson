package foodratingstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/foodratings/foodratingsmodel"
)

func (s *sqlStore) Create(ctx context.Context, data *foodratingsmodel.FoodRatings) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
