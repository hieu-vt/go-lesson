package restaurantmodel

import (
	"errors"
	"strings"
)

type Restaurant struct {
	Id   int    `json:"id,omitempty" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) ValidateRestaurantData() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return errors.New("this restaurant name is not blank")
	}

	return nil
}
