package cartbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/carts/cartmodel"
)

type updateCartStore interface {
	Update(ctx context.Context, foodId int, userId int, updateData *cartmodel.UpdateCart) error
}

type updateCartBiz struct {
	store updateCartStore
}

func NewBizUpdateCart(store updateCartStore) *updateCartBiz {
	return &updateCartBiz{store: store}
}

func (biz *updateCartBiz) UpdateCart(ctx context.Context, userId int, foodId int, updateData *cartmodel.UpdateCart) error {
	if err := biz.store.Update(ctx, foodId, userId, updateData); err != nil {
		return common.ErrCannotUpdateEntity(cartmodel.CartTableName, err)
	}

	return nil
}
