package pubsublocal

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/pubsub"
	"log"
	"sync"
)

type pubsublocal struct {
	messageQueue chan *pubsub.Message
	mapChannel   map[pubsub.Topic][]chan *pubsub.Message
	locker       *sync.RWMutex
}

func NewPubSub() *pubsublocal {
	pb := &pubsublocal{
		messageQueue: make(chan *pubsub.Message, 10000),
		mapChannel:   make(map[pubsub.Topic][]chan *pubsub.Message),
		locker:       new(sync.RWMutex),
	}

	pb.run()

	return pb
}

func (ps *pubsublocal) Publish(ctx context.Context, topic pubsub.Topic, data *pubsub.Message) error {
	data.SetChannel(topic)

	go func() {
		defer common.AppRecover()
		ps.messageQueue <- data
		log.Println("New event published:", data.String(), "with data", data.Data())
	}()
	return nil
}

func (ps *pubsublocal) Subscribe(ctx context.Context, topic pubsub.Topic) (ch <-chan *pubsub.Message, close func()) {
	c := make(chan *pubsub.Message)

	ps.locker.Lock()

	if val, ok := ps.mapChannel[topic]; ok {
		val = append(ps.mapChannel[topic], c)
		ps.mapChannel[topic] = val
	} else {
		ps.mapChannel[topic] = []chan *pubsub.Message{c}
	}

	ps.locker.Unlock()

	return c, func() {
		log.Println("Unsubscribe")

		if chans, ok := ps.mapChannel[topic]; ok {
			for i := range chans {
				if chans[i] == c {
					// remove element at index in chans
					chans = append(chans[:i], chans[i+1:]...)

					ps.locker.Lock()
					ps.mapChannel[topic] = chans
					ps.locker.Unlock()
					break
				}
			}
		}
	}
}

func (ps *pubsublocal) run() error {
	log.Println("Pubsub started")

	go func() {
		for {
			mess := <-ps.messageQueue
			log.Println("Message dequeue:", mess)

			if subs, ok := ps.mapChannel[mess.Channel()]; ok {
				for i := range subs {
					go func(c chan *pubsub.Message) {
						c <- mess
						//f(mess)
					}(subs[i])
				}
			}
			//else {
			//	ps.messageQueue <- mess
			//}
		}
	}()

	return nil
}
