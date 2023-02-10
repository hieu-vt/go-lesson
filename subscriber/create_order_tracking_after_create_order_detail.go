package subscriber

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/ordertracking/ordertrackingmodel"
	"lesson-5-goland/modules/ordertracking/ordertrackingstorage"
	"lesson-5-goland/pubsub"
)

func CreateOrderTrackingAfterCreateOrderDetail(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Create order tracking after create order details",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := ordertrackingstorage.NewSqlStore(appCtx.GetMainDBConnection())
			data := message.Data().(ordertrackingmodel.CreateOrderTracking)

			return store.Create(ctx, &ordertrackingmodel.OrderTracking{
				SqlModel: common.SqlModel{
					Status: 1,
				},
				OrderId: data.OrderId,
				State:   data.State,
			})
		},
	}
}
