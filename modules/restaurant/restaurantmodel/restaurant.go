package restaurantmodel

import (
	"errors"
	"lesson-5-goland/common"
	"strings"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SqlModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	Addr            string         `json:"addr" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;""`
	Cover           *common.Images `json:"cover" gorm:"column:cover;""`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	common.SqlModel `json:",inline"`
	Name            *string        `json:"name" gorm:"column:name;"`
	Addr            *string        `json:"addr" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	Name  string         `json:"name" gorm:"column:name;"`
	Addr  string         `json:"addr" gorm:"column:addr;"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
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

type RestaurantDelete struct {
	Status *int `json:"status" gorm:"column:status;"`
}

func (RestaurantDelete) TableName() string {
	return Restaurant{}.TableName()
}
