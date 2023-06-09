// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package appserver

import (
	time "time"

	mock "github.com/stretchr/testify/mock"

	appnet "github.com/skycoin/skywire/pkg/app/appnet"
	routing "github.com/skycoin/skywire/pkg/routing"
)

// MockRPCIngressClient is an autogenerated mock type for the RPCIngressClient type
type MockRPCIngressClient struct {
	mock.Mock
}

// Accept provides a mock function with given fields: lisID
func (_m *MockRPCIngressClient) Accept(lisID uint16) (uint16, appnet.Addr, error) {
	ret := _m.Called(lisID)

	var r0 uint16
	if rf, ok := ret.Get(0).(func(uint16) uint16); ok {
		r0 = rf(lisID)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	var r1 appnet.Addr
	if rf, ok := ret.Get(1).(func(uint16) appnet.Addr); ok {
		r1 = rf(lisID)
	} else {
		r1 = ret.Get(1).(appnet.Addr)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(uint16) error); ok {
		r2 = rf(lisID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CloseConn provides a mock function with given fields: id
func (_m *MockRPCIngressClient) CloseConn(id uint16) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint16) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CloseListener provides a mock function with given fields: id
func (_m *MockRPCIngressClient) CloseListener(id uint16) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint16) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Dial provides a mock function with given fields: remote
func (_m *MockRPCIngressClient) Dial(remote appnet.Addr) (uint16, routing.Port, error) {
	ret := _m.Called(remote)

	var r0 uint16
	if rf, ok := ret.Get(0).(func(appnet.Addr) uint16); ok {
		r0 = rf(remote)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	var r1 routing.Port
	if rf, ok := ret.Get(1).(func(appnet.Addr) routing.Port); ok {
		r1 = rf(remote)
	} else {
		r1 = ret.Get(1).(routing.Port)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(appnet.Addr) error); ok {
		r2 = rf(remote)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Listen provides a mock function with given fields: local
func (_m *MockRPCIngressClient) Listen(local appnet.Addr) (uint16, error) {
	ret := _m.Called(local)

	var r0 uint16
	if rf, ok := ret.Get(0).(func(appnet.Addr) uint16); ok {
		r0 = rf(local)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(appnet.Addr) error); ok {
		r1 = rf(local)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Read provides a mock function with given fields: connID, b
func (_m *MockRPCIngressClient) Read(connID uint16, b []byte) (int, error) {
	ret := _m.Called(connID, b)

	var r0 int
	if rf, ok := ret.Get(0).(func(uint16, []byte) int); ok {
		r0 = rf(connID, b)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint16, []byte) error); ok {
		r1 = rf(connID, b)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetConnectionDuration provides a mock function with given fields: dur
func (_m *MockRPCIngressClient) SetConnectionDuration(dur int64) error {
	ret := _m.Called(dur)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(dur)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDeadline provides a mock function with given fields: connID, d
func (_m *MockRPCIngressClient) SetDeadline(connID uint16, d time.Time) error {
	ret := _m.Called(connID, d)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint16, time.Time) error); ok {
		r0 = rf(connID, d)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDetailedStatus provides a mock function with given fields: status
func (_m *MockRPCIngressClient) SetDetailedStatus(status string) error {
	ret := _m.Called(status)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetError provides a mock function with given fields: appErr
func (_m *MockRPCIngressClient) SetError(appErr string) error {
	ret := _m.Called(appErr)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(appErr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDetailedStatus provides a mock function with given fields: status
func (_m *MockRPCIngressClient) SetAppPort(port routing.Port) error {
	ret := _m.Called(port)

	var r0 error
	if rf, ok := ret.Get(0).(func(routing.Port) error); ok {
		r0 = rf(port)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetReadDeadline provides a mock function with given fields: connID, d
func (_m *MockRPCIngressClient) SetReadDeadline(connID uint16, d time.Time) error {
	ret := _m.Called(connID, d)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint16, time.Time) error); ok {
		r0 = rf(connID, d)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetWriteDeadline provides a mock function with given fields: connID, d
func (_m *MockRPCIngressClient) SetWriteDeadline(connID uint16, d time.Time) error {
	ret := _m.Called(connID, d)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint16, time.Time) error); ok {
		r0 = rf(connID, d)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Write provides a mock function with given fields: connID, b
func (_m *MockRPCIngressClient) Write(connID uint16, b []byte) (int, error) {
	ret := _m.Called(connID, b)

	var r0 int
	if rf, ok := ret.Get(0).(func(uint16, []byte) int); ok {
		r0 = rf(connID, b)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint16, []byte) error); ok {
		r1 = rf(connID, b)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
