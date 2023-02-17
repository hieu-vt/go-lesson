package foodratingstorage

import (
	"context"
	"lesson-5-goland/common"
)

func (s *sqlStore) Delete(ctx context.Context, ratingId int) error {
	if err := s.db.Where("id = ?", ratingId).Update("status", 0).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
