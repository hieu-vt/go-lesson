package foodratingsbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/foodratings/foodratingsmodel"
)

type listFoodRatingStore interface {
	List(ctx context.Context, userId int, paging *common.Paging) ([]foodratingsmodel.FoodRatings, error)
}

type listFoodRatingBiz struct {
	store listFoodRatingStore
}

func NewListFoodRatingBiz(store listFoodRatingStore) *listFoodRatingBiz {
	return &listFoodRatingBiz{store: store}
}

func (biz *listFoodRatingBiz) ListFoodRating(ctx context.Context, foodId int, paging *common.Paging) ([]foodratingsmodel.FoodRatings, error) {
	result, err := biz.store.List(ctx, foodId, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(foodratingsmodel.FoodRatings{}.TableName(), err)
	}

	return result, nil
}
