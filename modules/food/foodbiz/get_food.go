package foodbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/food/foodmodel"
)

type getFoodStore interface {
	List(ctx context.Context, paging *common.Paging, moreKeys ...string) ([]foodmodel.Food, error)
}

type getFoodBiz struct {
	store getFoodStore
}

func NewGetFoodBiz(store getFoodStore) *getFoodBiz {
	return &getFoodBiz{store: store}
}

func (biz *getFoodBiz) GetFood(ctx context.Context, paging *common.Paging) ([]foodmodel.Food, error) {
	return biz.store.List(ctx, paging)
}
