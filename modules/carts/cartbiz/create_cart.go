package cartbiz

import (
	"context"
	"errors"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/carts/cartmodel"
)

type CreateCartStore interface {
	Create(ctx context.Context, cart *cartmodel.Cart) error
}

type bizCreateCart struct {
	store CreateCartStore
}

func NewCreateCartBiz(store CreateCartStore) *bizCreateCart {
	return &bizCreateCart{store: store}
}

func (biz *bizCreateCart) CreateCart(ctx context.Context, cart *cartmodel.Cart) error {
	if cart.Quantity <= 0 {
		return common.ErrNoPermission(errors.New(cartmodel.QuantityCannotEmpty))
	}

	if err := biz.store.Create(ctx, cart); err != nil {
		return common.ErrCannotCreateEntity(cartmodel.CartTableName, err)
	}

	return nil
}
