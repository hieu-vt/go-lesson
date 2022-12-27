package skio

import (
	"lesson-5-goland/common"
	"net"
	"net/http"
	"net/url"
)

type Conn interface {
	// ID returns session id
	ID() string
	Close() error
	URL() url.URL
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	RemoteHeader() http.Header

	// Context of this connection. You can save one context for one
	// connection, and share it between all handlers. The handlers
	// is called in one goroutine, so no need to lock context if it
	// only be accessed in one connection.
	Context() interface{}
	SetContext(ctx interface{})

	Namespace() string
	Emit(eventName string, v ...interface{})

	Join(room string)
	Leave(room string)
	LeaveAll()
	Rooms() []string
}

type AppSocket interface {
	Conn
	common.Requester
}

type appSocket struct {
	Conn
	common.Requester
}

func NewAppSocket(conn Conn, requester common.Requester) *appSocket {
	return &appSocket{conn, requester}
}
