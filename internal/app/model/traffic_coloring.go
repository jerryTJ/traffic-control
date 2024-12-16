package model

import (
	"time"

	"gorm.io/gorm"
)

type ServerInfo struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Namespace string         `gorm:"namespace"`
	Name      string         `gorm:"column:name"`
	Color     string         `gorm:"column:color"`
	Domain    string         `gorm:"column:domain"`
	Port      string         `gorm:"column:port"`
	Image     string         `gorm:"column:image"`
	Tag       string         `gorm:"column:tag"`
	Version   string         `gorm:"column:version"`
	IfDown    bool           `gorm:"column:if_down"`
	//Chains    []Chain        `gorm:"many2many:t_chain_servers;foreignKey:ID;joinForeignKey:ServerInfoID;"`
	ChainId uint `gorm:"column:chain_id"`
}

// TableName 方法指定表名
func (si *ServerInfo) TableName() string {
	return "t_server_infos"
}

type Chain struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"column:name"`
	Version   string         `gorm:"column:version"`
	IfClean   bool           `gorm:"column_if_clean"` // true clean all that server have related chains on server_info table
	//ServerInfos []ServerInfo   `gorm:"many2many:t_chain_servers;foreignKey:ID;joinForeignKey:ChainID;"`
	ServerInfos []ServerInfo `gorm:"foreignKey:ChainId;"`
}

// TableName 方法指定表名
func (ch *Chain) TableName() string {
	return "t_chains"
}

type ChainServer struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	ChainID      uint           `gorm:"primaryKey"`
	ServerInfoID uint           `gorm:"primaryKey"`
	Rank         uint           `gorom:"rank"`
}

func (cs *ChainServer) TableName() string {
	return "t_chain_servers"
}
