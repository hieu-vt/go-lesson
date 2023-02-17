package categoriesbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/categories/categoriesmodel"
)

type createCategoryStore interface {
	Create(ctx context.Context, data *categoriesmodel.Categories) error
}

type createCategoryBiz struct {
	store createCategoryStore
}

func NewCreateCategoryBiz(store createCategoryStore) *createCategoryBiz {
	return &createCategoryBiz{store: store}
}

func (biz *createCategoryBiz) CreateCategory(ctx context.Context, data *categoriesmodel.Categories) error {
	if err := data.Validate(); err != nil {
		return common.ErrNoPermission(err)
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(categoriesmodel.Categories{}.TableName(), err)
	}

	return nil
}
