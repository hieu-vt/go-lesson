package restaurantstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) ListRestaurantWithCondition(ctx context.Context, condition map[string]interface{},
	filter restaurantmodel.Filter,
	paging common.Paging,
	moreOptions ...string,
) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant
	db := s.db

	for v := range moreOptions {
		db = s.db.Preload(string(v))
	}

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).Where(condition)

	if filter.CityId > 0 {
		db = db.Where("city_id = ?", filter.CityId)
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	offset := (paging.Page - 1) * paging.Limit

	err := db.Offset(offset).Limit(paging.Limit).Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}