// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Layr-Labs/incredible-squaring-avs/core/chainio (interfaces: AvsReaderer)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/avs_reader.go -package=mocks github.com/Layr-Labs/incredible-squaring-avs/core/chainio AvsReaderer
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	contractECDSAOperatorStateRetriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ECDSAOperatorStateRetriever"
	contractERC20Mock "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/ERC20Mock"
	contractIncredibleSquaringTaskManager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"
	gomock "go.uber.org/mock/gomock"
)

// MockAvsReaderer is a mock of AvsReaderer interface.
type MockAvsReaderer struct {
	ctrl     *gomock.Controller
	recorder *MockAvsReadererMockRecorder
}

// MockAvsReadererMockRecorder is the mock recorder for MockAvsReaderer.
type MockAvsReadererMockRecorder struct {
	mock *MockAvsReaderer
}

// NewMockAvsReaderer creates a new mock instance.
func NewMockAvsReaderer(ctrl *gomock.Controller) *MockAvsReaderer {
	mock := &MockAvsReaderer{ctrl: ctrl}
	mock.recorder = &MockAvsReadererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAvsReaderer) EXPECT() *MockAvsReadererMockRecorder {
	return m.recorder
}

// CheckSignatures mocks base method.
func (m *MockAvsReaderer) CheckSignatures(arg0 context.Context, arg1 [32]byte, arg2 []byte, arg3 uint32, arg4 contractIncredibleSquaringTaskManager.ECDSASignatureCheckerSignerStakeIndicesAndSignatures) (contractIncredibleSquaringTaskManager.ECDSASignatureCheckerQuorumStakeTotals, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSignatures", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(contractIncredibleSquaringTaskManager.ECDSASignatureCheckerQuorumStakeTotals)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckSignatures indicates an expected call of CheckSignatures.
func (mr *MockAvsReadererMockRecorder) CheckSignatures(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSignatures", reflect.TypeOf((*MockAvsReaderer)(nil).CheckSignatures), arg0, arg1, arg2, arg3, arg4)
}

// GetCheckSignaturesIndices mocks base method.
func (m *MockAvsReaderer) GetCheckSignaturesIndices(arg0 *bind.CallOpts, arg1 uint32, arg2 []byte, arg3 []common.Address) (contractECDSAOperatorStateRetriever.ECDSAOperatorStateRetrieverCheckSignaturesIndices, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckSignaturesIndices", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(contractECDSAOperatorStateRetriever.ECDSAOperatorStateRetrieverCheckSignaturesIndices)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckSignaturesIndices indicates an expected call of GetCheckSignaturesIndices.
func (mr *MockAvsReadererMockRecorder) GetCheckSignaturesIndices(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckSignaturesIndices", reflect.TypeOf((*MockAvsReaderer)(nil).GetCheckSignaturesIndices), arg0, arg1, arg2, arg3)
}

// GetErc20Mock mocks base method.
func (m *MockAvsReaderer) GetErc20Mock(arg0 context.Context, arg1 common.Address) (*contractERC20Mock.ContractERC20Mock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetErc20Mock", arg0, arg1)
	ret0, _ := ret[0].(*contractERC20Mock.ContractERC20Mock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetErc20Mock indicates an expected call of GetErc20Mock.
func (mr *MockAvsReadererMockRecorder) GetErc20Mock(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetErc20Mock", reflect.TypeOf((*MockAvsReaderer)(nil).GetErc20Mock), arg0, arg1)
}

// GetOperatorsStakeInQuorumsAtBlock mocks base method.
func (m *MockAvsReaderer) GetOperatorsStakeInQuorumsAtBlock(arg0 *bind.CallOpts, arg1 []byte, arg2 uint32) ([][]contractECDSAOperatorStateRetriever.ECDSAOperatorStateRetrieverOperator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorsStakeInQuorumsAtBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].([][]contractECDSAOperatorStateRetriever.ECDSAOperatorStateRetrieverOperator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorsStakeInQuorumsAtBlock indicates an expected call of GetOperatorsStakeInQuorumsAtBlock.
func (mr *MockAvsReadererMockRecorder) GetOperatorsStakeInQuorumsAtBlock(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorsStakeInQuorumsAtBlock", reflect.TypeOf((*MockAvsReaderer)(nil).GetOperatorsStakeInQuorumsAtBlock), arg0, arg1, arg2)
}

// IsOperatorRegistered mocks base method.
func (m *MockAvsReaderer) IsOperatorRegistered(arg0 *bind.CallOpts, arg1 common.Address) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOperatorRegistered", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsOperatorRegistered indicates an expected call of IsOperatorRegistered.
func (mr *MockAvsReadererMockRecorder) IsOperatorRegistered(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOperatorRegistered", reflect.TypeOf((*MockAvsReaderer)(nil).IsOperatorRegistered), arg0, arg1)
}
