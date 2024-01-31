// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go
//
// Generated by this command:
//
//	mockgen -source=usecase.go -destination=mocks/mock_interface.go
//

// Package mock__interface is a generated GoMock package.
package mock__interface

import (
	usecase "currency-rate/internal/usecase"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockCurrencyUseCase is a mock of CurrencyUseCase interface.
type MockCurrencyUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockCurrencyUseCaseMockRecorder
}

// MockCurrencyUseCaseMockRecorder is the mock recorder for MockCurrencyUseCase.
type MockCurrencyUseCaseMockRecorder struct {
	mock *MockCurrencyUseCase
}

// NewMockCurrencyUseCase creates a new mock instance.
func NewMockCurrencyUseCase(ctrl *gomock.Controller) *MockCurrencyUseCase {
	mock := &MockCurrencyUseCase{ctrl: ctrl}
	mock.recorder = &MockCurrencyUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCurrencyUseCase) EXPECT() *MockCurrencyUseCaseMockRecorder {
	return m.recorder
}

// GetCurrency mocks base method.
func (m *MockCurrencyUseCase) GetCurrency(date time.Time, curName string) (*usecase.CurrencyDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrency", date, curName)
	ret0, _ := ret[0].(*usecase.CurrencyDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrency indicates an expected call of GetCurrency.
func (mr *MockCurrencyUseCaseMockRecorder) GetCurrency(date, curName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrency", reflect.TypeOf((*MockCurrencyUseCase)(nil).GetCurrency), date, curName)
}

// MockCurrencyGateway is a mock of CurrencyGateway interface.
type MockCurrencyGateway struct {
	ctrl     *gomock.Controller
	recorder *MockCurrencyGatewayMockRecorder
}

// MockCurrencyGatewayMockRecorder is the mock recorder for MockCurrencyGateway.
type MockCurrencyGatewayMockRecorder struct {
	mock *MockCurrencyGateway
}

// NewMockCurrencyGateway creates a new mock instance.
func NewMockCurrencyGateway(ctrl *gomock.Controller) *MockCurrencyGateway {
	mock := &MockCurrencyGateway{ctrl: ctrl}
	mock.recorder = &MockCurrencyGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCurrencyGateway) EXPECT() *MockCurrencyGatewayMockRecorder {
	return m.recorder
}

// GetQuotes mocks base method.
func (m *MockCurrencyGateway) GetQuotes(date string) (*usecase.AllCur, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuotes", date)
	ret0, _ := ret[0].(*usecase.AllCur)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuotes indicates an expected call of GetQuotes.
func (mr *MockCurrencyGatewayMockRecorder) GetQuotes(date any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuotes", reflect.TypeOf((*MockCurrencyGateway)(nil).GetQuotes), date)
}

// MockCurrencyCache is a mock of CurrencyCache interface.
type MockCurrencyCache struct {
	ctrl     *gomock.Controller
	recorder *MockCurrencyCacheMockRecorder
}

// MockCurrencyCacheMockRecorder is the mock recorder for MockCurrencyCache.
type MockCurrencyCacheMockRecorder struct {
	mock *MockCurrencyCache
}

// NewMockCurrencyCache creates a new mock instance.
func NewMockCurrencyCache(ctrl *gomock.Controller) *MockCurrencyCache {
	mock := &MockCurrencyCache{ctrl: ctrl}
	mock.recorder = &MockCurrencyCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCurrencyCache) EXPECT() *MockCurrencyCacheMockRecorder {
	return m.recorder
}

// GetCurrencyDate mocks base method.
func (m *MockCurrencyCache) GetCurrencyDate(date, charName string) (*usecase.CurrencyDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrencyDate", date, charName)
	ret0, _ := ret[0].(*usecase.CurrencyDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrencyDate indicates an expected call of GetCurrencyDate.
func (mr *MockCurrencyCacheMockRecorder) GetCurrencyDate(date, charName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrencyDate", reflect.TypeOf((*MockCurrencyCache)(nil).GetCurrencyDate), date, charName)
}

// InsertCurrencyDate mocks base method.
func (m *MockCurrencyCache) InsertCurrencyDate(value []usecase.CurrencyDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertCurrencyDate", value)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertCurrencyDate indicates an expected call of InsertCurrencyDate.
func (mr *MockCurrencyCacheMockRecorder) InsertCurrencyDate(value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertCurrencyDate", reflect.TypeOf((*MockCurrencyCache)(nil).InsertCurrencyDate), value)
}
