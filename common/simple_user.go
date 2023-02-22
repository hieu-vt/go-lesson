package common

type SimpleUser struct {
	SqlModel  `json:",inline"`
	LastName  string `json:"last_name" gorm:"last_name"`
	FirstName string `json:"first_name" gorm:"first_name"`
	avatar    *Image `json:"avatar" gorm:"avatar"`
}

func (*SimpleUser) TableName() string {
	return "users"
}
