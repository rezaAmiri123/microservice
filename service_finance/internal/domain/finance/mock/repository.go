// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package finance_mock is a generated GoMock package.
package finance_mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	finance "github.com/rezaAmiri123/microservice/service_finance/internal/domain/finance"
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

// AddAccountBalance mocks base method.
func (m *MockRepository) AddAccountBalance(ctx context.Context, arg finance.AddAccountBalanceParams) (finance.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAccountBalance", ctx, arg)
	ret0, _ := ret[0].(finance.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAccountBalance indicates an expected call of AddAccountBalance.
func (mr *MockRepositoryMockRecorder) AddAccountBalance(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccountBalance", reflect.TypeOf((*MockRepository)(nil).AddAccountBalance), ctx, arg)
}

// CreateAccount mocks base method.
func (m *MockRepository) CreateAccount(ctx context.Context, arg *finance.CreateAccountParams) (*finance.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", ctx, arg)
	ret0, _ := ret[0].(*finance.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockRepositoryMockRecorder) CreateAccount(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockRepository)(nil).CreateAccount), ctx, arg)
}

// CreateEntry mocks base method.
func (m *MockRepository) CreateEntry(ctx context.Context, arg finance.CreateEntryParams) (finance.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEntry", ctx, arg)
	ret0, _ := ret[0].(finance.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEntry indicates an expected call of CreateEntry.
func (mr *MockRepositoryMockRecorder) CreateEntry(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntry", reflect.TypeOf((*MockRepository)(nil).CreateEntry), ctx, arg)
}

// CreateTransfer mocks base method.
func (m *MockRepository) CreateTransfer(ctx context.Context, arg finance.CreateTransferParams) (finance.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransfer", ctx, arg)
	ret0, _ := ret[0].(finance.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransfer indicates an expected call of CreateTransfer.
func (mr *MockRepositoryMockRecorder) CreateTransfer(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransfer", reflect.TypeOf((*MockRepository)(nil).CreateTransfer), ctx, arg)
}

// ExecTx mocks base method.
func (m *MockRepository) ExecTx(ctx context.Context, fn func(finance.Repository) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecTx", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecTx indicates an expected call of ExecTx.
func (mr *MockRepositoryMockRecorder) ExecTx(ctx, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecTx", reflect.TypeOf((*MockRepository)(nil).ExecTx), ctx, fn)
}

// GetAccountByID mocks base method.
func (m *MockRepository) GetAccountByID(ctx context.Context, accountID uuid.UUID) (*finance.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByID", ctx, accountID)
	ret0, _ := ret[0].(*finance.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByID indicates an expected call of GetAccountByID.
func (mr *MockRepositoryMockRecorder) GetAccountByID(ctx, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByID", reflect.TypeOf((*MockRepository)(nil).GetAccountByID), ctx, accountID)
}

// ListTransfer mocks base method.
func (m *MockRepository) ListTransfer(ctx context.Context, arg finance.ListTransferParams) (*finance.ListTransferResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTransfer", ctx, arg)
	ret0, _ := ret[0].(*finance.ListTransferResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTransfer indicates an expected call of ListTransfer.
func (mr *MockRepositoryMockRecorder) ListTransfer(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTransfer", reflect.TypeOf((*MockRepository)(nil).ListTransfer), ctx, arg)
}

// TransferTx mocks base method.
func (m *MockRepository) TransferTx(ctx context.Context, arg finance.TransferTxParams) (finance.TransferTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferTx", ctx, arg)
	ret0, _ := ret[0].(finance.TransferTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferTx indicates an expected call of TransferTx.
func (mr *MockRepositoryMockRecorder) TransferTx(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferTx", reflect.TypeOf((*MockRepository)(nil).TransferTx), ctx, arg)
}
