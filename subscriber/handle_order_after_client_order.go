package subscriber

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/order/ordermodel"
	"lesson-5-goland/modules/order/orderstorage"
	"lesson-5-goland/pubsub"
)

type TopicEmitEvenOrderMessageData struct {
	OrderId   string `json:"orderId"`
	ShipperId int    `json:"shipperId"`
	UserId    int    `json:"userId"`
}

type HasTotalPrice interface {
	GetTotalPrice() float64
}

func HandleOrderAfterClientOrder(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Handle find shipper and push socket to user",
		Hld: func(c context.Context, message *pubsub.Message) error {
			store := orderstorage.NewSqlStore(appCtx.GetMainDBConnection())

			orderData := message.Data().(ordermodel.Order)

			orderData.Status = 1

			if err := store.Create(c, &orderData); err != nil {
				return err
			}

			orderData.Mask(false)

			pub := appCtx.GetPubsub()
			pub.Publish(c, common.TopicEmitEvenWhenUserCreateOrderSuccess, pubsub.NewMessage(TopicEmitEvenOrderMessageData{
				OrderId:   orderData.FakeId.String(),
				ShipperId: orderData.ShipperId,
				UserId:    orderData.UserId,
			}))

			return nil
		},
	}
}
