package orderdetailbiz

import (
	"context"
	"errors"
	"lesson-5-goland/modules/orderdetails/orderdetailmodel"
	"testing"
)

type mockCreateStore struct {
}

func (s *mockCreateStore) Create(ctx context.Context, orderDetail *orderdetailmodel.CreateOrderDetail) error {
	if orderDetail.OrderId == 2 {
		return errors.New("something wrong with db")
	}
	return nil
}

type testingItem struct {
	Input    orderdetailmodel.CreateOrderDetail
	Expected error
	Actual   error
}

func TestOrderDetailBiz_CreateOrderDetail(t *testing.T) {
	store := &mockCreateStore{}
	biz := NewOrderDetailBiz(store)

	dataTable := []testingItem{
		{
			Input: orderdetailmodel.CreateOrderDetail{
				OrderId:    0,
				FoodOrigin: "12",
				Price:      12,
			},
			Expected: errors.New(orderdetailmodel.OrderIdINotBeEmpty),
			Actual:   nil,
		},
		{
			Input: orderdetailmodel.CreateOrderDetail{
				OrderId:    1,
				FoodOrigin: "",
				Price:      12,
			},
			Expected: errors.New(orderdetailmodel.FoodOriginIsNotEmpty),
			Actual:   nil,
		},
		{
			Input: orderdetailmodel.CreateOrderDetail{
				OrderId:    1,
				FoodOrigin: "222",
				Price:      0,
			},
			Expected: errors.New(orderdetailmodel.PriceMustMoreThan0),
			Actual:   nil,
		},
		{
			Input: orderdetailmodel.CreateOrderDetail{
				OrderId:    1,
				FoodOrigin: "222",
				Price:      12,
			},
			Expected: nil,
			Actual:   nil,
		},
		{
			Input: orderdetailmodel.CreateOrderDetail{
				OrderId:    2,
				FoodOrigin: "222",
				Price:      12,
			},
			Expected: errors.New("something wrong with db"),
			Actual:   nil,
		},
	}

	for _, item := range dataTable {
		actual := biz.CreateOrderDetail(context.Background(), &item.Input)

		if actual == nil {
			if item.Expected != nil {
				t.Errorf("expect error is %s but actual is %v", item.Expected.Error(), actual.Error())
			}

			continue
		}

		if actual.Error() != item.Expected.Error() {
			t.Errorf("expect error is %s but actual is %v", item.Expected.Error(), actual.Error())
		}
	}
}
