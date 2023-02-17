package foodratingstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/foodratings/foodratingsmodel"
)

func (s *sqlStore) List(ctx context.Context, foodId int, paging *common.Paging) ([]foodratingsmodel.FoodRatings, error) {
	db := s.db

	db = db.Table(foodratingsmodel.FoodRatings{}.TableName()).Where("status = (1)").Where("food_id = ?", foodId)

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if paging.FakeCursor != "" {
		if uid, err := common.FromBase58(paging.FakeCursor); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		} else {
			return nil, common.ErrDB(err)
		}
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	var result []foodratingsmodel.FoodRatings

	if err := db.Limit(paging.Limit).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
