package cartmodel

type UpdateCart struct {
	Quantity int `json:"quantity" gorm:"column:quantity"`
}

func (UpdateCart) TableName() string {
	return Cart{}.TableName()
}
