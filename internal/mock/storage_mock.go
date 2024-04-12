// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/paulbahush/projects/yp/url-shortener/internal/storage/in_file_storage.go

// Package mock_storage is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	"go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// FindByID mocks base method.
func (m *MockRepository) FindByID(ctx context.Context, ID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, ID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockRepositoryMockRecorder) FindByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockRepository)(nil).FindByID), ctx, ID)
}

// Store mocks base method.
func (m *MockRepository) Store(ctx context.Context, Data string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, Data)
	ret0, _ := ret[0].(string)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockRepositoryMockRecorder) Store(ctx, Data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockRepository)(nil).Store), ctx, Data)
}
