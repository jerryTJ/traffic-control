package dao

import (
	"errors"

	"github.com/jerryTJ/controller/internal/app/model"
	"gorm.io/gorm"
)

type ServerDaoImpl struct {
	DB *gorm.DB
}

func (si *ServerDaoImpl) Add(serverInfo *model.ServerInfo) error {
	if createResult := si.DB.Create(serverInfo); createResult.Error != nil {
		return createResult.Error
	}
	return nil
}

func (si *ServerDaoImpl) Update(serverInfo *model.ServerInfo) error {
	if updateResult := si.DB.Create(serverInfo); updateResult.Error != nil {
		return updateResult.Error
	}
	return nil
}
func (si *ServerDaoImpl) Delete(serverInfo *model.ServerInfo) error {
	if deleteResult := si.DB.Create(serverInfo); deleteResult.Error != nil {
		return deleteResult.Error
	}
	return nil
}
func (si *ServerDaoImpl) List(serverInfo *model.ServerInfo) ([]model.ServerInfo, error) {
	var list []model.ServerInfo
	if listResult := si.DB.Where("name like ?", serverInfo.Name).Find(&list); listResult != nil {
		return list, nil
	}
	return nil, errors.New("no match condition")
}
func (si *ServerDaoImpl) ListById(id uint) (*model.ServerInfo, error) {
	var server model.ServerInfo
	if listResult := si.DB.Where("id = ?", id).Find(&server); listResult != nil {
		return &server, nil
	}
	return nil, errors.New("no exist")

}
func (si *ServerDaoImpl) QueryByVersion(name string, version string) (*model.ServerInfo, error) {
	var serverInfo model.ServerInfo
	if queryResult := si.DB.Where("name like ? and version = ?", name+"%", version).Find(&serverInfo); queryResult != nil {
		return &serverInfo, nil
	}
	return nil, errors.New("no match record")
}
