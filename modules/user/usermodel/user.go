package usermodel

import "lesson-5-goland/common"

type roleType string

const (
	USER    roleType = "user"
	ADMIN   roleType = "admin"
	SHIPPER roleType = "shipper"
)

const EntityName = "Users"

type User struct {
	common.SqlModel `json:"inline"`
	Email           string        `json:"email" gorm:"email"`
	Password        string        `json:"-" gorm:"password"`
	Salt            string        `json:"-" gorm:"salt"`
	LastName        string        `json:"last_name" gorm:"last_name"`
	FirstName       string        `json:"first_name" gorm:"first_name"`
	Phone           string        `json:"phone" gorm:"phone"`
	Role            roleType      `json:"role" gorm:"role"`
	avatar          *common.Image `json:"avatar" gorm:"avatar"`
}

func (User) TableName() string {
	return "users"
}
