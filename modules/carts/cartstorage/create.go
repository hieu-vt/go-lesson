package cartstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/carts/cartmodel"
)

func (s *sqlStore) Create(ctx context.Context, cart *cartmodel.Cart) error {
	if err := s.db.Create(&cart).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
