// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	entities "github.com/uiansol/go-clean-architecture/internal/core/entities"
)

// IUserRepository is an autogenerated mock type for the IUserRepository type
type IUserRepository struct {
	mock.Mock
}

// Save provides a mock function with given fields: user
func (_m *IUserRepository) Save(user entities.User) (string, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(entities.User) (string, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(entities.User) string); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(entities.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIUserRepository creates a new instance of IUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IUserRepository {
	mock := &IUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
