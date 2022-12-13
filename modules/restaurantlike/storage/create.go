package restaurantlikestorage

import (
	"context"
	"lesson-5-goland/common"
	restaurantlikemodel "lesson-5-goland/modules/restaurantlike/model"
)

func (s *sqlStore) CreateLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantCreateLike) error {

	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
