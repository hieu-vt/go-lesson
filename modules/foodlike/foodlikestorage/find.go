package foodlikestorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/foodlike/foodlikemodel"
)

func (s *sqlStore) FindLikeFood(ctx context.Context, data *foodlikemodel.FoodLikes) (bool, error) {
	var dataLike foodlikemodel.FoodLikes

	if err := s.db.Where("food_id = ?", data.FoodId).Where("user_id = ?", data.UserId).First(&dataLike).Error; err != nil {
		return false, common.ErrDB(err)
	}

	if dataLike.FoodId > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
