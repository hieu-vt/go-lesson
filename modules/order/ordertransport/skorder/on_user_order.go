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
}

type RealtimeEngine interface {
	EmitToRoom(room string, key string, data interface{}) error
}

type TopicEmitEvenOrderMessageData struct {
	OrderId   string              `json:"orderId"`
	ShipperId int                 `json:"shipperId"`
	UserId    int                 `json:"userId"`
	Type      common.TrackingType `json:"type"`
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

func OnOrderTracking(appCtx component.AppContext, requester common.Requester, rtEngine RealtimeEngine) func(s socketio.Conn, data TopicEmitEvenOrderMessageData) {
	return func(s socketio.Conn, data TopicEmitEvenOrderMessageData) {
		// Đoạn này phần shipper call socket chỉ handle test realtime
		// Thực chất khi shipper nhận socket order start
		// Shipper accept/reject package --> call request --> create pubsub để emit to room là tạo đã accept hoặc reject
		// Từ đó khi shipper cứ update trạng thái vào database thì nó sẽ emit to room cái trạng thái cho user
		// Khi nàp successfully thì update lại ở database và clear process

		roomKey := common.OrderTracking + data.OrderId
		if data.Type == common.OrderShipperAccept {
			log.Println("Tracking order", data.Type)
			// handle join shipper to room and update tracking type
			s.Join(roomKey)
			rtEngine.EmitToRoom(roomKey, common.OrderTracking, TopicEmitEvenOrderMessageData{
				OrderId:   data.OrderId,
				ShipperId: data.ShipperId,
				UserId:    data.UserId,
				Type:      common.OrderProcess,
			})
		}

		if data.Type == common.OrderShipperReject {
			log.Println("Tracking order", data.Type)
			// handle find another shipper
			rtEngine.EmitToRoom(roomKey, common.OrderTracking, TopicEmitEvenOrderMessageData{
				OrderId:   data.OrderId,
				ShipperId: data.ShipperId,
				UserId:    data.UserId,
				Type:      common.OrderShipperReject,
			})
		}

		if data.Type == common.OrderSuccessfully {
			log.Println("Tracking order", data.Type)
			// handle update database
			// handle clear rooms
			rtEngine.EmitToRoom(roomKey, common.OrderTracking, TopicEmitEvenOrderMessageData{
				OrderId:   data.OrderId,
				ShipperId: data.ShipperId,
				UserId:    data.UserId,
				Type:      common.OrderSuccessfully,
			})
		}
	}
}
