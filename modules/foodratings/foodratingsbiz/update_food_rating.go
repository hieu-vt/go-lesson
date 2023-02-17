package foodratingsbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/foodratings/foodratingsmodel"
)

type updateFoodRatingStore interface {
	Update(ctx context.Context, ratingId int, data *foodratingsmodel.UpdateFoodRatings) error
}

type updateFoodRatingBiz struct {
	store updateFoodRatingStore
}

func NewUpdateFoodRatingBiz(store updateFoodRatingStore) *updateFoodRatingBiz {
	return &updateFoodRatingBiz{store: store}
}

func (biz *updateFoodRatingBiz) UpdateFoodRating(ctx context.Context, rId int, dataUpdate *foodratingsmodel.UpdateFoodRatings) error {
	if err := biz.store.Update(ctx, rId, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(foodratingsmodel.FoodRatings{}.TableName(), err)
	}

	return nil
}
