package foodlikestorage

import (
	"context"
	"lesson-5-goland/common"
	restaurantlikemodel "lesson-5-goland/modules/restaurantlike/model"
)

func (s *sqlStore) GetFoodLike(ctx context.Context, ids []int) (map[int]int, error) {
	rLikeIds := make(map[int]int, len(ids))

	type sqlData struct {
		FoodId    int `json:"food_id" gorm:"column:food_id"`
		LikeCount int `gorm:"column:count;"`
	}

	var listData []sqlData

	if err := s.db.Table(restaurantlikemodel.RestaurantLike{}.TableName()).
		Select("food_id, count(food_id) as count").
		Where("food_id in (?)", ids).
		Group("food_id").
		Find(&listData).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listData {
		rLikeIds[item.FoodId] = item.LikeCount
	}

	return rLikeIds, nil
}
