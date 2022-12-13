package restaurantlikestorage

import (
	"context"
	"lesson-5-goland/common"
	restaurantlikemodel "lesson-5-goland/modules/restaurantlike/model"
)

func (s *sqlStore) DeleteLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantCreateLike) error {

	if err := s.db.Table(restaurantlikemodel.RestaurantLike{}.TableName()).Where("restaurant_id = ?", data.RestaurantId).Where("user_id = ?", data.UserId).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
