package usermodel

import (
	"errors"
	"lesson-5-goland/common"
)

type RoleType string

const (
	USER    RoleType = "user"
	ADMIN   RoleType = "admin"
	SHIPPER RoleType = "shipper"
)

const EntityName = "Users"

type User struct {
	common.SqlModel `json:",inline"`
	Email           string        `json:"email" gorm:"email"`
	Password        string        `json:"-" gorm:"password"`
	Salt            string        `json:"-" gorm:"salt"`
	LastName        string        `json:"last_name" gorm:"last_name"`
	FirstName       string        `json:"first_name" gorm:"first_name"`
	Phone           string        `json:"phone" gorm:"phone"`
	Role            RoleType      `json:"role" gorm:"role"`
	avatar          *common.Image `json:"avatar" gorm:"avatar"`
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return string(u.Role)
}

func (User) TableName() string {
	return "users"
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"email"`
	Password string `json:"password" form:"password" gorm:"password"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

//func NewAccount(at, rt *tokenprovider.Token) *Account {
//	return &Account{
//		AccessToken:  at,
//		RefreshToken: rt,
//	}
//}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)
)

func (data *User) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}
