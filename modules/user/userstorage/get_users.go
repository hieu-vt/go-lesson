package userstorage

import (
	"context"
	"gorm.io/gorm"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
)

func (s *sqlStore) GetUsers(ctx context.Context, ids []int) ([]common.SimpleUser, error) {
	_, span := component.Tracer.Start(ctx, "user.storage.GetUsersByIds")
	defer span.End()
	db := s.db
	var data []common.SimpleUser

	if err := db.Where("id in (?)", ids).Find(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return data, nil
}
