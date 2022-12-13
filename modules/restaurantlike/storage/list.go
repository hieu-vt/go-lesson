package restaurantlikestorage

import (
	"context"
	"lesson-5-goland/common"
	restaurantlikemodel "lesson-5-goland/modules/restaurantlike/model"
)

func (s *sqlStore) GetRestaurantLike(ctx context.Context, ids []int) (map[int]int, error) {
	rLikeIds := make(map[int]int, len(ids))

	type sqlData struct {
		RestaurantId int `json:"restaurant_id" gorm:"column:restaurant_id"`
		LikeCount    int `gorm:"column:count;"`
	}

	var listData []sqlData

	if err := s.db.Table(restaurantlikemodel.RestaurantLike{}.TableName()).
		Select("restaurant_id, count(restaurant_id) as count").
		Where("restaurant_id in (?)", ids).
		Group("restaurant_id").
		Find(&listData).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listData {
		rLikeIds[item.RestaurantId] = item.LikeCount
	}

	return rLikeIds, nil
}
