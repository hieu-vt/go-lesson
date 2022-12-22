package subscriber

import (
	"context"
	"lesson-5-goland/component"
)

func Setup(appCtx component.AppContext) {
	IncreaseLikeCountAfterUserLikeRestaurant(appCtx, context.Background())
}
