// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	types "github.com/0xPolygon/zkevm-ethtx-manager/types"
)

// StorageInterface is an autogenerated mock type for the StorageInterface type
type StorageInterface struct {
	mock.Mock
}

type StorageInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *StorageInterface) EXPECT() *StorageInterface_Expecter {
	return &StorageInterface_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: ctx, mTx
func (_m *StorageInterface) Add(ctx context.Context, mTx types.MonitoredTx) error {
	ret := _m.Called(ctx, mTx)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.MonitoredTx) error); ok {
		r0 = rf(ctx, mTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StorageInterface_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type StorageInterface_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - ctx context.Context
//   - mTx types.MonitoredTx
func (_e *StorageInterface_Expecter) Add(ctx interface{}, mTx interface{}) *StorageInterface_Add_Call {
	return &StorageInterface_Add_Call{Call: _e.mock.On("Add", ctx, mTx)}
}

func (_c *StorageInterface_Add_Call) Run(run func(ctx context.Context, mTx types.MonitoredTx)) *StorageInterface_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.MonitoredTx))
	})
	return _c
}

func (_c *StorageInterface_Add_Call) Return(_a0 error) *StorageInterface_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StorageInterface_Add_Call) RunAndReturn(run func(context.Context, types.MonitoredTx) error) *StorageInterface_Add_Call {
	_c.Call.Return(run)
	return _c
}

// Empty provides a mock function with given fields: ctx
func (_m *StorageInterface) Empty(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Empty")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StorageInterface_Empty_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Empty'
type StorageInterface_Empty_Call struct {
	*mock.Call
}

// Empty is a helper method to define mock.On call
//   - ctx context.Context
func (_e *StorageInterface_Expecter) Empty(ctx interface{}) *StorageInterface_Empty_Call {
	return &StorageInterface_Empty_Call{Call: _e.mock.On("Empty", ctx)}
}

func (_c *StorageInterface_Empty_Call) Run(run func(ctx context.Context)) *StorageInterface_Empty_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *StorageInterface_Empty_Call) Return(_a0 error) *StorageInterface_Empty_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StorageInterface_Empty_Call) RunAndReturn(run func(context.Context) error) *StorageInterface_Empty_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, id
func (_m *StorageInterface) Get(ctx context.Context, id common.Hash) (types.MonitoredTx, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 types.MonitoredTx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) (types.MonitoredTx, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) types.MonitoredTx); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(types.MonitoredTx)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StorageInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type StorageInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id common.Hash
func (_e *StorageInterface_Expecter) Get(ctx interface{}, id interface{}) *StorageInterface_Get_Call {
	return &StorageInterface_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *StorageInterface_Get_Call) Run(run func(ctx context.Context, id common.Hash)) *StorageInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Hash))
	})
	return _c
}

func (_c *StorageInterface_Get_Call) Return(_a0 types.MonitoredTx, _a1 error) *StorageInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StorageInterface_Get_Call) RunAndReturn(run func(context.Context, common.Hash) (types.MonitoredTx, error)) *StorageInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByBlock provides a mock function with given fields: ctx, fromBlock, toBlock
func (_m *StorageInterface) GetByBlock(ctx context.Context, fromBlock *uint64, toBlock *uint64) ([]types.MonitoredTx, error) {
	ret := _m.Called(ctx, fromBlock, toBlock)

	if len(ret) == 0 {
		panic("no return value specified for GetByBlock")
	}

	var r0 []types.MonitoredTx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *uint64, *uint64) ([]types.MonitoredTx, error)); ok {
		return rf(ctx, fromBlock, toBlock)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *uint64, *uint64) []types.MonitoredTx); ok {
		r0 = rf(ctx, fromBlock, toBlock)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.MonitoredTx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *uint64, *uint64) error); ok {
		r1 = rf(ctx, fromBlock, toBlock)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StorageInterface_GetByBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByBlock'
type StorageInterface_GetByBlock_Call struct {
	*mock.Call
}

// GetByBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - fromBlock *uint64
//   - toBlock *uint64
func (_e *StorageInterface_Expecter) GetByBlock(ctx interface{}, fromBlock interface{}, toBlock interface{}) *StorageInterface_GetByBlock_Call {
	return &StorageInterface_GetByBlock_Call{Call: _e.mock.On("GetByBlock", ctx, fromBlock, toBlock)}
}

func (_c *StorageInterface_GetByBlock_Call) Run(run func(ctx context.Context, fromBlock *uint64, toBlock *uint64)) *StorageInterface_GetByBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*uint64), args[2].(*uint64))
	})
	return _c
}

func (_c *StorageInterface_GetByBlock_Call) Return(_a0 []types.MonitoredTx, _a1 error) *StorageInterface_GetByBlock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StorageInterface_GetByBlock_Call) RunAndReturn(run func(context.Context, *uint64, *uint64) ([]types.MonitoredTx, error)) *StorageInterface_GetByBlock_Call {
	_c.Call.Return(run)
	return _c
}

// GetByStatus provides a mock function with given fields: ctx, statuses
func (_m *StorageInterface) GetByStatus(ctx context.Context, statuses []types.MonitoredTxStatus) ([]types.MonitoredTx, error) {
	ret := _m.Called(ctx, statuses)

	if len(ret) == 0 {
		panic("no return value specified for GetByStatus")
	}

	var r0 []types.MonitoredTx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []types.MonitoredTxStatus) ([]types.MonitoredTx, error)); ok {
		return rf(ctx, statuses)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []types.MonitoredTxStatus) []types.MonitoredTx); ok {
		r0 = rf(ctx, statuses)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.MonitoredTx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []types.MonitoredTxStatus) error); ok {
		r1 = rf(ctx, statuses)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StorageInterface_GetByStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByStatus'
type StorageInterface_GetByStatus_Call struct {
	*mock.Call
}

// GetByStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - statuses []types.MonitoredTxStatus
func (_e *StorageInterface_Expecter) GetByStatus(ctx interface{}, statuses interface{}) *StorageInterface_GetByStatus_Call {
	return &StorageInterface_GetByStatus_Call{Call: _e.mock.On("GetByStatus", ctx, statuses)}
}

func (_c *StorageInterface_GetByStatus_Call) Run(run func(ctx context.Context, statuses []types.MonitoredTxStatus)) *StorageInterface_GetByStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]types.MonitoredTxStatus))
	})
	return _c
}

func (_c *StorageInterface_GetByStatus_Call) Return(_a0 []types.MonitoredTx, _a1 error) *StorageInterface_GetByStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StorageInterface_GetByStatus_Call) RunAndReturn(run func(context.Context, []types.MonitoredTxStatus) ([]types.MonitoredTx, error)) *StorageInterface_GetByStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Remove provides a mock function with given fields: ctx, id
func (_m *StorageInterface) Remove(ctx context.Context, id common.Hash) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Remove")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StorageInterface_Remove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remove'
type StorageInterface_Remove_Call struct {
	*mock.Call
}

// Remove is a helper method to define mock.On call
//   - ctx context.Context
//   - id common.Hash
func (_e *StorageInterface_Expecter) Remove(ctx interface{}, id interface{}) *StorageInterface_Remove_Call {
	return &StorageInterface_Remove_Call{Call: _e.mock.On("Remove", ctx, id)}
}

func (_c *StorageInterface_Remove_Call) Run(run func(ctx context.Context, id common.Hash)) *StorageInterface_Remove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Hash))
	})
	return _c
}

func (_c *StorageInterface_Remove_Call) Return(_a0 error) *StorageInterface_Remove_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StorageInterface_Remove_Call) RunAndReturn(run func(context.Context, common.Hash) error) *StorageInterface_Remove_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, mTx
func (_m *StorageInterface) Update(ctx context.Context, mTx types.MonitoredTx) error {
	ret := _m.Called(ctx, mTx)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.MonitoredTx) error); ok {
		r0 = rf(ctx, mTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StorageInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type StorageInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - mTx types.MonitoredTx
func (_e *StorageInterface_Expecter) Update(ctx interface{}, mTx interface{}) *StorageInterface_Update_Call {
	return &StorageInterface_Update_Call{Call: _e.mock.On("Update", ctx, mTx)}
}

func (_c *StorageInterface_Update_Call) Run(run func(ctx context.Context, mTx types.MonitoredTx)) *StorageInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.MonitoredTx))
	})
	return _c
}

func (_c *StorageInterface_Update_Call) Return(_a0 error) *StorageInterface_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StorageInterface_Update_Call) RunAndReturn(run func(context.Context, types.MonitoredTx) error) *StorageInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewStorageInterface creates a new instance of StorageInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStorageInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *StorageInterface {
	mock := &StorageInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}