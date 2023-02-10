package cartmodel

import "lesson-5-goland/common"

const (
	CartTableName       = "carts"
	QuantityCannotEmpty = "QuantityCannotEmpty"
)

type Cart struct {
	common.SqlModel `json:",inline"`
	UserId          int `json:"userId" gorm:"column:user_id"`
	FoodId          int `json:"foodId" gorm:"column:food_id"`
	Quantity        int `json:"quantity" gorm:"column:quantity"`
}

func (Cart) TableName() string {
	return CartTableName
}

func (cart *Cart) Mask() {
	cart.GenUID(common.DbTypeCart)
}
