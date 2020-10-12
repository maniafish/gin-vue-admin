// 自动生成模板ProjectA
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type ProjectA struct {
	gorm.Model
	Name string `json:"name" form:"name" gorm:"column:name;comment:"`
	Male *bool  `json:"male" form:"male" gorm:"column:male;comment:"`
	Age  int    `json:"age" form:"age" gorm:"column:age;comment:"`
}

func (ProjectA) TableName() string {
	return "ProjectA"
}
