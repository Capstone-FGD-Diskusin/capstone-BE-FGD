// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	users "github.com/dragranzer/capstone-BE-FGD/features/users"
	mock "github.com/stretchr/testify/mock"
)

// Data is an autogenerated mock type for the Data type
type Data struct {
	mock.Mock
}

// CheckEmailPass provides a mock function with given fields: email, pass
func (_m *Data) CheckEmailPass(email string, pass string) (bool, users.Core, error) {
	ret := _m.Called(email, pass)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(email, pass)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 users.Core
	if rf, ok := ret.Get(1).(func(string, string) users.Core); ok {
		r1 = rf(email, pass)
	} else {
		r1 = ret.Get(1).(users.Core)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(email, pass)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CreateUser provides a mock function with given fields: data
func (_m *Data) CreateUser(data users.Core) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.Core) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDataUserbyId provides a mock function with given fields: data
func (_m *Data) DeleteDataUserbyId(data users.Core) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.Core) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Ranking provides a mock function with given fields:
func (_m *Data) Ranking() ([]users.Core, error) {
	ret := _m.Called()

	var r0 []users.Core
	if rf, ok := ret.Get(0).(func() []users.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectAllUser provides a mock function with given fields: data
func (_m *Data) SelectAllUser(data users.Core) ([]users.Core, error) {
	ret := _m.Called(data)

	var r0 []users.Core
	if rf, ok := ret.Get(0).(func(users.Core) []users.Core); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(users.Core) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectDatabyEmail provides a mock function with given fields: data
func (_m *Data) SelectDatabyEmail(data users.Core) (users.Core, error) {
	ret := _m.Called(data)

	var r0 users.Core
	if rf, ok := ret.Get(0).(func(users.Core) users.Core); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(users.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(users.Core) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectDatabyID provides a mock function with given fields: data
func (_m *Data) SelectDatabyID(data users.Core) (users.Core, error) {
	ret := _m.Called(data)

	var r0 users.Core
	if rf, ok := ret.Get(0).(func(users.Core) users.Core); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(users.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(users.Core) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDataUser provides a mock function with given fields: data
func (_m *Data) UpdateDataUser(data users.Core) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.Core) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateFolbyOne provides a mock function with given fields: data
func (_m *Data) UpdateFolbyOne(data users.Core) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.Core) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateFollowingbyOne provides a mock function with given fields: data
func (_m *Data) UpdateFollowingbyOne(data users.Core) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.Core) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateLikebyOne provides a mock function with given fields: data
func (_m *Data) UpdateLikebyOne(data users.Core) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.Core) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateMinFolbyOne provides a mock function with given fields: data
func (_m *Data) UpdateMinFolbyOne(data users.Core) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.Core) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateMinFollowingbyOne provides a mock function with given fields: data
func (_m *Data) UpdateMinFollowingbyOne(data users.Core) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.Core) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateMinLikebyOne provides a mock function with given fields: data
func (_m *Data) UpdateMinLikebyOne(data users.Core) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.Core) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateMinThreadbyOne provides a mock function with given fields: data
func (_m *Data) UpdateMinThreadbyOne(data users.Core) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.Core) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateThreadbyOne provides a mock function with given fields: data
func (_m *Data) UpdateThreadbyOne(data users.Core) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.Core) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUserToModerator provides a mock function with given fields: data
func (_m *Data) UpdateUserToModerator(data users.Core) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.Core) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
