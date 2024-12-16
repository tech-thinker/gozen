// Code generated by mockery v2.50.0. DO NOT EDIT.

package helpers

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/tech-thinker/gozen/models"
)

// MockProjectHelper is an autogenerated mock type for the ProjectHelper type
type MockProjectHelper struct {
	mock.Mock
}

// CreateRestAPI provides a mock function with given fields: project
func (_m *MockProjectHelper) CreateRestAPI(project models.Project) error {
	ret := _m.Called(project)

	if len(ret) == 0 {
		panic("no return value specified for CreateRestAPI")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Project) error); ok {
		r0 = rf(project)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreategRPCAPI provides a mock function with given fields: project
func (_m *MockProjectHelper) CreategRPCAPI(project models.Project) error {
	ret := _m.Called(project)

	if len(ret) == 0 {
		panic("no return value specified for CreategRPCAPI")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Project) error); ok {
		r0 = rf(project)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InitProject provides a mock function with given fields: project
func (_m *MockProjectHelper) InitProject(project models.Project) error {
	ret := _m.Called(project)

	if len(ret) == 0 {
		panic("no return value specified for InitProject")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Project) error); ok {
		r0 = rf(project)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetupConfig provides a mock function with given fields: project
func (_m *MockProjectHelper) SetupConfig(project models.Project) error {
	ret := _m.Called(project)

	if len(ret) == 0 {
		panic("no return value specified for SetupConfig")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Project) error); ok {
		r0 = rf(project)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetupConstants provides a mock function with given fields: project
func (_m *MockProjectHelper) SetupConstants(project models.Project) error {
	ret := _m.Called(project)

	if len(ret) == 0 {
		panic("no return value specified for SetupConstants")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Project) error); ok {
		r0 = rf(project)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetupDocker provides a mock function with given fields: project
func (_m *MockProjectHelper) SetupDocker(project models.Project) error {
	ret := _m.Called(project)

	if len(ret) == 0 {
		panic("no return value specified for SetupDocker")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Project) error); ok {
		r0 = rf(project)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetupEnv provides a mock function with given fields: project
func (_m *MockProjectHelper) SetupEnv(project models.Project) error {
	ret := _m.Called(project)

	if len(ret) == 0 {
		panic("no return value specified for SetupEnv")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Project) error); ok {
		r0 = rf(project)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetupLogger provides a mock function with given fields: project
func (_m *MockProjectHelper) SetupLogger(project models.Project) error {
	ret := _m.Called(project)

	if len(ret) == 0 {
		panic("no return value specified for SetupLogger")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Project) error); ok {
		r0 = rf(project)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetupModel provides a mock function with given fields: project
func (_m *MockProjectHelper) SetupModel(project models.Project) error {
	ret := _m.Called(project)

	if len(ret) == 0 {
		panic("no return value specified for SetupModel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Project) error); ok {
		r0 = rf(project)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetupRepository provides a mock function with given fields: project
func (_m *MockProjectHelper) SetupRepository(project models.Project) error {
	ret := _m.Called(project)

	if len(ret) == 0 {
		panic("no return value specified for SetupRepository")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Project) error); ok {
		r0 = rf(project)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetupService provides a mock function with given fields: project
func (_m *MockProjectHelper) SetupService(project models.Project) error {
	ret := _m.Called(project)

	if len(ret) == 0 {
		panic("no return value specified for SetupService")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Project) error); ok {
		r0 = rf(project)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetupUtils provides a mock function with given fields: project
func (_m *MockProjectHelper) SetupUtils(project models.Project) error {
	ret := _m.Called(project)

	if len(ret) == 0 {
		panic("no return value specified for SetupUtils")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Project) error); ok {
		r0 = rf(project)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockProjectHelper creates a new instance of MockProjectHelper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockProjectHelper(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockProjectHelper {
	mock := &MockProjectHelper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
