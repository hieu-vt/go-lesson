package categoriesbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/categories/categoriesmodel"
)

type getCategoriesStore interface {
	Find(ctx context.Context, condition map[string]interface{}, moreKeys ...string) ([]categoriesmodel.Categories, error)
}

type getCategoriesBiz struct {
	store getCategoriesStore
}

func NewGetCategoriesBiz(store getCategoriesStore) *getCategoriesBiz {
	return &getCategoriesBiz{store: store}
}

func (biz *getCategoriesBiz) GetCategories(ctx context.Context) ([]categoriesmodel.Categories, error) {
	result, err := biz.store.Find(ctx, map[string]interface{}{"status": 1})

	if err != nil {
		return nil, common.ErrEntityNotFound(categoriesmodel.Categories{}.TableName(), err)
	}

	return result, nil
}
