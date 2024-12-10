package model

import (
	"gorm.io/gorm"
)

type ServerInfo struct {
	gorm.Model
	Namespace string  `gorm:"namespace"`
	Name      string  `gorm:"column:name"`
	Color     string  `gorm:"column:color"`
	Domain    string  `gorm:"column:domain"`
	Port      string  `gorm:"column:port"`
	Image     string  `gorm:"column:image"`
	Tag       string  `gorm:"column:tag"`
	Version   string  `gorm:"column:version"`
	IfDown    bool    `gorm:"column:if_down"`
	Chains    []Chain `gorm:"many2many:t_chain_servers;foreignKey:ID;joinForeignKey:ServerInfoID;"`
}

// TableName 方法指定表名
func (si *ServerInfo) TableName() string {
	return "t_server_infos"
}

type Chain struct {
	gorm.Model
	Name        string       `gorm:"column:name"`
	Version     string       `gorm:"column:version"`
	If_clean    bool         `gorm:"column_if_clean"` // true clean all that server have related chains on server_info table
	ServerInfos []ServerInfo `gorm:"many2many:t_chain_servers;foreignKey:ID;joinForeignKey:ChainID;"`
}

// TableName 方法指定表名
func (ch *Chain) TableName() string {
	return "t_chains"
}

type ChainServer struct {
	gorm.Model
	ChainID      uint `gorm:"primaryKey"`
	ServerInfoID uint `gorm:"primaryKey"`
	Rank         uint `gorom:"rank"`
}

func (cs *ChainServer) TableName() string {
	return "t_chain_servers"
}
