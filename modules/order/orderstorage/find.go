package orderstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/order/ordermodel"
)

func (s *sqlStore) Find(ctx context.Context, userId int) (*[]ordermodel.GetOrderType, error) {
	_, span := component.Tracer.Start(ctx, "order.DB.GetOrder")
	defer span.End()
	db := s.db

	var orders []ordermodel.GetOrderType

	if err := db.Table(ordermodel.TableOrderName).
		Joins("JOIN order_details ON orders.id = order_details.order_id").
		Joins("JOIN restaurants on JSON_EXTRACT(order_details.food_origin, '$.restaurantId') = restaurants.id").
		Joins("JOIN order_trackings ON orders.id = order_trackings.order_id").
		Select("orders.*, order_details.*, order_trackings.*, restaurants.*").
		Where("orders.user_id = ?", userId).
		Find(&orders).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &orders, nil
}
