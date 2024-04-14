// Code generated by MockGen. DO NOT EDIT.
// Source: banner_repo.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "avito-banner/pkg/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIRepository is a mock of IRepository interface.
type MockIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryMockRecorder
}

// MockIRepositoryMockRecorder is the mock recorder for MockIRepository.
type MockIRepositoryMockRecorder struct {
	mock *MockIRepository
}

// NewMockIRepository creates a new mock instance.
func NewMockIRepository(ctrl *gomock.Controller) *MockIRepository {
	mock := &MockIRepository{ctrl: ctrl}
	mock.recorder = &MockIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepository) EXPECT() *MockIRepositoryMockRecorder {
	return m.recorder
}

// CheckBanner mocks base method.
func (m *MockIRepository) CheckBanner(bannerId uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckBanner", bannerId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckBanner indicates an expected call of CheckBanner.
func (mr *MockIRepositoryMockRecorder) CheckBanner(bannerId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckBanner", reflect.TypeOf((*MockIRepository)(nil).CheckBanner), bannerId)
}

// CreateBanner mocks base method.
func (m *MockIRepository) CreateBanner(banner *models.BannerRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBanner", banner)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBanner indicates an expected call of CreateBanner.
func (mr *MockIRepositoryMockRecorder) CreateBanner(banner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBanner", reflect.TypeOf((*MockIRepository)(nil).CreateBanner), banner)
}

// DeleteBanner mocks base method.
func (m *MockIRepository) DeleteBanner(bannerId uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBanner", bannerId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBanner indicates an expected call of DeleteBanner.
func (mr *MockIRepositoryMockRecorder) DeleteBanner(bannerId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBanner", reflect.TypeOf((*MockIRepository)(nil).DeleteBanner), bannerId)
}

// GetBanners mocks base method.
func (m *MockIRepository) GetBanners(tagId, featureId, offset, limit uint64) (*[]models.BannerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBanners", tagId, featureId, offset, limit)
	ret0, _ := ret[0].(*[]models.BannerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBanners indicates an expected call of GetBanners.
func (mr *MockIRepositoryMockRecorder) GetBanners(tagId, featureId, offset, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBanners", reflect.TypeOf((*MockIRepository)(nil).GetBanners), tagId, featureId, offset, limit)
}

// GetUserBanner mocks base method.
func (m *MockIRepository) GetUserBanner(tagId, featureId uint64) (*models.UserBanner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserBanner", tagId, featureId)
	ret0, _ := ret[0].(*models.UserBanner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserBanner indicates an expected call of GetUserBanner.
func (mr *MockIRepositoryMockRecorder) GetUserBanner(tagId, featureId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserBanner", reflect.TypeOf((*MockIRepository)(nil).GetUserBanner), tagId, featureId)
}

// UpdateBanner mocks base method.
func (m *MockIRepository) UpdateBanner(banner *models.BannerRequest) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBanner", banner)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBanner indicates an expected call of UpdateBanner.
func (mr *MockIRepositoryMockRecorder) UpdateBanner(banner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBanner", reflect.TypeOf((*MockIRepository)(nil).UpdateBanner), banner)
}