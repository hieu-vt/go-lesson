package common

import "time"

type SqlModel struct {
	Id        int       `json:"id,omitempty" gorm:"column:id;"`
	Status    int       `json:"status" gorm:"column:status;"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdateAt  time.Time `json:"update_at" gorm:"column:update_at;"`
}
