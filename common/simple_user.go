package common

type SimpleUser struct {
	SqlModel  `json:",inline"`
	LastName  string `json:"last_name" gorm:"last_name"`
	FirstName string `json:"first_name" gorm:"first_name"`
	Avatar    *Image `json:"avatar" gorm:"avatar"`
	Role      string `json:"role" gorm:"role"`
}

func (*SimpleUser) TableName() string {
	return "users"
}

func (s *SimpleUser) Mask() {
	s.GenUID(DbTypeUser)
}
