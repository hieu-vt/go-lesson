package skuser

import (
	socketio "github.com/googollee/go-socket.io"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"log"
	"time"
)

type LocationData struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func OnUserUpdateLocation(appCtx component.AppContext, requester common.Requester) func(s socketio.Conn, location LocationData) {
	return func(s socketio.Conn, location LocationData) {
		time.Sleep(2 * time.Second)
		reddit := appCtx.GetReddit()
		log.Println("User before location: user id is", requester.GetUserId(), "at location", reddit.Get(requester.GetUserId()))
		reddit.Save(requester.GetUserId(), location)
		log.Println("User update location: user id is", requester.GetUserId(), "at location", location)
	}
}
