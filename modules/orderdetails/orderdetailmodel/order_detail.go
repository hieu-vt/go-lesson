package orderdetailmodel

import (
	"errors"
	"lesson-5-goland/common"
	"strings"
)

const (
	TableNameOrderDetail = "order_details"
	FoodOriginIsNotEmpty = "FoodOrigin is not empty"
	PriceMustMoreThan0   = "Price must more than 0"
)

type OrderDetail struct {
	common.SqlModel `json:",inline"`
	OrderId         int     `json:"orderId" gorm:"order_id"`
	FoodOrigin      string  `json:"foodOrigin" gorm:"food_origin"`
	Price           float32 `json:"price" gorm:"price"`
	Quantity        int     `json:"quantity" gorm:"quantity"`
	Discount        float32 `json:"discount" gorm:"quantity"`
}

func (OrderDetail) TableName() string {
	return TableNameOrderDetail
}

type CreateOrderDetail struct {
	OrderId    string  `json:"orderId" gorm:"_"`
	FoodOrigin string  `json:"foodOrigin" gorm:"food_origin"`
	Price      float32 `json:"price" gorm:"price"`
	Quantity   int     `json:"quantity" gorm:"quantity"`
	Discount   float32 `json:"discount" gorm:"quantity"`
}

func (CreateOrderDetail) TableName() string {
	return OrderDetail{}.TableName()
}

func (res *CreateOrderDetail) ValidateOrderDetailData() error {
	res.FoodOrigin = strings.TrimSpace(res.FoodOrigin)
	if res.FoodOrigin == "" {
		return errors.New(FoodOriginIsNotEmpty)
	}

	if res.Price <= 0 {
		return errors.New(PriceMustMoreThan0)
	}

	return nil
}
