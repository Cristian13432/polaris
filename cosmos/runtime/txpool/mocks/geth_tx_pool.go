// Code generated by mockery v2.36.1. DO NOT EDIT.

package mocks

import (
	common "github.com/ethereum/go-ethereum/common"
	core "github.com/ethereum/go-ethereum/core"

	coretxpool "github.com/ethereum/go-ethereum/core/txpool"

	event "github.com/ethereum/go-ethereum/event"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// GethTxPool is an autogenerated mock type for the GethTxPool type
type GethTxPool struct {
	mock.Mock
}

type GethTxPool_Expecter struct {
	mock *mock.Mock
}

func (_m *GethTxPool) EXPECT() *GethTxPool_Expecter {
	return &GethTxPool_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: _a0, _a1, _a2
func (_m *GethTxPool) Add(_a0 []*types.Transaction, _a1 bool, _a2 bool) []error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []error
	if rf, ok := ret.Get(0).(func([]*types.Transaction, bool, bool) []error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]error)
		}
	}

	return r0
}

// GethTxPool_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type GethTxPool_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - _a0 []*types.Transaction
//   - _a1 bool
//   - _a2 bool
func (_e *GethTxPool_Expecter) Add(_a0 interface{}, _a1 interface{}, _a2 interface{}) *GethTxPool_Add_Call {
	return &GethTxPool_Add_Call{Call: _e.mock.On("Add", _a0, _a1, _a2)}
}

func (_c *GethTxPool_Add_Call) Run(run func(_a0 []*types.Transaction, _a1 bool, _a2 bool)) *GethTxPool_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]*types.Transaction), args[1].(bool), args[2].(bool))
	})
	return _c
}

func (_c *GethTxPool_Add_Call) Return(_a0 []error) *GethTxPool_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GethTxPool_Add_Call) RunAndReturn(run func([]*types.Transaction, bool, bool) []error) *GethTxPool_Add_Call {
	_c.Call.Return(run)
	return _c
}

// Has provides a mock function with given fields: hash
func (_m *GethTxPool) Has(hash common.Hash) bool {
	ret := _m.Called(hash)

	var r0 bool
	if rf, ok := ret.Get(0).(func(common.Hash) bool); ok {
		r0 = rf(hash)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GethTxPool_Has_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Has'
type GethTxPool_Has_Call struct {
	*mock.Call
}

// Has is a helper method to define mock.On call
//   - hash common.Hash
func (_e *GethTxPool_Expecter) Has(hash interface{}) *GethTxPool_Has_Call {
	return &GethTxPool_Has_Call{Call: _e.mock.On("Has", hash)}
}

func (_c *GethTxPool_Has_Call) Run(run func(hash common.Hash)) *GethTxPool_Has_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Hash))
	})
	return _c
}

func (_c *GethTxPool_Has_Call) Return(_a0 bool) *GethTxPool_Has_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GethTxPool_Has_Call) RunAndReturn(run func(common.Hash) bool) *GethTxPool_Has_Call {
	_c.Call.Return(run)
	return _c
}

// Stats provides a mock function with given fields:
func (_m *GethTxPool) Stats() (int, int) {
	ret := _m.Called()

	var r0 int
	var r1 int
	if rf, ok := ret.Get(0).(func() (int, int)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func() int); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// GethTxPool_Stats_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stats'
type GethTxPool_Stats_Call struct {
	*mock.Call
}

// Stats is a helper method to define mock.On call
func (_e *GethTxPool_Expecter) Stats() *GethTxPool_Stats_Call {
	return &GethTxPool_Stats_Call{Call: _e.mock.On("Stats")}
}

func (_c *GethTxPool_Stats_Call) Run(run func()) *GethTxPool_Stats_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GethTxPool_Stats_Call) Return(_a0 int, _a1 int) *GethTxPool_Stats_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GethTxPool_Stats_Call) RunAndReturn(run func() (int, int)) *GethTxPool_Stats_Call {
	_c.Call.Return(run)
	return _c
}

// Status provides a mock function with given fields: hash
func (_m *GethTxPool) Status(hash common.Hash) coretxpool.TxStatus {
	ret := _m.Called(hash)

	var r0 coretxpool.TxStatus
	if rf, ok := ret.Get(0).(func(common.Hash) coretxpool.TxStatus); ok {
		r0 = rf(hash)
	} else {
		r0 = ret.Get(0).(coretxpool.TxStatus)
	}

	return r0
}

// GethTxPool_Status_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Status'
type GethTxPool_Status_Call struct {
	*mock.Call
}

// Status is a helper method to define mock.On call
//   - hash common.Hash
func (_e *GethTxPool_Expecter) Status(hash interface{}) *GethTxPool_Status_Call {
	return &GethTxPool_Status_Call{Call: _e.mock.On("Status", hash)}
}

func (_c *GethTxPool_Status_Call) Run(run func(hash common.Hash)) *GethTxPool_Status_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Hash))
	})
	return _c
}

func (_c *GethTxPool_Status_Call) Return(_a0 coretxpool.TxStatus) *GethTxPool_Status_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GethTxPool_Status_Call) RunAndReturn(run func(common.Hash) coretxpool.TxStatus) *GethTxPool_Status_Call {
	_c.Call.Return(run)
	return _c
}

// SubscribeTransactions provides a mock function with given fields: ch, reorgs
func (_m *GethTxPool) SubscribeTransactions(ch chan<- core.NewTxsEvent, reorgs bool) event.Subscription {
	ret := _m.Called(ch, reorgs)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(chan<- core.NewTxsEvent, bool) event.Subscription); ok {
		r0 = rf(ch, reorgs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	return r0
}

// GethTxPool_SubscribeTransactions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubscribeTransactions'
type GethTxPool_SubscribeTransactions_Call struct {
	*mock.Call
}

// SubscribeTransactions is a helper method to define mock.On call
//   - ch chan<- core.NewTxsEvent
//   - reorgs bool
func (_e *GethTxPool_Expecter) SubscribeTransactions(ch interface{}, reorgs interface{}) *GethTxPool_SubscribeTransactions_Call {
	return &GethTxPool_SubscribeTransactions_Call{Call: _e.mock.On("SubscribeTransactions", ch, reorgs)}
}

func (_c *GethTxPool_SubscribeTransactions_Call) Run(run func(ch chan<- core.NewTxsEvent, reorgs bool)) *GethTxPool_SubscribeTransactions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(chan<- core.NewTxsEvent), args[1].(bool))
	})
	return _c
}

func (_c *GethTxPool_SubscribeTransactions_Call) Return(_a0 event.Subscription) *GethTxPool_SubscribeTransactions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GethTxPool_SubscribeTransactions_Call) RunAndReturn(run func(chan<- core.NewTxsEvent, bool) event.Subscription) *GethTxPool_SubscribeTransactions_Call {
	_c.Call.Return(run)
	return _c
}

// NewGethTxPool creates a new instance of GethTxPool. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGethTxPool(t interface {
	mock.TestingT
	Cleanup(func())
}) *GethTxPool {
	mock := &GethTxPool{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
