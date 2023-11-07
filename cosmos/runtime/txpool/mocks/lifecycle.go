// Code generated by mockery v2.36.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Lifecycle is an autogenerated mock type for the Lifecycle type
type Lifecycle struct {
	mock.Mock
}

type Lifecycle_Expecter struct {
	mock *mock.Mock
}

func (_m *Lifecycle) EXPECT() *Lifecycle_Expecter {
	return &Lifecycle_Expecter{mock: &_m.Mock}
}

// Start provides a mock function with given fields:
func (_m *Lifecycle) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Lifecycle_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type Lifecycle_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
func (_e *Lifecycle_Expecter) Start() *Lifecycle_Start_Call {
	return &Lifecycle_Start_Call{Call: _e.mock.On("Start")}
}

func (_c *Lifecycle_Start_Call) Run(run func()) *Lifecycle_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Lifecycle_Start_Call) Return(_a0 error) *Lifecycle_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Lifecycle_Start_Call) RunAndReturn(run func() error) *Lifecycle_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *Lifecycle) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Lifecycle_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type Lifecycle_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *Lifecycle_Expecter) Stop() *Lifecycle_Stop_Call {
	return &Lifecycle_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *Lifecycle_Stop_Call) Run(run func()) *Lifecycle_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Lifecycle_Stop_Call) Return(_a0 error) *Lifecycle_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Lifecycle_Stop_Call) RunAndReturn(run func() error) *Lifecycle_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// NewLifecycle creates a new instance of Lifecycle. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLifecycle(t interface {
	mock.TestingT
	Cleanup(func())
}) *Lifecycle {
	mock := &Lifecycle{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
