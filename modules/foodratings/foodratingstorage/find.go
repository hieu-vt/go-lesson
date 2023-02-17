package foodratingstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/foodratings/foodratingsmodel"
)

func (s *sqlStore) Find(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (foodratingsmodel.FoodRatings, error) {
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var result foodratingsmodel.FoodRatings
	if err := db.Where(condition).First(&result).Error; err != nil {
		return result, common.ErrDB(err)
	}

	return result, nil
}
