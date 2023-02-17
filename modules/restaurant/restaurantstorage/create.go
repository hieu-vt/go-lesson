package restaurantstorage

import (
	"context"
	"gorm.io/gorm"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	db := s.db

	if err := db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) InCreateLikeCount(ctx context.Context, restaurantId int) error {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", restaurantId).
		Updates(map[string]interface{}{"like_count": gorm.Expr("like_count + ?", 1)}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) DeCreateLikeCount(ctx context.Context, restaurantId int) error {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", restaurantId).
		Updates(map[string]interface{}{"like_count": gorm.Expr("like_count - ?", 1)}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
