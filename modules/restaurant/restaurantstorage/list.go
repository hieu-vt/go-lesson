package restaurantstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) ListRestaurantWithCondition(ctx context.Context, condition map[string]interface{},
	filter restaurantmodel.Filter,
	paging common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant
	db := s.db

	for i := range moreKeys {
		db = s.db.Preload(moreKeys[i])
	}

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).Order("id desc").Where(condition).Where("status in (1)")

	if filter.CityId > 0 {
		db = db.Where("city_id = ?", filter.CityId)
	}

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

	err := db.Limit(paging.Limit).Find(&result).Error

	if err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
