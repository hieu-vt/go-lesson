package foodmodel

import "lesson-5-goland/common"

type Food struct {
	common.SqlModel `json:",inline"`
	RestaurantId    int            `json:"-" gorm:"column:restaurant_id"`
	CategoryId      int            `json:"-" gorm:"column:category_id"`
	Name            string         `json:"name" gorm:"column:name"`
	Description     string         `json:"description" gorm:"column:description"`
	Price           float64        `json:"price" gorm:"column:price"`
	Images          *common.Images `json:",inline" gorm:"column:images"`
}

func (*Food) TableName() string {
	return "foods"
}

type CreateFood struct {
	RestaurantId string         `json:"restaurantId" gorm:"-"`
	CategoryId   string         `json:"categoryId" gorm:"-"`
	Name         string         `json:"name" gorm:"column:name"`
	Description  string         `json:"description" gorm:"column:description"`
	Price        float64        `json:"price" gorm:"column:price"`
	Images       *common.Images `json:"images" gorm:"column:images"`
}

func (*CreateFood) TableName() string {
	return "foods"
}
