package foodstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/food/foodmodel"
)

func (s *sqlStore) Create(ctx context.Context, food *foodmodel.Food) error {
	if err := s.db.Create(&food).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
