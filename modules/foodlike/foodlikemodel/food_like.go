package foodlikemodel

import (
	"fmt"
	"lesson-5-goland/common"
	"time"
)

type FoodLikes struct {
	UserId    int       `json:"userId" gorm:"column:user_id"`
	FoodId    int       `json:"foodId" gorm:"column:food_id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (FoodLikes) TableName() string {
	return "food_likes"
}

func ErrCannotLikeFood(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot like food"),
		fmt.Sprintf("ErrCannotLikeFood"),
	)
}

func ErrLikeFoodExist(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Like food have been exist"),
		fmt.Sprintf("ErrLikeFoodtExist"),
	)
}

func ErrCannotUnlikeFood(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot unlike food"),
		fmt.Sprintf("ErrCannotUnlikeFood"),
	)
}

func (rLike *FoodLikes) GetFoodId() int {
	return rLike.FoodId
}
