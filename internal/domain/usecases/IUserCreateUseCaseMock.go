// Code generated by mockery v2.39.1. DO NOT EDIT.

package usecases

import (
	mock "github.com/stretchr/testify/mock"
)

// IUserCreateUseCaseMock is an autogenerated mock type for the IUserCreateUseCaseMock type
type IUserCreateUseCaseMock struct {
	mock.Mock
}

// Execute provides a mock function with given fields: userCreateInput
func (_m *IUserCreateUseCaseMock) Execute(userCreateInput UserCreateInput) (UserCreateOutput, error) {
	ret := _m.Called(userCreateInput)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 UserCreateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(UserCreateInput) (UserCreateOutput, error)); ok {
		return rf(userCreateInput)
	}
	if rf, ok := ret.Get(0).(func(UserCreateInput) UserCreateOutput); ok {
		r0 = rf(userCreateInput)
	} else {
		r0 = ret.Get(0).(UserCreateOutput)
	}

	if rf, ok := ret.Get(1).(func(UserCreateInput) error); ok {
		r1 = rf(userCreateInput)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIUserCreateUseCase creates a new instance of IUserCreateUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIUserCreateUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *IUserCreateUseCaseMock {
	mock := &IUserCreateUseCaseMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
