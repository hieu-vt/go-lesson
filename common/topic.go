package common

const (
	// restaurant
	TopicUserLikeRestaurant    = "TopicUserLikeRestaurant"
	TopicUserDislikeRestaurant = "TopicUserDislikeRestaurant"

	// socket
	TopicHandleOrderWhenUserOrderFood       = "TopicHandleOrderWhenUserOrderFood"
	TopicEmitEvenWhenUserCreateOrderSuccess = "TopicEmitEvenWhenUserCreateOrderSuccess"

	// order
	TopicCreateOrderTrackingAfterCreateOrderDetail = "TopicCreateOrderTrackingAfterCreateOrderDetail"
)
