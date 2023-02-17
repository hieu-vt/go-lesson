package foodlikestorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/foodlike/foodlikemodel"
)

func (s *sqlStore) CreateLikeFood(ctx context.Context, data *foodlikemodel.FoodLikes) error {

	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
