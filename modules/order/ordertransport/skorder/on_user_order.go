package skorder

import (
	"context"
	socketio "github.com/googollee/go-socket.io"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/order/ordermodel"
	pubsub2 "lesson-5-goland/pubsub"
	"log"
)

type DataOrder struct {
	TotalPrice float64 `json:"totalPrice"`
	//ShipperId  int     `json:"shipperId"`
}

func OnUserOrder(appCtx component.AppContext, requester common.Requester, shipperId int) func(s socketio.Conn, data DataOrder) {
	return func(s socketio.Conn, data DataOrder) {
		pubsub := appCtx.GetPubsub()

		pubsub.Publish(context.Background(), common.TopicHandleOrderWhenUserOrderFood, pubsub2.NewMessage(ordermodel.Order{
			TotalPrice: data.TotalPrice,
			ShipperId:  shipperId,
			UserId:     requester.GetUserId(),
		}))
	}
}

func OnOrderTracking(appCtx component.AppContext, requester common.Requester) func(s socketio.Conn, data interface{}) {
	return func(s socketio.Conn, data interface{}) {
		log.Println("Data receiver from user: ", data)
	}
}
