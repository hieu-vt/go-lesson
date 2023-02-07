package ordertrackingmodel

import "lesson-5-goland/common"

const (
	TableNameOrderTracking = "order_trackings"
	OrderIdIsNotEmpty      = "orderId is not empty"
	StateIsNotEmpty        = "state is not empty"
)

type OrderTracking struct {
	common.SqlModel `json:",inline"`
	OrderId         int                 `json:"orderId" gorm:"column:order_id"`
	State           common.TrackingType `json:"state" gorm:"column:state"`
}

type CreateOrderTracking struct {
	OrderId int                 `json:"orderId" gorm:"column:order_id"`
	State   common.TrackingType `json:"state" gorm:"column:state"`
}

func (OrderTracking) TableName() string {
	return TableNameOrderTracking
}

func (oTracking *OrderTracking) Mask() {
	oTracking.GenUID(common.DbTypeOrder)
}
