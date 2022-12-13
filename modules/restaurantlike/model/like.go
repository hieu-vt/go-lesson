package restaurantlikemodel

import "time"

type RestaurantLike struct {
	RestaurantId int       `json:"restaurant_id" gorm:"column:restaurant_id"`
	UserId       int       `json:"user_id" gorm:"column:user_id;"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
}

func (*RestaurantLike) TableName() string {
	return "restaurant_likes"
}
