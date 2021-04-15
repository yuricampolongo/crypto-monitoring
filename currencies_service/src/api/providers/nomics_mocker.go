// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/yuricampolongo/crypto-monitoring/currencies_service/src/api/providers (interfaces: CurrenciesInterface)

// Package providers is a generated GoMock package.
package providers

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/yuricampolongo/crypto-monitoring/currencies_service/src/api/domain"
)

// MockCurrenciesInterface is a mock of CurrenciesInterface interface.
type MockCurrenciesInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCurrenciesInterfaceMockRecorder
}

// MockCurrenciesInterfaceMockRecorder is the mock recorder for MockCurrenciesInterface.
type MockCurrenciesInterfaceMockRecorder struct {
	mock *MockCurrenciesInterface
}

// NewMockCurrenciesInterface creates a new mock instance.
func NewMockCurrenciesInterface(ctrl *gomock.Controller) *MockCurrenciesInterface {
	mock := &MockCurrenciesInterface{ctrl: ctrl}
	mock.recorder = &MockCurrenciesInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCurrenciesInterface) EXPECT() *MockCurrenciesInterfaceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockCurrenciesInterface) Get(arg0 domain.CurrencyRequest) (*[]domain.CurrencyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*[]domain.CurrencyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCurrenciesInterfaceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCurrenciesInterface)(nil).Get), arg0)
}