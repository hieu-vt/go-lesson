package orderstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/order/ordermodel"
)

func (s *sqlStore) Find(ctx context.Context, userId int) (*[]ordermodel.GetOrderType, error) {
	db := s.db

	var orders []ordermodel.GetOrderType

	if err := db.
		Joins("JOIN order_trackings ON orders.id = order_trackings.orders_id").
		Joins("JOIN order_details ON orders.id = order_details..orders_id").
		Where("orders.user_id = ?", userId).
		Find(&orders).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &orders, nil
}
