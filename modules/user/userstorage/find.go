package userstorage

import (
	"context"
	"gorm.io/gorm"
	"lesson-5-goland/common"
	"lesson-5-goland/component"
	"lesson-5-goland/modules/user/usermodel"
)

func (s *sqlStore) FindUser(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*usermodel.User, error) {
	_, span := component.Tracer.Start(ctx, "user.storage.FindUser")
	defer span.End()
	db := s.db
	var data usermodel.User

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
