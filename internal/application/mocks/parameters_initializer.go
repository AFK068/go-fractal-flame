// Code generated by mockery v2.47.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ParametersInitializer is an autogenerated mock type for the ParametersInitializer type
type ParametersInitializer struct {
	mock.Mock
}

type ParametersInitializer_Expecter struct {
	mock *mock.Mock
}

func (_m *ParametersInitializer) EXPECT() *ParametersInitializer_Expecter {
	return &ParametersInitializer_Expecter{mock: &_m.Mock}
}

// InitializeConcurrent provides a mock function with given fields:
func (_m *ParametersInitializer) InitializeConcurrent() (bool, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for InitializeConcurrent")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func() (bool, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParametersInitializer_InitializeConcurrent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InitializeConcurrent'
type ParametersInitializer_InitializeConcurrent_Call struct {
	*mock.Call
}

// InitializeConcurrent is a helper method to define mock.On call
func (_e *ParametersInitializer_Expecter) InitializeConcurrent() *ParametersInitializer_InitializeConcurrent_Call {
	return &ParametersInitializer_InitializeConcurrent_Call{Call: _e.mock.On("InitializeConcurrent")}
}

func (_c *ParametersInitializer_InitializeConcurrent_Call) Run(run func()) *ParametersInitializer_InitializeConcurrent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ParametersInitializer_InitializeConcurrent_Call) Return(_a0 bool, _a1 error) *ParametersInitializer_InitializeConcurrent_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ParametersInitializer_InitializeConcurrent_Call) RunAndReturn(run func() (bool, error)) *ParametersInitializer_InitializeConcurrent_Call {
	_c.Call.Return(run)
	return _c
}

// InitializeGamma provides a mock function with given fields:
func (_m *ParametersInitializer) InitializeGamma() (float64, bool, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for InitializeGamma")
	}

	var r0 float64
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func() (float64, bool, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ParametersInitializer_InitializeGamma_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InitializeGamma'
type ParametersInitializer_InitializeGamma_Call struct {
	*mock.Call
}

// InitializeGamma is a helper method to define mock.On call
func (_e *ParametersInitializer_Expecter) InitializeGamma() *ParametersInitializer_InitializeGamma_Call {
	return &ParametersInitializer_InitializeGamma_Call{Call: _e.mock.On("InitializeGamma")}
}

func (_c *ParametersInitializer_InitializeGamma_Call) Run(run func()) *ParametersInitializer_InitializeGamma_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ParametersInitializer_InitializeGamma_Call) Return(gamma float64, flag bool, err error) *ParametersInitializer_InitializeGamma_Call {
	_c.Call.Return(gamma, flag, err)
	return _c
}

func (_c *ParametersInitializer_InitializeGamma_Call) RunAndReturn(run func() (float64, bool, error)) *ParametersInitializer_InitializeGamma_Call {
	_c.Call.Return(run)
	return _c
}

// InitializeIterations provides a mock function with given fields:
func (_m *ParametersInitializer) InitializeIterations() (float64, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for InitializeIterations")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func() (float64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParametersInitializer_InitializeIterations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InitializeIterations'
type ParametersInitializer_InitializeIterations_Call struct {
	*mock.Call
}

// InitializeIterations is a helper method to define mock.On call
func (_e *ParametersInitializer_Expecter) InitializeIterations() *ParametersInitializer_InitializeIterations_Call {
	return &ParametersInitializer_InitializeIterations_Call{Call: _e.mock.On("InitializeIterations")}
}

func (_c *ParametersInitializer_InitializeIterations_Call) Run(run func()) *ParametersInitializer_InitializeIterations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ParametersInitializer_InitializeIterations_Call) Return(iterations float64, err error) *ParametersInitializer_InitializeIterations_Call {
	_c.Call.Return(iterations, err)
	return _c
}

func (_c *ParametersInitializer_InitializeIterations_Call) RunAndReturn(run func() (float64, error)) *ParametersInitializer_InitializeIterations_Call {
	_c.Call.Return(run)
	return _c
}

// InitializeSymmetry provides a mock function with given fields:
func (_m *ParametersInitializer) InitializeSymmetry() (bool, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for InitializeSymmetry")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func() (bool, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParametersInitializer_InitializeSymmetry_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InitializeSymmetry'
type ParametersInitializer_InitializeSymmetry_Call struct {
	*mock.Call
}

// InitializeSymmetry is a helper method to define mock.On call
func (_e *ParametersInitializer_Expecter) InitializeSymmetry() *ParametersInitializer_InitializeSymmetry_Call {
	return &ParametersInitializer_InitializeSymmetry_Call{Call: _e.mock.On("InitializeSymmetry")}
}

func (_c *ParametersInitializer_InitializeSymmetry_Call) Run(run func()) *ParametersInitializer_InitializeSymmetry_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ParametersInitializer_InitializeSymmetry_Call) Return(_a0 bool, _a1 error) *ParametersInitializer_InitializeSymmetry_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ParametersInitializer_InitializeSymmetry_Call) RunAndReturn(run func() (bool, error)) *ParametersInitializer_InitializeSymmetry_Call {
	_c.Call.Return(run)
	return _c
}

// NewParametersInitializer creates a new instance of ParametersInitializer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewParametersInitializer(t interface {
	mock.TestingT
	Cleanup(func())
}) *ParametersInitializer {
	mock := &ParametersInitializer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
