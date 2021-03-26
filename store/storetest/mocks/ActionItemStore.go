// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	actionitem "github.com/mattermost/mattermost-server/v5/actionitem"
	mock "github.com/stretchr/testify/mock"
)

// ActionItemStore is an autogenerated mock type for the ActionItemStore type
type ActionItemStore struct {
	mock.Mock
}

// GetForUser provides a mock function with given fields: userid
func (_m *ActionItemStore) GetForUser(userid string) ([]actionitem.ActionItem, error) {
	ret := _m.Called(userid)

	var r0 []actionitem.ActionItem
	if rf, ok := ret.Get(0).(func(string) []actionitem.ActionItem); ok {
		r0 = rf(userid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]actionitem.ActionItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}