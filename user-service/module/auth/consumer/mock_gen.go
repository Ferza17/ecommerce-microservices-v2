// Code generated by mocktail; DO NOT EDIT.

package consumer

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

// iAuthConsumerMock mock of IAuthConsumer.
type iAuthConsumerMock struct{ mock.Mock }

// NewIAuthConsumerMock creates a new iAuthConsumerMock.
func NewIAuthConsumerMock(tb testing.TB) *iAuthConsumerMock {
	tb.Helper()

	m := &iAuthConsumerMock{}
	m.Mock.Test(tb)

	tb.Cleanup(func() { m.AssertExpectations(tb) })

	return m
}

func (_m *iAuthConsumerMock) UserLogin(_ context.Context) error {
	_ret := _m.Called()

	if _rf, ok := _ret.Get(0).(func() error); ok {
		return _rf()
	}

	_ra0 := _ret.Error(0)

	return _ra0
}

func (_m *iAuthConsumerMock) OnUserLogin() *iAuthConsumerUserLoginCall {
	return &iAuthConsumerUserLoginCall{Call: _m.Mock.On("UserLogin"), Parent: _m}
}

func (_m *iAuthConsumerMock) OnUserLoginRaw() *iAuthConsumerUserLoginCall {
	return &iAuthConsumerUserLoginCall{Call: _m.Mock.On("UserLogin"), Parent: _m}
}

type iAuthConsumerUserLoginCall struct {
	*mock.Call
	Parent *iAuthConsumerMock
}

func (_c *iAuthConsumerUserLoginCall) Panic(msg string) *iAuthConsumerUserLoginCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *iAuthConsumerUserLoginCall) Once() *iAuthConsumerUserLoginCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *iAuthConsumerUserLoginCall) Twice() *iAuthConsumerUserLoginCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *iAuthConsumerUserLoginCall) Times(i int) *iAuthConsumerUserLoginCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *iAuthConsumerUserLoginCall) WaitUntil(w <-chan time.Time) *iAuthConsumerUserLoginCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *iAuthConsumerUserLoginCall) After(d time.Duration) *iAuthConsumerUserLoginCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *iAuthConsumerUserLoginCall) Run(fn func(args mock.Arguments)) *iAuthConsumerUserLoginCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *iAuthConsumerUserLoginCall) Maybe() *iAuthConsumerUserLoginCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *iAuthConsumerUserLoginCall) TypedReturns(a error) *iAuthConsumerUserLoginCall {
	_c.Call = _c.Return(a)
	return _c
}

func (_c *iAuthConsumerUserLoginCall) ReturnsFn(fn func() error) *iAuthConsumerUserLoginCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *iAuthConsumerUserLoginCall) TypedRun(fn func()) *iAuthConsumerUserLoginCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		fn()
	})
	return _c
}

func (_c *iAuthConsumerUserLoginCall) OnUserLogin() *iAuthConsumerUserLoginCall {
	return _c.Parent.OnUserLogin()
}

func (_c *iAuthConsumerUserLoginCall) OnUserLoginRaw() *iAuthConsumerUserLoginCall {
	return _c.Parent.OnUserLoginRaw()
}
