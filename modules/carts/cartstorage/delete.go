package cartstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/carts/cartmodel"
)

func (s *sqlStore) Delete(ctx context.Context, ids []int) error {
	if err := s.db.Where("id in (?)", ids).Delete(cartmodel.Cart{}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
