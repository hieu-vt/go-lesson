package orderstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/order/ordermodel"
)

func (s *sqlStore) Create(ctx context.Context, order *ordermodel.Order) error {
	if err := s.db.Create(&order).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
