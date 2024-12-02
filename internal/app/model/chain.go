package model

import (
	"gorm.io/gorm"
)

type Chain struct {
	gorm.Model
	name       string `gorm:"column:name"`
	ServerName string `gorm:"column:server_name"`
	Color      string `gorm:"column:color"`
	NextId     string `gorm:"column:next_id"`
	IfLast     bool   `gorm:"column:if_last"`
}

// TableName 方法指定表名
func (ch *Chain) TableName() string {
	return "chain"
}
