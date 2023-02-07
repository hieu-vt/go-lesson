package subscriber

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/order/ordermodel"
	"lesson-5-goland/modules/order/ordertransport/skorder"
	"lesson-5-goland/pubsub"
)

type TrackingType string

func HandleOrderAfterClientOrder(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Handle find shipper and push socket to user",
		Hld: func(c context.Context, message *pubsub.Message) error {
			orderData := message.Data().(ordermodel.CreateOrder)

			orderData.Mask(false)

			pub := appCtx.GetPubsub()
			pub.Publish(c, common.TopicEmitEvenWhenUserCreateOrderSuccess, pubsub.NewMessage(skorder.TopicEmitEvenOrderMessageData{
				OrderId: orderData.FakeId.String(),
				CreateOrder: ordermodel.CreateOrder{
					ShipperId: orderData.ShipperId,
					UserId:    orderData.UserId,
				},
				Type: common.WaitingForShipper,
			}))

			return nil
		},
	}
}
