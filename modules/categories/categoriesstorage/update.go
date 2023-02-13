package categoriesstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/categories/categoriesmodel"
)

func (s *sqlStore) Update(ctx context.Context, categoryId int, data *categoriesmodel.Categories) error {
	if err := s.db.Where("id = (?)", categoryId).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
