package orderdetailbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/order/ordermodel"
	"lesson-5-goland/modules/orderdetails/orderdetailmodel"
)

type OrderDetailStore interface {
	Create(ctx context.Context, orderDetail *orderdetailmodel.OrderDetail) error
}

type OrderStore interface {
	FindByCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*ordermodel.Order, error)
}

type orderDetailBiz struct {
	store      OrderDetailStore
	orderStore OrderStore
}

func NewOrderDetailBiz(store OrderDetailStore, orderStore OrderStore) *orderDetailBiz {
	return &orderDetailBiz{
		store:      store,
		orderStore: orderStore,
	}
}

func (biz *orderDetailBiz) CreateOrderDetail(ctx context.Context, data *orderdetailmodel.OrderDetail) error {
	if err := data.ValidateOrderDetailData(); err != nil {
		return common.ErrNoPermission(err)
	}

	order, err := biz.orderStore.FindByCondition(ctx, map[string]interface{}{"orderId": data.OrderId})

	if err != nil {
		return common.ErrEntityNotFound(ordermodel.TableOrderName, err)
	}

	if order.Status == 0 {
		return common.ErrEntityNotFound(ordermodel.TableOrderName, err)
	}

	data.Status = 1

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(orderdetailmodel.TableNameOrderDetail, err)
	}

	return nil
}
