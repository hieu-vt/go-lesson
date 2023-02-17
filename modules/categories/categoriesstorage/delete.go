package categoriesstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/categories/categoriesmodel"
)

func (s *sqlStore) Delete(ctx context.Context, categoryId int) error {
	if err := s.db.Table(categoriesmodel.Categories{}.TableName()).Where("id = (?)", categoryId).Update("status", 0).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
