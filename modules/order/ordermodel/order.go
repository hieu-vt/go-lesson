package ordermodel

import "lesson-5-goland/common"

//`id` int(11) NOT NULL AUTO_INCREMENT,
//`user_id` int(11) NOT NULL,
//`total_price` float NOT NULL,
//`shipper_id` int(11) DEFAULT NULL,
//`status` int(11) NOT NULL DEFAULT '1',
//`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
//`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

type Order struct {
	common.SqlModel `json:",inline"`

	UserId     int     `json:"-" gorm:"user_id"`
	ShipperId  int     `json:"-" gorm:"shipper_id"`
	TotalPrice float64 `json:"totalPrice" gorm:"column:total_price"`
}

func (Order) TableName() string {
	return "orders"
}

type CreateOrder struct {
	UserId     int     `json:"userId" gorm:"-"`
	ShipperId  int     `json:"shipperId" gorm:"-"`
	TotalPrice float64 `json:"totalPrice" gorm:"column:total_price"`
}

func (*CreateOrder) TableName() string {
	return Order{}.TableName()
}

func (order *Order) Mask(isAdminOrOwner bool) {
	order.GenUID(common.DbTypeFood)
}

func (order *CreateOrder) GetTotalPrice() float64 {
	return order.TotalPrice
}
