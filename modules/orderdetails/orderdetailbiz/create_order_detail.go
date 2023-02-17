package orderdetailbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/order/ordermodel"
	"lesson-5-goland/modules/orderdetails/orderdetailmodel"
	"lesson-5-goland/modules/ordertracking/ordertrackingmodel"
	"lesson-5-goland/pubsub"
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
	pubsub     pubsub.Pubsub
}

func NewOrderDetailBiz(store OrderDetailStore, orderStore OrderStore, pubsub pubsub.Pubsub) *orderDetailBiz {
	return &orderDetailBiz{
		store:      store,
		orderStore: orderStore,
		pubsub:     pubsub,
	}
}

func (biz *orderDetailBiz) CreateOrderDetail(ctx context.Context, data *orderdetailmodel.OrderDetail) error {
	if err := data.ValidateOrderDetailData(); err != nil {
		return common.ErrNoPermission(err)
	}

	order, err := biz.orderStore.FindByCondition(ctx, map[string]interface{}{"id": data.OrderId})

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

	biz.pubsub.Publish(ctx, common.TopicCreateOrderTrackingAfterCreateOrderDetail, pubsub.NewMessage(ordertrackingmodel.CreateOrderTracking{
		//SqlModel: common.SqlModel{},
		OrderId: data.OrderId,
		State:   common.WaitingForShipper,
	}))

	return nil
}
