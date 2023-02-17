package foodratingsbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/foodratings/foodratingsmodel"
)

type deleteFoodRatingStore interface {
	Delete(ctx context.Context, ratingId int) error
}

type deleteFoodRatingBiz struct {
	store deleteFoodRatingStore
}

func NewDeleteFoodRatingBiz(store deleteFoodRatingStore) *deleteFoodRatingBiz {
	return &deleteFoodRatingBiz{store: store}
}

func (biz *deleteFoodRatingBiz) DeleteFoodRating(ctx context.Context, rId int) error {
	if err := biz.store.Delete(ctx, rId); err != nil {
		return common.ErrCannotDeleteEntity(foodratingsmodel.FoodRatings{}.TableName(), err)
	}

	return nil
}
