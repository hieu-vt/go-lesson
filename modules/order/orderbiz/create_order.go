package orderbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/order/ordermodel"
)

type OrderStore interface {
	Create(ctx context.Context, order *ordermodel.CreateOrder) error
}

type createOrderBiz struct {
	orderStore OrderStore
}

func NewCreateOrderBiz(orderStore OrderStore) *createOrderBiz {
	return &createOrderBiz{orderStore: orderStore}
}

func (biz *createOrderBiz) CreateOrder(ctx context.Context, order *ordermodel.CreateOrder) error {
	if err := order.ValidateOrderData(); err != nil {
		return common.ErrNoPermission(err)
	}

	if err := biz.orderStore.Create(ctx, order); err != nil {
		return common.ErrCannotCreateEntity(ordermodel.TableOrderName, err)
	}

	return nil
}
