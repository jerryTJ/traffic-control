package model

import (
	"github.com/jerryTJ/controller/init/db"
	"gorm.io/gorm"
)

type ServerInfo struct {
	gorm.Model
	Name    string `gorm:"column:name"`
	Color   string `gorm:"column:color"`
	Domain  string `gorm:"column:domain"`
	Port    string `gorm:"column:port"`
	Image   string `gorm:"column:image"`
	Tag     string `gorm:"column:tag"`
	ChainID string `gorm:"column:chain_id"`
	Version string `gorm:"column:version"`
	IfDown  bool   `gorm:"column:if_down"`
}

// TableName 方法指定表名
func (si *ServerInfo) TableName() string {
	return "server_info"
}
func (c *ServerInfo) QueryServerInfos() (serverInfos []ServerInfo) {
	var servers []ServerInfo
	db.DB.Find(&servers)
	return servers
}
