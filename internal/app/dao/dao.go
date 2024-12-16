package dao

import (
	"github.com/jerryTJ/controller/internal/app/entity"
	"github.com/jerryTJ/controller/internal/app/model"
)

type ServerDao interface {
	Add(server *model.ServerInfo) error
	Update(server *model.ServerInfo) error
	Delete(server *model.ServerInfo) error
	GetPaginatedServerInfos(server *model.ServerInfo, page, pageSize int) (list []model.ServerInfo, err error)
	ListById(id uint) (server *model.ServerInfo, err error)
	QueryByVersion(name string, version string) (server *model.ServerInfo, err error)
	BindChain(chainId, serverId uint) error
	UnBindChain(serverId, defaultChainId uint) error
}

type ChainDao interface {
	Add(chain *model.Chain) error
	Update(chain *model.Chain) error
	QueryById(id uint) (chain *model.Chain, err error)
	Query(name string) (chain *model.Chain, err error)
	QueryServerInfos(chainID uint) []model.ServerInfo
	AssociationServerInfo(chain *model.Chain, serverInfos []model.ServerInfo) error
	GetPaginatedChains(chain *model.Chain, page, pageSize int) entity.PaginatedResult
}
