package app

import (
	"github.com/jerryTJ/controller/init/db"
	"gorm.io/gorm"
)

type ServerInfo struct {
	gorm.Model
	Name    string
	Color   string
	Domain  string
	Port    string
	Image   string
	Tag     string
	ChainID string
	Version string
	IfDown  bool
}

type Chain struct {
	gorm.Model
	ID         int64
	ServerName string
	Color      string
	NextChain  *Chain
	IfLast     bool
}

// TableName 方法指定表名
func (ServerInfo) TableName() string {
	return "server_info"
}
func (c *ServerInfo) QueryServerInfos() (serverInfos []ServerInfo) {
	var servers []ServerInfo
	db.DB.Find(&servers)
	return servers
}
