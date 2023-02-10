package cartbiz

import (
	"context"
	"errors"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/carts/cartmodel"
	"testing"
)

type mockCreateCartStore struct{}

func (s *mockCreateCartStore) Create(ctx context.Context, cart *cartmodel.Cart) error {
	if cart.FoodId == 3 {
		return errors.New("something wrong with DB")
	}
	return nil
}

type testingItem struct {
	Input    cartmodel.Cart
	Expected error
	Actual   error
}

func TestCreateCart(t *testing.T) {
	store := &mockCreateCartStore{}
	biz := NewCreateCartBiz(store)

	dataTable := []testingItem{
		{
			Input: cartmodel.Cart{
				FoodId:   1,
				Quantity: 0,
			},
			Expected: errors.New(cartmodel.QuantityCannotEmpty),
			Actual:   nil,
		},
		{
			Input: cartmodel.Cart{
				FoodId:   3,
				Quantity: 1,
			},
			Expected: common.ErrCannotCreateEntity(cartmodel.CartTableName, errors.New("something wrong with DB")),
			Actual:   nil,
		},
		{
			Input: cartmodel.Cart{
				FoodId:   1,
				Quantity: 2,
			},
			Expected: nil,
			Actual:   nil,
		},
	}

	for _, item := range dataTable {
		actual := biz.CreateCart(context.Background(), &item.Input)

		if actual == nil {
			if item.Expected != nil {
				t.Errorf("with input %#v expect error is %s but actual is %v", item.Input, item.Expected.Error(), actual.Error())
			}

			continue
		}

		if actual.Error() != item.Expected.Error() {
			t.Errorf("with input %#v expect error is %s but actual is %v", item.Input, item.Expected.Error(), actual.Error())
		}
	}
}
