package foodlikestorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/foodlike/foodlikemodel"
)

func (s *sqlStore) DeleteLikeFood(ctx context.Context, data *foodlikemodel.FoodLikes) error {

	if err := s.db.Table(foodlikemodel.FoodLikes{}.TableName()).
		Where("food_id = ? and user_id = ?", data.FoodId, data.UserId).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
