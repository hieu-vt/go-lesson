package orderdetailstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/orderdetails/orderdetailmodel"
)

func (s *sqlStore) Create(ctx context.Context, orderDetail *orderdetailmodel.OrderDetail) error {
	if err := s.db.Create(orderDetail).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
