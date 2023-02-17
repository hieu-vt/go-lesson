package categoriesstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/categories/categoriesmodel"
)

func (s *sqlStore) Find(ctx context.Context, condition map[string]interface{}, moreKeys ...string) ([]categoriesmodel.Categories, error) {
	db := s.db
	db = db.Table(categoriesmodel.Categories{}.TableName())

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var result []categoriesmodel.Categories

	if err := db.Where(condition).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
