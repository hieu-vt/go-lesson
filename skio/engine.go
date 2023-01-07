package skio

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/component/tokenprovider/jwt"
	"lesson-5-goland/modules/order/ordertransport/skorder"
	"lesson-5-goland/modules/user/usermodel"
	"lesson-5-goland/modules/user/userstorage"
	"lesson-5-goland/modules/user/usertransport/skuser"
	"lesson-5-goland/reddit"
	"log"
	"math"
	"sync"
)

type RealtimeEngine interface {
	UserSockets(userId int) []AppSocket
	EmitToRoom(room string, key string, data interface{}) error
	EmitToUser(userId int, key string, data interface{}) error
	JoinRoom(userId int, room string) error
	Run(ctx component.AppContext, engine *gin.Engine) error
	GetShipper(reddit reddit.RedditEngine, id int, location interface{}) int
	//Emit(userId int) error
}

type rtEngine struct {
	server  *socketio.Server
	storage map[int][]AppSocket
	locker  *sync.RWMutex
}

func NewEngine() *rtEngine {
	return &rtEngine{
		storage: make(map[int][]AppSocket),
		locker:  new(sync.RWMutex),
	}
}

func (engine *rtEngine) saveAppSocket(userId int, appSck AppSocket) {
	engine.locker.Lock()

	//appSck.Join("order-{ordID}")

	if v, ok := engine.storage[userId]; ok {
		engine.storage[userId] = append(v, appSck)
	} else {
		engine.storage[userId] = []AppSocket{appSck}
	}

	engine.locker.Unlock()
}

func (engine *rtEngine) getAppSocket(userId int) []AppSocket {
	engine.locker.RLock()
	defer engine.locker.RUnlock()

	return engine.storage[userId]
}

func (engine *rtEngine) removeAppSocket(userId int, appSck AppSocket) {
	engine.locker.Lock()
	defer engine.locker.Unlock()

	if v, ok := engine.storage[userId]; ok {
		for i := range v {
			if v[i] == appSck {
				engine.storage[userId] = append(v[:i], v[i+1:]...)
				break
			}
		}
	}
}

func (engine *rtEngine) UserSockets(userId int) []AppSocket {
	var sockets []AppSocket

	if scks, ok := engine.storage[userId]; ok {
		return scks
	}

	return sockets
}

func (engine *rtEngine) EmitToRoom(room string, key string, data interface{}) error {
	engine.server.BroadcastToRoom("/", room, key, data)
	return nil
}

func (engine *rtEngine) OnEvent(userId int, key string, f interface{}) error {
	engine.server.OnEvent("/", key, f)

	return nil
}

func (engine *rtEngine) EmitToUser(userId int, key string, data interface{}) error {
	sockets := engine.getAppSocket(userId)
	for _, s := range sockets {
		s.Emit(key, data)
	}

	return nil
}

func (engine *rtEngine) JoinRoom(userId int, room string) error {
	sockets := engine.getAppSocket(userId)

	for _, s := range sockets {
		s.Join(room)
	}

	return nil
}

//:::    optional: unit = the unit you desire for results                     :::
//:::           where: 'M' is statute miles (default, or omitted)             :::
//:::                  'K' is kilometers                                      :::
//:::                  'N' is nautical miles                                  :::
//:::

func calculatorDistance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	radlat1 := float64(math.Pi * lat1 / 180)
	radlat2 := float64(math.Pi * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}

func (engine *rtEngine) GetShipper(reddit reddit.RedditEngine, id int, location interface{}) int {
	userLocation := location.(skuser.LocationData)
	distanceAround := 5 // KM
	var minDistance float64
	minDistance = 0
	var shipperId int
	isSecondCheck := false

	for {
		for _, user := range engine.storage {
			if isSecondCheck {
				currentUser := user[0]
				if id != currentUser.GetUserId() && currentUser.GetRole() == string(usermodel.SHIPPER) && reddit.Get(currentUser.GetUserId()) != nil {
					shipperLocation := reddit.Get(currentUser.GetUserId()).(skuser.LocationData)
					distance := calculatorDistance(userLocation.Lat, userLocation.Lng, shipperLocation.Lat, shipperLocation.Lng)

					if minDistance == 0 {
						minDistance = distance
						shipperId = currentUser.GetUserId()
					} else if minDistance > distance {
						minDistance = distance
						shipperId = currentUser.GetUserId()
					}
				}
			} else {
				currentUser := user[0]
				if id != currentUser.GetUserId() && currentUser.GetRole() == string(usermodel.SHIPPER) && reddit.Get(currentUser.GetUserId()) != nil {
					shipperLocation := reddit.Get(currentUser.GetUserId()).(skuser.LocationData)
					distance := calculatorDistance(userLocation.Lat, userLocation.Lng, shipperLocation.Lat, shipperLocation.Lng)

					if distance < float64(distanceAround) {
						shipperId = currentUser.GetUserId()
						return shipperId
					}
				}
			}
		}
		if isSecondCheck {
			return shipperId
		}

		isSecondCheck = true
	}

	return shipperId
}

func authenticated(appCtx component.AppContext, engine *rtEngine) func(s socketio.Conn, token string) {
	return func(s socketio.Conn, token string) {
		server := engine.server
		db := appCtx.GetMainDBConnection()
		store := userstorage.NewSqlStore(db)

		tokenProvider := jwt.NewTokenJwt(appCtx.SecretKey())

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			s.Emit("authentication_failed", err.Error())
			s.Close()
			return
		}

		user, err := store.FindUser(context.Background(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			s.Emit("authentication_failed", err.Error())
			s.Close()
			return
		}

		if user.Status == 0 {
			s.Emit("authentication_failed", errors.New("you has been banned/deleted"))
			s.Close()
			return
		}

		user.Mask(false)

		appSck := NewAppSocket(s, user)
		engine.saveAppSocket(user.Id, appSck)

		s.Emit(common.EmitAuthenticated, user)

		//appSck.Join(user.GetRole()) // the same
		//if user.GetRole() == "admin" {
		//	appSck.Join("admin")
		//}

		log.Println(user.Id)

		server.OnEvent("/", common.EventUserUpdateLocation+user.FakeId.String(), skuser.OnUserUpdateLocation(appCtx, user))
		server.OnEvent("/", common.EvenUserCreateOrder+user.FakeId.String(), skorder.OnUserOrder(appCtx, user, engine))
		server.OnEvent("/", common.OrderTracking, skorder.OnOrderTracking(appCtx, user, engine))
	}
}

func (engine *rtEngine) Run(appCtx component.AppContext, r *gin.Engine) error {
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{websocket.Default},
	})

	engine.server = server

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID(), " IP:", s.RemoteAddr(), s.ID())
		return nil
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	// Setup

	server.OnEvent("/", common.EvenAuthenticated, authenticated(appCtx, engine))

	go server.Serve()

	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))

	return nil
}
