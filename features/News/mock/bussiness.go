// Code generated by mockery 2.9.4. DO NOT EDIT.

package mock

import (
	news "capstone/backend/features/News"

	mock "github.com/stretchr/testify/mock"
)

// Business is an autogenerated mock type for the Business type
type Business struct {
	mock.Mock
}

// Createnews provides a mock function with given fields: accData
func (_m *Business) CreateNews(accData news.NewsCore) error {
	ret := _m.Called(accData)

	var r0 error
	if rf, ok := ret.Get(0).(func(news.NewsCore) error); ok {
		r0 = rf(accData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Getnews provides a mock function with given fields: _a0
func (_m *Business) Getnews(_a0 news.NewsCore) ([]news.NewsCore, error) {
	ret := _m.Called(_a0)

	var r0 []news.NewsCore
	if rf, ok := ret.Get(0).(func(news.NewsCore) []news.NewsCore); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]news.NewsCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(news.NewsCore) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetnewsByID provides a mock function with given fields: id
func (_m *Business) GetnewsByID(id int) (news.NewsCore, error) {
	ret := _m.Called(id)

	var r0 news.NewsCore
	if rf, ok := ret.Get(0).(func(int) news.NewsCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(news.NewsCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Loginnews provides a mock function with given fields: _a0
func (_m *Business) Loginnews(_a0 news.NewsCore) (news.NewsCore, error) {
	ret := _m.Called(_a0)

	var r0 news.NewsCore
	if rf, ok := ret.Get(0).(func(news.NewsCore) news.NewsCore); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(news.NewsCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(news.NewsCore) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
