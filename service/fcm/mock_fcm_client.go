// Code generated by mockery v2.14.0. DO NOT EDIT.

package fcm

import (
	go_fcm "github.com/appleboy/go-fcm"
	mock "github.com/stretchr/testify/mock"
)

// mockFcmClient is an autogenerated mock type for the fcmClient type
type mockFcmClient struct {
	mock.Mock
}

// SendWithRetry provides a mock function with given fields: _a0, _a1
func (_m *mockFcmClient) SendWithRetry(_a0 *go_fcm.Message, _a1 int) (*go_fcm.Response, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *go_fcm.Response
	if rf, ok := ret.Get(0).(func(*go_fcm.Message, int) *go_fcm.Response); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*go_fcm.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*go_fcm.Message, int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTnewMockFcmClient interface {
	mock.TestingT
	Cleanup(func())
}

// newMockFcmClient creates a new instance of mockFcmClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockFcmClient(t mockConstructorTestingTnewMockFcmClient) *mockFcmClient {
	mock := &mockFcmClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
