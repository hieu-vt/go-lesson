package restaurantbiz

import (
	"context"
	"errors"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/restaurant/restaurantmodel"
	"testing"
)

type mockDeleteStore struct {
}

func (s *mockDeleteStore) FindByCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error) {
	var result *restaurantmodel.Restaurant
	var error error
	error = nil
	for i := range condition {
		if condition[i] == 1 {
			result = &restaurantmodel.Restaurant{
				SqlModel: common.SqlModel{
					Status: 1,
				},
				Name: "Hieu",
				Addr: "Somewhere",
			}
			break
		}

		if condition[i] == 2 {
			result = &restaurantmodel.Restaurant{
				SqlModel: common.SqlModel{
					Status: 0,
				},
				Name: "Hieu",
				Addr: "Somewhere",
			}
			break
		}

		if condition[i] == -2 {
			result = nil
			error = errors.New(restaurantmodel.RestaurantNotFound)
			break
		}

		if condition[i] == -1 {
			result = &restaurantmodel.Restaurant{
				SqlModel: common.SqlModel{
					Status: 1,
				},
				Name: "HieuDZ",
				Addr: "Somewhere",
			}

			break
		}

		result = nil
	}

	return result, error
}

func (s *mockDeleteStore) DeleteRestaurantWithCondition(ctx context.Context, id int) error {
	if id == -1 {
		return errors.New(restaurantmodel.RestaurantNotFound)
	}
	return nil
}

type testingDeleteItem struct {
	Input    int
	Expected error
	Actual   error
}

func TestDeleteRestaurantBiz_DeleteRestaurant(t *testing.T) {
	store := &mockDeleteStore{}
	biz := NewDeleteRestaurantBiz(store)

	dataTesting := []testingDeleteItem{
		{
			Input:    1,
			Expected: nil,
			Actual:   nil,
		},
		{
			Input:    0,
			Expected: common.ErrEntityNotFound(restaurantmodel.EntityName, errors.New(restaurantmodel.RestaurantNotFound)),
			Actual:   nil,
		},
		{
			Input:    2,
			Expected: common.ErrEntityNotFound(restaurantmodel.EntityName, errors.New(restaurantmodel.RestaurantNotFound)),
			Actual:   nil,
		},
		{
			Input:    -1,
			Expected: common.ErrCannotDeleteEntity(restaurantmodel.EntityName, errors.New(restaurantmodel.RestaurantNotFound)),
			Actual:   nil,
		},
		{
			Input:    -2,
			Expected: common.ErrEntityNotFound(restaurantmodel.EntityName, errors.New(restaurantmodel.RestaurantNotFound)),
			Actual:   nil,
		},
	}

	for _, item := range dataTesting {
		actual := biz.DeleteRestaurant(context.Background(), item.Input)

		if actual == nil {
			if item.Expected != nil {
				t.Errorf("expect error is %s but actual is %v", item.Expected.Error(), actual.Error())
			}

			continue
		}

		if item.Expected == nil {
			if actual != nil {
				t.Errorf("expect error is %s but actual is %v", item.Expected, actual.Error())
			}

			continue
		}

		if actual.Error() != item.Expected.Error() {
			t.Errorf("expect error is %s but actual is %v", item.Expected.Error(), actual.Error())
		}
	}
}
