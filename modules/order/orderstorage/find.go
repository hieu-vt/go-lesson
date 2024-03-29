package orderstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/order/ordermodel"
)

func (s *sqlStore) Find(ctx context.Context, userId int, paging common.Paging) ([]ordermodel.GetOrderType, error) {
	_, span := component.Tracer.Start(ctx, "order.DB.GetOrder")
	defer span.End()
	db := s.db

	var orders []ordermodel.GetOrderType

	if paging.FakeCursor != "" {
		if uid, err := common.FromBase58(paging.FakeCursor); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		} else {
			return nil, common.ErrDB(err)
		}
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	if err := db.Limit(paging.Limit).Table(ordermodel.TableOrderName).
		Joins("JOIN order_details ON orders.id = order_details.order_id AND order_details.status = 1").
		Joins("JOIN restaurants on JSON_EXTRACT(order_details.food_origin, '$.restaurantId') = restaurants.id").
		Joins("JOIN order_trackings ON orders.id = order_trackings.order_id AND order_trackings.status = 1").
		Select("orders.*, order_details.*, order_trackings.*, restaurants.*").
		Where("orders.user_id = ?", userId).
		Where("orders.status = 1").
		Find(&orders).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return orders, nil
}
