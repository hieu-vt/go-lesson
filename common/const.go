package common

const (
	DbTypeRestaurant = 1
	DbTypeFood       = 2
	DbTypeCategory   = 3
	DbTypeUser       = 4
)

const (
	CurrentUser = "user"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

const (
	EmitUserCreateOrderSuccess = "EmitUserCreateOrderSuccess"
	EmitUserOrderFailure       = "EmitUserOrderFailure"
	EmitAuthenticated          = "EmitAuthenticated"
)

const (
	EvenAuthenticated       = "EvenAuthenticated"
	EvenUserCreateOrder     = "EvenUserCreateOrder"
	EventUserUpdateLocation = "EventUserUpdateLocation"
)
