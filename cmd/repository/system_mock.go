// Code generated by mockery v2.50.0. DO NOT EDIT.

package repository

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/tech-thinker/gozen/models"
)

// SystemRepoMock is an autogenerated mock type for the SystemRepo type
type SystemRepoMock struct {
	mock.Mock
}

// ExecShell provides a mock function with given fields: command, args
func (_m *SystemRepoMock) ExecShell(command string, args ...string) ([]string, error) {
	_va := make([]interface{}, len(args))
	for _i := range args {
		_va[_i] = args[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, command)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ExecShell")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...string) ([]string, error)); ok {
		return rf(command, args...)
	}
	if rf, ok := ret.Get(0).(func(string, ...string) []string); ok {
		r0 = rf(command, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, ...string) error); ok {
		r1 = rf(command, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecShellRaw provides a mock function with given fields: command, args
func (_m *SystemRepoMock) ExecShellRaw(command string, args ...string) ([]byte, error) {
	_va := make([]interface{}, len(args))
	for _i := range args {
		_va[_i] = args[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, command)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ExecShellRaw")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...string) ([]byte, error)); ok {
		return rf(command, args...)
	}
	if rf, ok := ret.Get(0).(func(string, ...string) []byte); ok {
		r0 = rf(command, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string, ...string) error); ok {
		r1 = rf(command, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Write provides a mock function with given fields: appDir, templatePath, destination, data
func (_m *SystemRepoMock) Write(appDir string, templatePath string, destination string, data interface{}) error {
	ret := _m.Called(appDir, templatePath, destination, data)

	if len(ret) == 0 {
		panic("no return value specified for Write")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, interface{}) error); ok {
		r0 = rf(appDir, templatePath, destination, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteAll provides a mock function with given fields: appDir, fileConfigs, data
func (_m *SystemRepoMock) WriteAll(appDir string, fileConfigs []models.FileConfig, data interface{}) error {
	ret := _m.Called(appDir, fileConfigs, data)

	if len(ret) == 0 {
		panic("no return value specified for WriteAll")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []models.FileConfig, interface{}) error); ok {
		r0 = rf(appDir, fileConfigs, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSystemRepoMock creates a new instance of SystemRepoMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSystemRepoMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *SystemRepoMock {
	mock := &SystemRepoMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
