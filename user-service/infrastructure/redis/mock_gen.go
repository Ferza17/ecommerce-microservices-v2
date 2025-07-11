// Code generated by mocktail; DO NOT EDIT.

package redis

import (
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/mock"
)

// iRedisInfrastructureMock mock of IRedisInfrastructure.
type iRedisInfrastructureMock struct{ mock.Mock }

// NewIRedisInfrastructureMock creates a new iRedisInfrastructureMock.
func NewIRedisInfrastructureMock(tb testing.TB) *iRedisInfrastructureMock {
	tb.Helper()

	m := &iRedisInfrastructureMock{}
	m.Mock.Test(tb)

	tb.Cleanup(func() { m.AssertExpectations(tb) })

	return m
}

func (_m *iRedisInfrastructureMock) Close() error {
	_ret := _m.Called()

	if _rf, ok := _ret.Get(0).(func() error); ok {
		return _rf()
	}

	_ra0 := _ret.Error(0)

	return _ra0
}

func (_m *iRedisInfrastructureMock) OnClose() *iRedisInfrastructureCloseCall {
	return &iRedisInfrastructureCloseCall{Call: _m.Mock.On("Close"), Parent: _m}
}

func (_m *iRedisInfrastructureMock) OnCloseRaw() *iRedisInfrastructureCloseCall {
	return &iRedisInfrastructureCloseCall{Call: _m.Mock.On("Close"), Parent: _m}
}

type iRedisInfrastructureCloseCall struct {
	*mock.Call
	Parent *iRedisInfrastructureMock
}

func (_c *iRedisInfrastructureCloseCall) Panic(msg string) *iRedisInfrastructureCloseCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *iRedisInfrastructureCloseCall) Once() *iRedisInfrastructureCloseCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *iRedisInfrastructureCloseCall) Twice() *iRedisInfrastructureCloseCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *iRedisInfrastructureCloseCall) Times(i int) *iRedisInfrastructureCloseCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *iRedisInfrastructureCloseCall) WaitUntil(w <-chan time.Time) *iRedisInfrastructureCloseCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *iRedisInfrastructureCloseCall) After(d time.Duration) *iRedisInfrastructureCloseCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *iRedisInfrastructureCloseCall) Run(fn func(args mock.Arguments)) *iRedisInfrastructureCloseCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *iRedisInfrastructureCloseCall) Maybe() *iRedisInfrastructureCloseCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *iRedisInfrastructureCloseCall) TypedReturns(a error) *iRedisInfrastructureCloseCall {
	_c.Call = _c.Return(a)
	return _c
}

func (_c *iRedisInfrastructureCloseCall) ReturnsFn(fn func() error) *iRedisInfrastructureCloseCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *iRedisInfrastructureCloseCall) TypedRun(fn func()) *iRedisInfrastructureCloseCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		fn()
	})
	return _c
}

func (_c *iRedisInfrastructureCloseCall) OnClose() *iRedisInfrastructureCloseCall {
	return _c.Parent.OnClose()
}

func (_c *iRedisInfrastructureCloseCall) OnGetClient() *iRedisInfrastructureGetClientCall {
	return _c.Parent.OnGetClient()
}

func (_c *iRedisInfrastructureCloseCall) OnCloseRaw() *iRedisInfrastructureCloseCall {
	return _c.Parent.OnCloseRaw()
}

func (_c *iRedisInfrastructureCloseCall) OnGetClientRaw() *iRedisInfrastructureGetClientCall {
	return _c.Parent.OnGetClientRaw()
}

func (_m *iRedisInfrastructureMock) GetClient() *redis.Client {
	_ret := _m.Called()

	if _rf, ok := _ret.Get(0).(func() *redis.Client); ok {
		return _rf()
	}

	_ra0, _ := _ret.Get(0).(*redis.Client)

	return _ra0
}

func (_m *iRedisInfrastructureMock) OnGetClient() *iRedisInfrastructureGetClientCall {
	return &iRedisInfrastructureGetClientCall{Call: _m.Mock.On("GetClient"), Parent: _m}
}

func (_m *iRedisInfrastructureMock) OnGetClientRaw() *iRedisInfrastructureGetClientCall {
	return &iRedisInfrastructureGetClientCall{Call: _m.Mock.On("GetClient"), Parent: _m}
}

type iRedisInfrastructureGetClientCall struct {
	*mock.Call
	Parent *iRedisInfrastructureMock
}

func (_c *iRedisInfrastructureGetClientCall) Panic(msg string) *iRedisInfrastructureGetClientCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *iRedisInfrastructureGetClientCall) Once() *iRedisInfrastructureGetClientCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *iRedisInfrastructureGetClientCall) Twice() *iRedisInfrastructureGetClientCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *iRedisInfrastructureGetClientCall) Times(i int) *iRedisInfrastructureGetClientCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *iRedisInfrastructureGetClientCall) WaitUntil(w <-chan time.Time) *iRedisInfrastructureGetClientCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *iRedisInfrastructureGetClientCall) After(d time.Duration) *iRedisInfrastructureGetClientCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *iRedisInfrastructureGetClientCall) Run(fn func(args mock.Arguments)) *iRedisInfrastructureGetClientCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *iRedisInfrastructureGetClientCall) Maybe() *iRedisInfrastructureGetClientCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *iRedisInfrastructureGetClientCall) TypedReturns(a *redis.Client) *iRedisInfrastructureGetClientCall {
	_c.Call = _c.Return(a)
	return _c
}

func (_c *iRedisInfrastructureGetClientCall) ReturnsFn(fn func() *redis.Client) *iRedisInfrastructureGetClientCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *iRedisInfrastructureGetClientCall) TypedRun(fn func()) *iRedisInfrastructureGetClientCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		fn()
	})
	return _c
}

func (_c *iRedisInfrastructureGetClientCall) OnClose() *iRedisInfrastructureCloseCall {
	return _c.Parent.OnClose()
}

func (_c *iRedisInfrastructureGetClientCall) OnGetClient() *iRedisInfrastructureGetClientCall {
	return _c.Parent.OnGetClient()
}

func (_c *iRedisInfrastructureGetClientCall) OnCloseRaw() *iRedisInfrastructureCloseCall {
	return _c.Parent.OnCloseRaw()
}

func (_c *iRedisInfrastructureGetClientCall) OnGetClientRaw() *iRedisInfrastructureGetClientCall {
	return _c.Parent.OnGetClientRaw()
}
