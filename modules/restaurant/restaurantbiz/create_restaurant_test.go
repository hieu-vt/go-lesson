package restaurantbiz

import (
	"context"
	"errors"
	"lesson-5-goland/modules/restaurant/restaurantmodel"
	"testing"
)

type mockCreateStore struct {
}

func (s *mockCreateStore) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if data.Name == "Hieu" {
		return errors.New("that have some wrong of the database")
	}
	return nil
}

type testingItem struct {
	Input    restaurantmodel.RestaurantCreate
	Expected error
	Actual   error
}

func TestCreateRestaurantBiz_CreateRestaurant(t *testing.T) {
	store := &mockCreateStore{}
	biz := NewCreateRestaurantBiz(store)

	dataTable := []testingItem{
		{
			Input: restaurantmodel.RestaurantCreate{
				Name: "",
				Addr: "",
			},
			Expected: errors.New(restaurantmodel.RestaurantNameIsNotBlank),
		},
		{
			Input: restaurantmodel.RestaurantCreate{
				Name: "Hieu",
				Addr: "",
			},
			Expected: errors.New("that have some wrong of the database"),
		},
		{
			Input: restaurantmodel.RestaurantCreate{
				Name: "Hop Tac Xa",
				Addr: "",
			},
			Expected: nil,
		},
	}

	for _, item := range dataTable {
		actual := biz.CreateRestaurant(context.Background(), &item.Input)

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
