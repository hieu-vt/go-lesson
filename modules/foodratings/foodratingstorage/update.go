package foodratingstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/foodratings/foodratingsmodel"
)

func (s *sqlStore) Update(ctx context.Context, ratingId int, data *foodratingsmodel.UpdateFoodRatings) error {
	if err := s.db.Where("id = (?)", ratingId).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
