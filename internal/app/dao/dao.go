package dao

import "github.com/jerryTJ/controller/internal/app/model"

type ServerDao interface {
	Add(server *model.ServerInfo) error
	Update(server *model.ServerInfo) error
	Delete(server *model.ServerInfo) error
	List(server *model.ServerInfo) (list []model.ServerInfo, err error)
	ListById(id uint) (server *model.ServerInfo, err error)
	QueryByVersion(name string, version string) (server *model.ServerInfo, err error)
}
