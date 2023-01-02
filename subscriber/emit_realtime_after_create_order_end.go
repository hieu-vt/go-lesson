package subscriber

import (
	"context"
	"lesson-5-goland/component"
	"lesson-5-goland/pubsub"
	"lesson-5-goland/skio"
	"log"
	"time"
)

type DataEmitCreateSuccessOrder struct {
	OrderId string `json:"orderId"`
}

func EmitRealtimeAfterOrderEnd(appCtx component.AppContext, rtEngine skio.RealtimeEngine) consumerJob {
	return consumerJob{
		Title: "Emit realtime after order end !",
		Hld: func(c context.Context, message *pubsub.Message) error {
			data := message.Data().(TopicEmitEvenOrderMessageData)

			log.Println(data.ShipperId, data.UserId)

			//rtEngine.EmitToUser(data.UserId, "OrderTracking", data)
			roomKey := "orders/" + data.OrderId
			rtEngine.JoinRoom(data.UserId, roomKey)
			time.Sleep(time.Second)
			//rtEngine.JoinRoom(data.ShipperId, roomKey)
			time.Sleep(time.Second)
			rtEngine.EmitToRoom(roomKey, "OrderTracking", data)
			return nil
		},
	}
}
