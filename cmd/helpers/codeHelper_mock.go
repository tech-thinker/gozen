// Code generated by mockery v2.51.1. DO NOT EDIT.

package helpers

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/tech-thinker/gozen/models"
)

// MockCodeHelper is an autogenerated mock type for the CodeHelper type
type MockCodeHelper struct {
	mock.Mock
}

// GenerateModel provides a mock function with given fields: doc
func (_m *MockCodeHelper) GenerateModel(doc models.Generator) error {
	ret := _m.Called(doc)

	if len(ret) == 0 {
		panic("no return value specified for GenerateModel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Generator) error); ok {
		r0 = rf(doc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockCodeHelper creates a new instance of MockCodeHelper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCodeHelper(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCodeHelper {
	mock := &MockCodeHelper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
