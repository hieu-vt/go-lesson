package foodratingsmodel

import "lesson-5-goland/common"

type FoodRatings struct {
	common.SqlModel `json:",inline"`
	UserId          int     `json:"userId" gorm:"column:user_id"`
	FoodId          int     `json:"foodId" gorm:"column:food_id"`
	Point           float32 `json:"point" gorm:"column:point"`
	Comment         string  `json:"comment" gorm:"column:comment"`
}

type CreateFoodRatings struct {
	Point   float32 `json:"point"`
	Comment string  `json:"comment"`
}

type UpdateFoodRatings struct {
	UserId  int     `json:"userId" gorm:"column:user_id"`
	FoodId  int     `json:"foodId" gorm:"column:food_id"`
	Point   float32 `json:"point" gorm:"column:point"`
	Comment string  `json:"comment" gorm:"column:comment"`
}

func (FoodRatings) TableName() string {
	return "food_ratings"
}

func (UpdateFoodRatings) TableName() string {
	return "food_ratings"
}

func (fR *FoodRatings) Mask() {
	fR.GenUID(common.DbTypeFoodRating)
}
