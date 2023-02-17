package subscriber

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/order/ordertransport/skorder"
	"lesson-5-goland/pubsub"
	"lesson-5-goland/skio"
	"log"
)

type DataEmitCreateSuccessOrder struct {
	OrderId string `json:"orderId"`
}

func EmitRealtimeAfterOrderEnd(appCtx component.AppContext, rtEngine skio.RealtimeEngine) consumerJob {
	return consumerJob{
		Title: "Emit realtime after order end !",
		Hld: func(c context.Context, message *pubsub.Message) error {
			data := message.Data().(skorder.TopicEmitEvenOrderMessageData)

			log.Println(data.ShipperId, data.UserId)

			roomKey := common.OrderTracking + data.OrderId
			rtEngine.JoinRoom(data.UserId, roomKey)
			rtEngine.EmitToUser(data.ShipperId, common.OrderTracking, data)
			return nil
		},
	}
}
