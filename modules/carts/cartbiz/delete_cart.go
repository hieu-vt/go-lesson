package cartbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/carts/cartmodel"
)

type deleteCartStore interface {
	Delete(ctx context.Context, userId int, foodId int) error
}

type deleteCartBiz struct {
	store deleteCartStore
}

func NewDeleteCartBiz(store deleteCartStore) *deleteCartBiz {
	return &deleteCartBiz{store: store}
}

func (biz *deleteCartBiz) DeleteCarts(ctx context.Context, userId int, foodId int) error {
	if err := biz.store.Delete(ctx, userId, foodId); err != nil {
		return common.ErrCannotDeleteEntity(cartmodel.CartTableName, err)
	}

	return nil
}
