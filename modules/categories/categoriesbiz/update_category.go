package categoriesbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/categories/categoriesmodel"
)

type updateCategoryStore interface {
	Update(ctx context.Context, categoryId int, data *categoriesmodel.Categories) error
}

type updateCategoryBiz struct {
	store updateCategoryStore
}

func NewUpdateCategoryBiz(store updateCategoryStore) *updateCategoryBiz {
	return &updateCategoryBiz{store: store}
}

func (biz *updateCategoryBiz) UpdateCategory(ctx context.Context, categoryId int, data *categoriesmodel.Categories) error {
	if err := biz.store.Update(ctx, categoryId, data); err != nil {
		return common.ErrCannotUpdateEntity(categoriesmodel.Categories{}.TableName(), err)
	}

	return nil
}
