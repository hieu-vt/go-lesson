package cartbiz

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/carts/cartmodel"
)

type listCartStore interface {
	List(ctx context.Context, userId int, paging common.Paging, moreKeys ...string) ([]cartmodel.Cart, error)
}

type listCartBiz struct {
	store listCartStore
}

func NewListCartBiz(store listCartStore) *listCartBiz {
	return &listCartBiz{store: store}
}

func (biz *listCartBiz) ListCart(ctx context.Context, userId int, paging common.Paging) ([]cartmodel.Cart, error) {
	result, err := biz.store.List(ctx, userId, paging)

	if err != nil {
		return nil, common.ErrCannotGetEntity(cartmodel.CartTableName, err)
	}

	return result, nil
}
