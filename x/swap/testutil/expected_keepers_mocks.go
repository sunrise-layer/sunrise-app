// Code generated by MockGen. DO NOT EDIT.
// Source: x/swap/types/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	context "context"
	reflect "reflect"

	math "cosmossdk.io/math"
	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	gomock "github.com/golang/mock/gomock"
	types1 "github.com/sunriselayer/sunrise/x/liquiditypool/types"
)

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(arg0 context.Context, arg1 types.AccAddress) types.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(types.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), arg0, arg1)
}

// GetModuleAddress mocks base method.
func (m *MockAccountKeeper) GetModuleAddress(moduleName string) types.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAddress", moduleName)
	ret0, _ := ret[0].(types.AccAddress)
	return ret0
}

// GetModuleAddress indicates an expected call of GetModuleAddress.
func (mr *MockAccountKeeperMockRecorder) GetModuleAddress(moduleName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAddress", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAddress), moduleName)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// IsSendEnabledCoins mocks base method.
func (m *MockBankKeeper) IsSendEnabledCoins(ctx context.Context, coins ...types.Coin) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range coins {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IsSendEnabledCoins", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsSendEnabledCoins indicates an expected call of IsSendEnabledCoins.
func (mr *MockBankKeeperMockRecorder) IsSendEnabledCoins(ctx interface{}, coins ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, coins...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSendEnabledCoins", reflect.TypeOf((*MockBankKeeper)(nil).IsSendEnabledCoins), varargs...)
}

// SendCoins mocks base method.
func (m *MockBankKeeper) SendCoins(ctx context.Context, fromAddr, toAddr types.AccAddress, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoins", ctx, fromAddr, toAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoins indicates an expected call of SendCoins.
func (mr *MockBankKeeperMockRecorder) SendCoins(ctx, fromAddr, toAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoins", reflect.TypeOf((*MockBankKeeper)(nil).SendCoins), ctx, fromAddr, toAddr, amt)
}

// SendCoinsFromModuleToAccount mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr types.AccAddress, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToAccount indicates an expected call of SendCoinsFromModuleToAccount.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToAccount", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToAccount), ctx, senderModule, recipientAddr, amt)
}

// SpendableCoins mocks base method.
func (m *MockBankKeeper) SpendableCoins(arg0 context.Context, arg1 types.AccAddress) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpendableCoins", arg0, arg1)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// SpendableCoins indicates an expected call of SpendableCoins.
func (mr *MockBankKeeperMockRecorder) SpendableCoins(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpendableCoins", reflect.TypeOf((*MockBankKeeper)(nil).SpendableCoins), arg0, arg1)
}

// MockTransferKeeper is a mock of TransferKeeper interface.
type MockTransferKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockTransferKeeperMockRecorder
}

// MockTransferKeeperMockRecorder is the mock recorder for MockTransferKeeper.
type MockTransferKeeperMockRecorder struct {
	mock *MockTransferKeeper
}

// NewMockTransferKeeper creates a new mock instance.
func NewMockTransferKeeper(ctrl *gomock.Controller) *MockTransferKeeper {
	mock := &MockTransferKeeper{ctrl: ctrl}
	mock.recorder = &MockTransferKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransferKeeper) EXPECT() *MockTransferKeeperMockRecorder {
	return m.recorder
}

// DenomPathFromHash mocks base method.
func (m *MockTransferKeeper) DenomPathFromHash(ctx types.Context, denom string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DenomPathFromHash", ctx, denom)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DenomPathFromHash indicates an expected call of DenomPathFromHash.
func (mr *MockTransferKeeperMockRecorder) DenomPathFromHash(ctx, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DenomPathFromHash", reflect.TypeOf((*MockTransferKeeper)(nil).DenomPathFromHash), ctx, denom)
}

// GetTotalEscrowForDenom mocks base method.
func (m *MockTransferKeeper) GetTotalEscrowForDenom(ctx types.Context, denom string) types.Coin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalEscrowForDenom", ctx, denom)
	ret0, _ := ret[0].(types.Coin)
	return ret0
}

// GetTotalEscrowForDenom indicates an expected call of GetTotalEscrowForDenom.
func (mr *MockTransferKeeperMockRecorder) GetTotalEscrowForDenom(ctx, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalEscrowForDenom", reflect.TypeOf((*MockTransferKeeper)(nil).GetTotalEscrowForDenom), ctx, denom)
}

// SetTotalEscrowForDenom mocks base method.
func (m *MockTransferKeeper) SetTotalEscrowForDenom(ctx types.Context, coin types.Coin) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTotalEscrowForDenom", ctx, coin)
}

// SetTotalEscrowForDenom indicates an expected call of SetTotalEscrowForDenom.
func (mr *MockTransferKeeperMockRecorder) SetTotalEscrowForDenom(ctx, coin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTotalEscrowForDenom", reflect.TypeOf((*MockTransferKeeper)(nil).SetTotalEscrowForDenom), ctx, coin)
}

// Transfer mocks base method.
func (m *MockTransferKeeper) Transfer(ctx context.Context, msg *types0.MsgTransfer) (*types0.MsgTransferResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transfer", ctx, msg)
	ret0, _ := ret[0].(*types0.MsgTransferResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Transfer indicates an expected call of Transfer.
func (mr *MockTransferKeeperMockRecorder) Transfer(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transfer", reflect.TypeOf((*MockTransferKeeper)(nil).Transfer), ctx, msg)
}

// MockLiquidityPoolKeeper is a mock of LiquidityPoolKeeper interface.
type MockLiquidityPoolKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockLiquidityPoolKeeperMockRecorder
}

// MockLiquidityPoolKeeperMockRecorder is the mock recorder for MockLiquidityPoolKeeper.
type MockLiquidityPoolKeeperMockRecorder struct {
	mock *MockLiquidityPoolKeeper
}

// NewMockLiquidityPoolKeeper creates a new mock instance.
func NewMockLiquidityPoolKeeper(ctrl *gomock.Controller) *MockLiquidityPoolKeeper {
	mock := &MockLiquidityPoolKeeper{ctrl: ctrl}
	mock.recorder = &MockLiquidityPoolKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLiquidityPoolKeeper) EXPECT() *MockLiquidityPoolKeeperMockRecorder {
	return m.recorder
}

// CalculateResultExactAmountIn mocks base method.
func (m *MockLiquidityPoolKeeper) CalculateResultExactAmountIn(ctx types.Context, pool types1.Pool, tokenIn types.Coin, denomOut string, feeEnabled bool) (math.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateResultExactAmountIn", ctx, pool, tokenIn, denomOut, feeEnabled)
	ret0, _ := ret[0].(math.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalculateResultExactAmountIn indicates an expected call of CalculateResultExactAmountIn.
func (mr *MockLiquidityPoolKeeperMockRecorder) CalculateResultExactAmountIn(ctx, pool, tokenIn, denomOut, feeEnabled interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateResultExactAmountIn", reflect.TypeOf((*MockLiquidityPoolKeeper)(nil).CalculateResultExactAmountIn), ctx, pool, tokenIn, denomOut, feeEnabled)
}

// CalculateResultExactAmountOut mocks base method.
func (m *MockLiquidityPoolKeeper) CalculateResultExactAmountOut(ctx types.Context, pool types1.Pool, tokenOut types.Coin, denomIn string, feeEnabled bool) (math.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateResultExactAmountOut", ctx, pool, tokenOut, denomIn, feeEnabled)
	ret0, _ := ret[0].(math.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalculateResultExactAmountOut indicates an expected call of CalculateResultExactAmountOut.
func (mr *MockLiquidityPoolKeeperMockRecorder) CalculateResultExactAmountOut(ctx, pool, tokenOut, denomIn, feeEnabled interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateResultExactAmountOut", reflect.TypeOf((*MockLiquidityPoolKeeper)(nil).CalculateResultExactAmountOut), ctx, pool, tokenOut, denomIn, feeEnabled)
}

// GetPool mocks base method.
func (m *MockLiquidityPoolKeeper) GetPool(ctx context.Context, id uint64) (types1.Pool, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPool", ctx, id)
	ret0, _ := ret[0].(types1.Pool)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetPool indicates an expected call of GetPool.
func (mr *MockLiquidityPoolKeeperMockRecorder) GetPool(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPool", reflect.TypeOf((*MockLiquidityPoolKeeper)(nil).GetPool), ctx, id)
}

// SwapExactAmountIn mocks base method.
func (m *MockLiquidityPoolKeeper) SwapExactAmountIn(ctx types.Context, sender types.AccAddress, pool types1.Pool, tokenIn types.Coin, denomOut string, feeEnabled bool) (math.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SwapExactAmountIn", ctx, sender, pool, tokenIn, denomOut, feeEnabled)
	ret0, _ := ret[0].(math.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SwapExactAmountIn indicates an expected call of SwapExactAmountIn.
func (mr *MockLiquidityPoolKeeperMockRecorder) SwapExactAmountIn(ctx, sender, pool, tokenIn, denomOut, feeEnabled interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SwapExactAmountIn", reflect.TypeOf((*MockLiquidityPoolKeeper)(nil).SwapExactAmountIn), ctx, sender, pool, tokenIn, denomOut, feeEnabled)
}

// SwapExactAmountOut mocks base method.
func (m *MockLiquidityPoolKeeper) SwapExactAmountOut(ctx types.Context, sender types.AccAddress, pool types1.Pool, tokenOut types.Coin, denomIn string, feeEnabled bool) (math.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SwapExactAmountOut", ctx, sender, pool, tokenOut, denomIn, feeEnabled)
	ret0, _ := ret[0].(math.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SwapExactAmountOut indicates an expected call of SwapExactAmountOut.
func (mr *MockLiquidityPoolKeeperMockRecorder) SwapExactAmountOut(ctx, sender, pool, tokenOut, denomIn, feeEnabled interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SwapExactAmountOut", reflect.TypeOf((*MockLiquidityPoolKeeper)(nil).SwapExactAmountOut), ctx, sender, pool, tokenOut, denomIn, feeEnabled)
}

// MockParamSubspace is a mock of ParamSubspace interface.
type MockParamSubspace struct {
	ctrl     *gomock.Controller
	recorder *MockParamSubspaceMockRecorder
}

// MockParamSubspaceMockRecorder is the mock recorder for MockParamSubspace.
type MockParamSubspaceMockRecorder struct {
	mock *MockParamSubspace
}

// NewMockParamSubspace creates a new mock instance.
func NewMockParamSubspace(ctrl *gomock.Controller) *MockParamSubspace {
	mock := &MockParamSubspace{ctrl: ctrl}
	mock.recorder = &MockParamSubspaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockParamSubspace) EXPECT() *MockParamSubspaceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockParamSubspace) Get(arg0 context.Context, arg1 []byte, arg2 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Get", arg0, arg1, arg2)
}

// Get indicates an expected call of Get.
func (mr *MockParamSubspaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockParamSubspace)(nil).Get), arg0, arg1, arg2)
}

// Set mocks base method.
func (m *MockParamSubspace) Set(arg0 context.Context, arg1 []byte, arg2 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", arg0, arg1, arg2)
}

// Set indicates an expected call of Set.
func (mr *MockParamSubspaceMockRecorder) Set(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockParamSubspace)(nil).Set), arg0, arg1, arg2)
}