package restaurantlikestorage

import (
	"context"
	"lesson-5-goland/common"
	restaurantlikemodel "lesson-5-goland/modules/restaurantlike/model"
)

func (s *sqlStore) FindLikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantCreateLike) (bool, error) {
	var dataLike restaurantlikemodel.RestaurantLike

	if err := s.db.Where("restaurant_id = ?", data.RestaurantId).Where("user_id = ?", data.UserId).First(&dataLike).Error; err != nil {
		return false, common.ErrDB(err)
	}

	if dataLike.RestaurantId > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
