package foodstorage

import (
	"context"
	"gorm.io/gorm"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/food/foodmodel"
)

func (s *sqlStore) Create(ctx context.Context, food *foodmodel.Food) error {
	if err := s.db.Create(&food).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) InCreateFoodLikeCount(ctx context.Context, foodId int) error {
	if err := s.db.Table(foodmodel.Food{}.TableName()).
		Where("id = ?", foodId).
		Updates(map[string]interface{}{"like_count": gorm.Expr("like_count + ?", 1)}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) DeCreateFoodLikeCount(ctx context.Context, foodId int) error {
	if err := s.db.Table(foodmodel.Food{}.TableName()).
		Where("id = ?", foodId).
		Updates(map[string]interface{}{"like_count": gorm.Expr("like_count - ?", 1)}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
