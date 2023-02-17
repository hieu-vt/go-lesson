package restaurantlikemodel

import (
	"fmt"
	"lesson-5-goland/common"
	"time"
)

type RestaurantLike struct {
	RestaurantId int       `json:"restaurant_id" gorm:"column:restaurant_id"`
	UserId       int       `json:"user_id" gorm:"column:user_id;"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
}

func (RestaurantLike) TableName() string {
	return "restaurant_likes"
}

type RestaurantCreateLike struct {
	RestaurantId int `json:"restaurant_id" gorm:"column:restaurant_id"`
	UserId       int `json:"user_id" gorm:"column:user_id;"`
}

func (RestaurantCreateLike) TableName() string {
	return "restaurant_likes"
}

func (rLike *RestaurantCreateLike) GetRestaurantId() int {
	return rLike.RestaurantId
}

func ErrCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot like restaurant"),
		fmt.Sprintf("ErrCannotLikeRestaurant"),
	)
}

func ErrLikeRestaurantExist(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Like restaurant have been exist"),
		fmt.Sprintf("ErrLikeRestaurantExist"),
	)
}

func ErrCannotUnlikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot unlike restaurant"),
		fmt.Sprintf("ErrCannotUnlikeRestaurant"),
	)
}

func ErrLikeRestaurantDidLikeThisRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("User did like restaurant before"),
		fmt.Sprintf("ErrLikeRestaurantDidLikeThisRestaurant"),
	)
}
