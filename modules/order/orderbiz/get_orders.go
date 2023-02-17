package orderbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/order/ordermodel"
)

type orderStore interface {
	Find(ctx context.Context, userId int, paging common.Paging) ([]ordermodel.GetOrderType, error)
}

type getOrderBiz struct {
	store orderStore
}

func NewGetOrderBiz(store orderStore) *getOrderBiz {
	return &getOrderBiz{store: store}
}

func (biz *getOrderBiz) GetOrders(
	ctx context.Context,
	userId int,
	paging common.Paging,
) ([]ordermodel.GetOrderType, error) {
	ctx, span := component.Tracer.Start(ctx, "order.biz.GetOrder")
	defer span.End()
	orders, err := biz.store.Find(ctx, userId, paging)

	if err != nil {
		return nil, common.ErrEntityNotFound(ordermodel.TableOrderName, err)
	}

	return orders, nil
}
