// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	entity "github.com/wizeline/CA-Microservices-Go/internal/domain/entity"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// Add provides a mock function with given fields: user
func (_m *UserService) Add(user entity.User) (int, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.User) (int, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(entity.User) int); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(entity.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChangeEmail provides a mock function with given fields: id, email
func (_m *UserService) ChangeEmail(id int, email string) error {
	ret := _m.Called(id, email)

	if len(ret) == 0 {
		panic("no return value specified for ChangeEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string) error); ok {
		r0 = rf(id, email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *UserService) Delete(id int) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: filter, value
func (_m *UserService) Find(filter string, value string) ([]entity.User, error) {
	ret := _m.Called(filter, value)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 []entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]entity.User, error)); ok {
		return rf(filter, value)
	}
	if rf, ok := ret.Get(0).(func(string, string) []entity.User); ok {
		r0 = rf(filter, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(filter, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *UserService) Get(id int) (entity.User, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (entity.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) entity.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsActive provides a mock function with given fields: id
func (_m *UserService) IsActive(id int) (bool, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for IsActive")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (bool, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, data
func (_m *UserService) Update(id int, data entity.User) error {
	ret := _m.Called(id, data)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, entity.User) error); ok {
		r0 = rf(id, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateLogin provides a mock function with given fields: username, password
func (_m *UserService) ValidateLogin(username string, password string) (entity.User, error) {
	ret := _m.Called(username, password)

	if len(ret) == 0 {
		panic("no return value specified for ValidateLogin")
	}

	var r0 entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (entity.User, error)); ok {
		return rf(username, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) entity.User); ok {
		r0 = rf(username, password)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
