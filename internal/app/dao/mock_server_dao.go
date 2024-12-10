package dao

import (
	"github.com/jerryTJ/controller/internal/app/model"
	"github.com/stretchr/testify/mock"
)

type MockServerDao struct {
	mock.Mock
}

func (m *MockServerDao) Add(server *model.ServerInfo) error {
	args := m.Called(server)
	return args.Error(0)
}

func (m *MockServerDao) Update(server *model.ServerInfo) error {
	args := m.Called(server)
	return args.Error(0)
}

func (m *MockServerDao) Delete(server *model.ServerInfo) error {
	args := m.Called(server)
	return args.Error(0)
}

func (m *MockServerDao) GetPaginatedServerInfos(server *model.ServerInfo, page, pageSize int) (serverinfos []model.ServerInfo, err error) {
	args := m.Called(server)
	return args.Get(0).([]model.ServerInfo), args.Error(1)
}

func (m *MockServerDao) ListById(id uint) (*model.ServerInfo, error) {
	args := m.Called(id)
	return args.Get(0).(*model.ServerInfo), args.Error(1)
}

func (m *MockServerDao) QueryByVersion(name string, version string) (*model.ServerInfo, error) {
	args := m.Called(name, version)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.ServerInfo), args.Error(1)
}
