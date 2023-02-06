package orderdetailbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/orderdetails/orderdetailmodel"
)

type OrderDetailStore interface {
	Create(ctx context.Context, orderDetail *orderdetailmodel.CreateOrderDetail) error
}

type orderDetailBiz struct {
	store OrderDetailStore
}

func NewOrderDetailBiz(store OrderDetailStore) *orderDetailBiz {
	return &orderDetailBiz{
		store: store,
	}
}

func (biz *orderDetailBiz) CreateOrderDetail(ctx context.Context, data *orderdetailmodel.CreateOrderDetail) error {
	if err := data.ValidateOrderDetailData(); err != nil {
		return common.ErrNoPermission(err)
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(orderdetailmodel.TableNameOrderDetail, err)
	}

	return nil
}
