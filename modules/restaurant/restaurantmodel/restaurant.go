package restaurantmodel

import (
	"errors"
	"lesson-5-goland/common"
	"strings"
)

const (
	EntityName               = "Restaurant"
	RestaurantNameIsNotBlank = "this restaurant name is not blank"
	RestaurantNotFound       = "restaurant not found"
)

type Restaurant struct {
	common.SqlModel `json:",inline"`
	Name            string             `json:"name" gorm:"column:name;"`
	Addr            string             `json:"addr" gorm:"column:addr;"`
	Logo            *common.Image      `json:"logo" gorm:"column:logo;""`
	OwnerId         int                `json:"-" gorm:"column:owner_id;""`
	Cover           *common.Images     `json:"cover" gorm:"column:cover;""`
	LikeCount       int                `json:"like_count" gorm:"column:like_count;"`
	Owner           *common.SimpleUser `json:"owner" gorm:"-"`
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
	common.SqlModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	Addr            string         `json:"addr" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	OwnerId         int            `json:"-" gorm:"column:owner_id;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) ValidateRestaurantData() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return errors.New(RestaurantNameIsNotBlank)
	}

	return nil
}

type RestaurantDelete struct {
	Status *int `json:"status" gorm:"column:status;"`
}

func (RestaurantDelete) TableName() string {
	return Restaurant{}.TableName()
}

func (data *Restaurant) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}

func (data *RestaurantCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}
