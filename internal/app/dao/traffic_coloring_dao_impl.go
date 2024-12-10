package dao

import (
	"errors"

	"github.com/jerryTJ/controller/internal/app/entity"
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
	if updateResult := si.DB.Updates(serverInfo); updateResult.Error != nil {
		return updateResult.Error
	}
	return nil
}
func (si *ServerDaoImpl) Delete(serverInfo *model.ServerInfo) error {
	if deleteResult := si.DB.Delete(serverInfo); deleteResult.Error != nil {
		return deleteResult.Error
	}
	return nil
}
func (si *ServerDaoImpl) GetPaginatedServerInfos(server *model.ServerInfo, page, pageSize int) (serverinfos []model.ServerInfo, err error) {
	var list []model.ServerInfo
	if listResult := si.DB.Limit(pageSize).Offset((page-1)*pageSize).Where("name like ?", server.Name+"%").Find(&list); listResult != nil {
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

type ChainDaoImpl struct {
	DB *gorm.DB
}

func (cd *ChainDaoImpl) Add(chain *model.Chain) error {
	if createResult := cd.DB.Create(chain); createResult.Error != nil {
		return createResult.Error
	}
	return nil
}
func (cd *ChainDaoImpl) Update(chain *model.Chain) error {
	if createResult := cd.DB.Updates(chain); createResult.Error != nil {
		return createResult.Error
	}
	return nil
}
func (cd *ChainDaoImpl) QueryById(id uint) (*model.Chain, error) {
	var chain model.Chain
	if listResult := cd.DB.Where("id = ?", id).Find(&chain); listResult != nil {
		return &chain, nil
	}
	return nil, errors.New("no exist")
}
func (cd *ChainDaoImpl) Query(name string) (*model.Chain, error) {
	var chain model.Chain
	if queryResult := cd.DB.Preload("ServerInfos").Where("name like ?", name+"%").Find(&chain); queryResult != nil {
		return &chain, nil
	}
	return nil, errors.New("no match record")
}
func (cd *ChainDaoImpl) QueryServerInfos(chainID uint) []model.ServerInfo {
	var chain model.Chain
	if queryResult := cd.DB.Preload("ServerInfos").Where("id = ?", chainID).Find(&chain); queryResult != nil {
		return chain.ServerInfos
	}
	return nil
}
func (cd *ChainDaoImpl) AssociationServerInfo(chain *model.Chain, serverInfos []model.ServerInfo) error {
	for _, serverinfo := range serverInfos {
		err := cd.DB.Model(&chain).Association("ServerInfos").Append(&serverinfo)
		if err != nil {
			return err
		}
	}
	return nil
}
func (cd *ChainDaoImpl) GetPaginatedChains(chain *model.Chain, page, pageSize int) entity.PaginatedResult {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	var maxPageSize int = 100
	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}
	var totalCount int64
	tmpDB := cd.DB.Where("name like ? and version like ?", chain.Name+"%", chain.Version+"%")
	tmpDB.Model(&model.Chain{}).Count(&totalCount)

	chains := make(
		[]model.Chain,
		pageSize,
		pageSize,
	)
	tmpDB.Limit(pageSize).Offset((page - 1) * pageSize).Find(&chains)
	return _getPaginatedResult(chains, totalCount, page, pageSize)
}
func _getPaginatedResult(data interface{}, totalCount int64, page, pageSize int) entity.PaginatedResult {
	//calcalate totalPage
	mode := totalCount % int64(pageSize)
	totalPage := int(totalCount) / pageSize
	if mode > 0 {
		totalPage = totalPage + 1
	}
	nextPage := page + 1
	if nextPage > totalPage {
		nextPage = totalPage
	}

	return entity.PaginatedResult{
		Data:       data,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
		TotalPage:  totalPage,
		NextPage:   nextPage,
	}
}
