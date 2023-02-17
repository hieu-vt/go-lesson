package categoriesbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/categories/categoriesmodel"
)

type deleteCategoryStore interface {
	Delete(ctx context.Context, categoryId int) error
}

type deleteCategoryBiz struct {
	store deleteCategoryStore
}

func NewDeleteCategoryBiz(store deleteCategoryStore) *deleteCategoryBiz {
	return &deleteCategoryBiz{store: store}
}

func (biz *deleteCategoryBiz) DeleteCategory(ctx context.Context, categoryId int) error {
	if err := biz.store.Delete(ctx, categoryId); err != nil {
		return common.ErrCannotDeleteEntity(categoriesmodel.Categories{}.TableName(), err)
	}

	return nil
}
