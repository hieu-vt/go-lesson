package ordertrackingstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/ordertracking/ordertrackingmodel"
)

func (s *sqlStore) Create(ctx context.Context, data *ordertrackingmodel.OrderTracking) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}