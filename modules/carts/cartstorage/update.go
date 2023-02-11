package cartstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/carts/cartmodel"
)

func (s *sqlStore) Update(ctx context.Context, foodId int, userId int, updateData *cartmodel.UpdateCart) error {
	if err := s.db.
		Where("user_id = ?", userId).
		Where("food_id = ?", foodId).
		Updates(&updateData).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
