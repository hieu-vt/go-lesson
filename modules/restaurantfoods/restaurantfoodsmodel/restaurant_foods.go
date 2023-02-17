package restaurantfoodsmodel

import "lesson-5-goland/common"

const RestaurantFoodsTableName = "restaurant_foods"

type RestaurantFoods struct {
	common.SqlModel `json:",inline"`
	RestaurantId    int `json:"restaurantId" gorm:"column:restaurant_id"`
	FoodId          int `json:"foodId" gorm:"column:food_id"`
}

func (RestaurantFoods) TableName() string {
	return RestaurantFoodsTableName
}

type CreateRestaurantFoods struct {
	RestaurantId int `json:"restaurantId" gorm:"column:restaurant_id"`
	FoodId       int `json:"foodId" gorm:"column:food_id"`
}
