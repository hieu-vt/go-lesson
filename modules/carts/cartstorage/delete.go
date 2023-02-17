package cartstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/carts/cartmodel"
)

func (s *sqlStore) Delete(ctx context.Context, userId int, foodId int) error {
	db := s.db

	db = db.Table(cartmodel.CartTableName)

	if foodId > 0 {
		db = db.Where("food_id = (?)", foodId)
	}

	if err := db.Where("user_id = (?)", userId).Delete(cartmodel.Cart{}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
