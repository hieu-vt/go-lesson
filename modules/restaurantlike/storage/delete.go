package restaurantlikestorage

import (
	"context"
	"lesson-5-goland/common"
	restaurantlikemodel "lesson-5-goland/modules/restaurantlike/model"
)

func (s *sqlStore) DeleteLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantCreateLike) error {

	if err := s.db.Table(restaurantlikemodel.RestaurantLike{}.TableName()).
		Where("restaurant_id = ? and user_id = ?", data.RestaurantId, data.UserId).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
