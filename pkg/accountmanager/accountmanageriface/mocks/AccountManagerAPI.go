// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// AccountManagerAPI is an autogenerated mock type for the AccountManagerAPI type
type AccountManagerAPI struct {
	mock.Mock
}

// Setup provides a mock function with given fields: adminRoleArn
func (_m *AccountManagerAPI) Setup(adminRoleArn string) error {
	ret := _m.Called(adminRoleArn)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(adminRoleArn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
