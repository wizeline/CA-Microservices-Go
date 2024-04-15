// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	entity "github.com/wizeline/CA-Microservices-Go/internal/entity"
)

// UserSvc is an autogenerated mock type for the UserSvc type
type UserSvc struct {
	mock.Mock
}

// Activate provides a mock function with given fields: id
func (_m *UserSvc) Activate(id uint64) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Activate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChangeEmail provides a mock function with given fields: id, email
func (_m *UserSvc) ChangeEmail(id uint64, email string) error {
	ret := _m.Called(id, email)

	if len(ret) == 0 {
		panic("no return value specified for ChangeEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, string) error); ok {
		r0 = rf(id, email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChangePasswd provides a mock function with given fields: id, passwd
func (_m *UserSvc) ChangePasswd(id uint64, passwd string) error {
	ret := _m.Called(id, passwd)

	if len(ret) == 0 {
		panic("no return value specified for ChangePasswd")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, string) error); ok {
		r0 = rf(id, passwd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: user
func (_m *UserSvc) Create(user entity.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *UserSvc) Delete(id uint64) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: filter, value
func (_m *UserSvc) Find(filter string, value string) ([]entity.User, error) {
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
func (_m *UserSvc) Get(id uint64) (entity.User, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (entity.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint64) entity.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *UserSvc) GetAll() ([]entity.User, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]entity.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []entity.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsActive provides a mock function with given fields: id
func (_m *UserSvc) IsActive(id uint64) (bool, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for IsActive")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (bool, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint64) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: user
func (_m *UserSvc) Update(user entity.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateLogin provides a mock function with given fields: username, passwd
func (_m *UserSvc) ValidateLogin(username string, passwd string) (entity.User, error) {
	ret := _m.Called(username, passwd)

	if len(ret) == 0 {
		panic("no return value specified for ValidateLogin")
	}

	var r0 entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (entity.User, error)); ok {
		return rf(username, passwd)
	}
	if rf, ok := ret.Get(0).(func(string, string) entity.User); ok {
		r0 = rf(username, passwd)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(username, passwd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserSvc creates a new instance of UserSvc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserSvc(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserSvc {
	mock := &UserSvc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}