package ordermodel

import (
	"lesson-5-goland/common"
)

type GetOrderType struct {
	Order      `json:",inline"`
	State      common.TrackingType `json:"state" gorm:"column:state"`
	FoodOrigin string              `json:"foodOrigin" gorm:"column:food_origin"`
}

func (gOrderType *GetOrderType) Mask() {
	gOrderType.GenUID(common.DbTypeOrder)
}
