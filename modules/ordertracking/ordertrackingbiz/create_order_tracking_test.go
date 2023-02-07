package ordertrackingbiz

import (
	"errors"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/ordertracking/ordertrackingmodel"
	"testing"
)

type mockOrderTrackingStore struct{}

type testingItem struct {
	Input    ordertrackingmodel.OrderTracking
	Expected error
	Actual   error
}

func TestCreateOrderTracking(t *testing.T) {
	store := &mockOrderTrackingStore{}
	
	dataTable := []testingItem{
		{
			Input: ordertrackingmodel.OrderTracking{
				OrderId: 0,
				State:   common.WaitingForShipper,
			},
			Expected: errors.New(ordertrackingmodel.OrderIdIsNotEmpty),
			Actual:   nil,
		},
		{
			Input: ordertrackingmodel.OrderTracking{
				OrderId: 1,
				State:   "",
			},
			Expected: errors.New(ordertrackingmodel.StateIsNotEmpty),
			Actual:   nil,
		},
		{
			Input: ordertrackingmodel.OrderTracking{
				OrderId: 2,
				State:   common.WaitingForShipper,
			},
			Expected: nil,
			Actual:   nil,
		},
	}

	for _, item := range dataTable {

	}
}
