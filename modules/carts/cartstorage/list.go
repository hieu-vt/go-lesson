package cartstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/carts/cartmodel"
)

func (s *sqlStore) List(ctx context.Context, userId int) ([]cartmodel.Cart, error) {
	var result []cartmodel.Cart

	if err := s.db.Where("user_id ?= ", userId).Where("status in (1)").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	
	return result, nil
}
