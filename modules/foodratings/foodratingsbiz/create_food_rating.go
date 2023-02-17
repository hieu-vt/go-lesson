package foodratingsbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/foodratings/foodratingsmodel"
)

type createFoodRatingStore interface {
	Create(ctx context.Context, data *foodratingsmodel.FoodRatings) error
}

type createFoodRatingBiz struct {
	store createFoodRatingStore
}

func NewCreateFoodRatingBiz(store createFoodRatingStore) *createFoodRatingBiz {
	return &createFoodRatingBiz{store: store}
}

func (biz *createFoodRatingBiz) CreateFoodRating(ctx context.Context, data *foodratingsmodel.FoodRatings) error {
	//foodRatingData, err := biz.store.Find(ctx, map[string]interface{}{"food_id": data.FoodId})
	//
	//if err != nil {
	//	return common.ErrEntityExisted(foodmodel.Food{}.TableName(), err)
	//}
	//
	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(foodratingsmodel.FoodRatings{}.TableName(), err)
	}

	return nil
}
