// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	followers "github.com/dragranzer/capstone-BE-FGD/features/followers"
	mock "github.com/stretchr/testify/mock"
)

// Data is an autogenerated mock type for the Data type
type Data struct {
	mock.Mock
}

// DeleteFollow provides a mock function with given fields: data
func (_m *Data) DeleteFollow(data followers.Core) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(followers.Core) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertFollow provides a mock function with given fields: data
func (_m *Data) InsertFollow(data followers.Core) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(followers.Core) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectFollowed provides a mock function with given fields: data
func (_m *Data) SelectFollowed(data followers.Core) ([]followers.Core, error) {
	ret := _m.Called(data)

	var r0 []followers.Core
	if rf, ok := ret.Get(0).(func(followers.Core) []followers.Core); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]followers.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(followers.Core) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectFollowing provides a mock function with given fields: data
func (_m *Data) SelectFollowing(data followers.Core) ([]followers.Core, error) {
	ret := _m.Called(data)

	var r0 []followers.Core
	if rf, ok := ret.Get(0).(func(followers.Core) []followers.Core); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]followers.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(followers.Core) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
