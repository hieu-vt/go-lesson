package foodbiz

import (
	"context"
	"lesson-5-goland/modules/food/foodmodel"
)

type CreateFoodStore interface {
	Create(ctx context.Context, food *foodmodel.Food) error
}

type createFoodBiz struct {
	store CreateFoodStore
}

func NewBizCreateFood(store CreateFoodStore) *createFoodBiz {
	return &createFoodBiz{store: store}
}

func (biz *createFoodBiz) CreateFood(ctx context.Context, data foodmodel.Food) error {
	data.Status = 1

	return biz.store.Create(ctx, &data)
}
