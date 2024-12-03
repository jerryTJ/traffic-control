package adapter

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jerryTJ/controller/internal/app/dao"
	"github.com/jerryTJ/controller/internal/app/model"
	"github.com/stretchr/testify/assert"
)

type mockDaoFactory struct {
	mockServerDao *dao.MockServerDao
}

func (df *mockDaoFactory) GetServerInfoDao() dao.ServerDao {
	return df.mockServerDao
}

func TestQueryServerInfos(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	mockDao := new(dao.MockServerDao)
	mockDS := &mockDaoFactory{mockServerDao: mockDao}

	tests := []struct {
		name         string
		serverName   string
		version      string
		mockResponse *model.ServerInfo
		mockError    error
		expectedCode int
	}{
		{
			name:       "successful_query",
			serverName: "test-server",
			version:    "1.0.0",
			mockResponse: &model.ServerInfo{
				Name:    "test-server",
				Version: "1.0.0",
				Color:   "blue",
			},
			mockError:    nil,
			expectedCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock expectations
			mockDao.On("QueryByVersion", tt.serverName, tt.version).
				Return(tt.mockResponse, tt.mockError)

			// Setup router
			router := gin.New()
			controller := NewServerInfoController(mockDS)
			router.GET("/:name/:version", controller.QueryServerInfos)

			// Create request
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/"+tt.serverName+"/"+tt.version, nil)
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, tt.expectedCode, w.Code)
			mockDao.AssertExpectations(t)
		})
	}
}
