package usermodel

import (
	"errors"
	"lesson-5-goland/common"
)

type UserCreate struct {
	common.SqlModel `json:",inline"`
	Email           string        `json:"email" gorm:"email"`
	Password        string        `json:"password" gorm:"password"`
	Salt            string        `json:"-" gorm:"salt"`
	LastName        string        `json:"last_name" gorm:"last_name"`
	FirstName       string        `json:"first_name" gorm:"first_name"`
	Phone           string        `json:"phone" gorm:"phone"`
	Role            RoleType      `json:"role" gorm:"role"`
	avatar          *common.Image `json:"avatar" gorm:"avatar"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

var (
	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)
