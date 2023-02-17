package ordertrackingstorage

import (
	"context"
	"lesson-5-goland/common"
	"lesson-5-goland/modules/ordertracking/ordertrackingmodel"
)

func (s *sqlStore) Update(ctx context.Context, data *ordertrackingmodel.UpdateOrderTracking) error {
	if err := s.db.Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
