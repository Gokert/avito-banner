// Code generated by MockGen. DO NOT EDIT.
// Source: core.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "avito-banner/pkg/models"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockICore is a mock of ICore interface.
type MockICore struct {
	ctrl     *gomock.Controller
	recorder *MockICoreMockRecorder
}

// MockICoreMockRecorder is the mock recorder for MockICore.
type MockICoreMockRecorder struct {
	mock *MockICore
}

// NewMockICore creates a new mock instance.
func NewMockICore(ctrl *gomock.Controller) *MockICore {
	mock := &MockICore{ctrl: ctrl}
	mock.recorder = &MockICoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICore) EXPECT() *MockICoreMockRecorder {
	return m.recorder
}

// CheckBanner mocks base method.
func (m *MockICore) CheckBanner(bannerId uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckBanner", bannerId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckBanner indicates an expected call of CheckBanner.
func (mr *MockICoreMockRecorder) CheckBanner(bannerId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckBanner", reflect.TypeOf((*MockICore)(nil).CheckBanner), bannerId)
}

// CreateBanner mocks base method.
func (m *MockICore) CreateBanner(banner *models.BannerRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBanner", banner)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBanner indicates an expected call of CreateBanner.
func (mr *MockICoreMockRecorder) CreateBanner(banner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBanner", reflect.TypeOf((*MockICore)(nil).CreateBanner), banner)
}

// DeleteBanner mocks base method.
func (m *MockICore) DeleteBanner(bannerId uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBanner", bannerId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBanner indicates an expected call of DeleteBanner.
func (mr *MockICoreMockRecorder) DeleteBanner(bannerId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBanner", reflect.TypeOf((*MockICore)(nil).DeleteBanner), bannerId)
}

// GetBanners mocks base method.
func (m *MockICore) GetBanners(tagId, featureId, offset, limit uint64) (*[]models.BannerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBanners", tagId, featureId, offset, limit)
	ret0, _ := ret[0].(*[]models.BannerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBanners indicates an expected call of GetBanners.
func (mr *MockICoreMockRecorder) GetBanners(tagId, featureId, offset, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBanners", reflect.TypeOf((*MockICore)(nil).GetBanners), tagId, featureId, offset, limit)
}

// GetRole mocks base method.
func (m *MockICore) GetRole(ctx context.Context, userId uint64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRole", ctx, userId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRole indicates an expected call of GetRole.
func (mr *MockICoreMockRecorder) GetRole(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*MockICore)(nil).GetRole), ctx, userId)
}

// GetUserBanner mocks base method.
func (m *MockICore) GetUserBanner(tagId, featureId uint64) (*models.UserBanner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserBanner", tagId, featureId)
	ret0, _ := ret[0].(*models.UserBanner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserBanner indicates an expected call of GetUserBanner.
func (mr *MockICoreMockRecorder) GetUserBanner(tagId, featureId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserBanner", reflect.TypeOf((*MockICore)(nil).GetUserBanner), tagId, featureId)
}

// GetUserId mocks base method.
func (m *MockICore) GetUserId(ctx context.Context, sid string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserId", ctx, sid)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserId indicates an expected call of GetUserId.
func (mr *MockICoreMockRecorder) GetUserId(ctx, sid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserId", reflect.TypeOf((*MockICore)(nil).GetUserId), ctx, sid)
}

// UpdateBanner mocks base method.
func (m *MockICore) UpdateBanner(banner *models.BannerRequest) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBanner", banner)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBanner indicates an expected call of UpdateBanner.
func (mr *MockICoreMockRecorder) UpdateBanner(banner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBanner", reflect.TypeOf((*MockICore)(nil).UpdateBanner), banner)
}
