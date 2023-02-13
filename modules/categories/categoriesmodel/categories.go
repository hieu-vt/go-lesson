package categoriesmodel

import (
	"errors"
	"lesson-5-goland/common"
	"strings"
)

const NameIsNotEmpty = "name is not empty"

type Categories struct {
	common.SqlModel `json:",inline"`
	Name            string        `json:"name" gorm:"column:name"`
	Description     string        `json:"description" gorm:"column:description"`
	Icon            *common.Image `json:"icon" gorm:"column:icon"`
}

func (Categories) TableName() string {
	return "categories"
}

func (c *Categories) Validate() error {
	c.Name = strings.TrimSpace(c.Name)
	if c.Name == "" {
		return errors.New(NameIsNotEmpty)
	}

	return nil
}

func (c *Categories) Mask() {
	c.GenUID(common.DbTypeCategory)
}

type CreateCategories struct {
	Name        string        `json:"name" gorm:"column:name"`
	Description string        `json:"description" gorm:"column:description"`
	Icon        *common.Image `json:"icon" gorm:"column:icon"`
}
