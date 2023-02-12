package foodstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/food/foodmodel"
)

func (s *sqlStore) List(ctx context.Context, paging *common.Paging, moreKeys ...string) ([]foodmodel.Food, error) {
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(foodmodel.Food{}.TableName()).Where("status in (1)")

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

	var result []foodmodel.Food

	if err := db.Limit(paging.Limit).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
