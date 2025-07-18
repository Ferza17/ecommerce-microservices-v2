// Code generated by mocktail; DO NOT EDIT.

package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

// iAccessControlUseCaseMock mock of IAccessControlUseCase.
type iAccessControlUseCaseMock struct{ mock.Mock }

// NewIAccessControlUseCaseMock creates a new iAccessControlUseCaseMock.
func NewIAccessControlUseCaseMock(tb testing.TB) *iAccessControlUseCaseMock {
	tb.Helper()

	m := &iAccessControlUseCaseMock{}
	m.Mock.Test(tb)

	tb.Cleanup(func() { m.AssertExpectations(tb) })

	return m
}

func (_m *iAccessControlUseCaseMock) IsExcludedHTTP(_ context.Context, requestId string, method string, url string) (bool, error) {
	_ret := _m.Called(requestId, method, url)

	if _rf, ok := _ret.Get(0).(func(string, string, string) (bool, error)); ok {
		return _rf(requestId, method, url)
	}

	_ra0 := _ret.Bool(0)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *iAccessControlUseCaseMock) OnIsExcludedHTTP(requestId string, method string, url string) *iAccessControlUseCaseIsExcludedHTTPCall {
	return &iAccessControlUseCaseIsExcludedHTTPCall{Call: _m.Mock.On("IsExcludedHTTP", requestId, method, url), Parent: _m}
}

func (_m *iAccessControlUseCaseMock) OnIsExcludedHTTPRaw(requestId interface{}, method interface{}, url interface{}) *iAccessControlUseCaseIsExcludedHTTPCall {
	return &iAccessControlUseCaseIsExcludedHTTPCall{Call: _m.Mock.On("IsExcludedHTTP", requestId, method, url), Parent: _m}
}

type iAccessControlUseCaseIsExcludedHTTPCall struct {
	*mock.Call
	Parent *iAccessControlUseCaseMock
}

func (_c *iAccessControlUseCaseIsExcludedHTTPCall) Panic(msg string) *iAccessControlUseCaseIsExcludedHTTPCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *iAccessControlUseCaseIsExcludedHTTPCall) Once() *iAccessControlUseCaseIsExcludedHTTPCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *iAccessControlUseCaseIsExcludedHTTPCall) Twice() *iAccessControlUseCaseIsExcludedHTTPCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *iAccessControlUseCaseIsExcludedHTTPCall) Times(i int) *iAccessControlUseCaseIsExcludedHTTPCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *iAccessControlUseCaseIsExcludedHTTPCall) WaitUntil(w <-chan time.Time) *iAccessControlUseCaseIsExcludedHTTPCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *iAccessControlUseCaseIsExcludedHTTPCall) After(d time.Duration) *iAccessControlUseCaseIsExcludedHTTPCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *iAccessControlUseCaseIsExcludedHTTPCall) Run(fn func(args mock.Arguments)) *iAccessControlUseCaseIsExcludedHTTPCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *iAccessControlUseCaseIsExcludedHTTPCall) Maybe() *iAccessControlUseCaseIsExcludedHTTPCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *iAccessControlUseCaseIsExcludedHTTPCall) TypedReturns(a bool, b error) *iAccessControlUseCaseIsExcludedHTTPCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *iAccessControlUseCaseIsExcludedHTTPCall) ReturnsFn(fn func(string, string, string) (bool, error)) *iAccessControlUseCaseIsExcludedHTTPCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *iAccessControlUseCaseIsExcludedHTTPCall) TypedRun(fn func(string, string, string)) *iAccessControlUseCaseIsExcludedHTTPCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_requestId := args.String(0)
		_method := args.String(1)
		_url := args.String(2)
		fn(_requestId, _method, _url)
	})
	return _c
}

func (_c *iAccessControlUseCaseIsExcludedHTTPCall) OnIsExcludedHTTP(requestId string, method string, url string) *iAccessControlUseCaseIsExcludedHTTPCall {
	return _c.Parent.OnIsExcludedHTTP(requestId, method, url)
}

func (_c *iAccessControlUseCaseIsExcludedHTTPCall) OnIsExcludedRPC(requestId string, url string) *iAccessControlUseCaseIsExcludedRPCCall {
	return _c.Parent.OnIsExcludedRPC(requestId, url)
}

func (_c *iAccessControlUseCaseIsExcludedHTTPCall) OnIsHasHTTPAccess(requestId string, role string, httpMethod string, httpUrl string) *iAccessControlUseCaseIsHasHTTPAccessCall {
	return _c.Parent.OnIsHasHTTPAccess(requestId, role, httpMethod, httpUrl)
}

func (_c *iAccessControlUseCaseIsExcludedHTTPCall) OnIsHasRPCAccess(requestId string, role string, fullMethodName string) *iAccessControlUseCaseIsHasRPCAccessCall {
	return _c.Parent.OnIsHasRPCAccess(requestId, role, fullMethodName)
}

func (_c *iAccessControlUseCaseIsExcludedHTTPCall) OnIsExcludedHTTPRaw(requestId interface{}, method interface{}, url interface{}) *iAccessControlUseCaseIsExcludedHTTPCall {
	return _c.Parent.OnIsExcludedHTTPRaw(requestId, method, url)
}

func (_c *iAccessControlUseCaseIsExcludedHTTPCall) OnIsExcludedRPCRaw(requestId interface{}, url interface{}) *iAccessControlUseCaseIsExcludedRPCCall {
	return _c.Parent.OnIsExcludedRPCRaw(requestId, url)
}

func (_c *iAccessControlUseCaseIsExcludedHTTPCall) OnIsHasHTTPAccessRaw(requestId interface{}, role interface{}, httpMethod interface{}, httpUrl interface{}) *iAccessControlUseCaseIsHasHTTPAccessCall {
	return _c.Parent.OnIsHasHTTPAccessRaw(requestId, role, httpMethod, httpUrl)
}

func (_c *iAccessControlUseCaseIsExcludedHTTPCall) OnIsHasRPCAccessRaw(requestId interface{}, role interface{}, fullMethodName interface{}) *iAccessControlUseCaseIsHasRPCAccessCall {
	return _c.Parent.OnIsHasRPCAccessRaw(requestId, role, fullMethodName)
}

func (_m *iAccessControlUseCaseMock) IsExcludedRPC(_ context.Context, requestId string, url string) (bool, error) {
	_ret := _m.Called(requestId, url)

	if _rf, ok := _ret.Get(0).(func(string, string) (bool, error)); ok {
		return _rf(requestId, url)
	}

	_ra0 := _ret.Bool(0)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *iAccessControlUseCaseMock) OnIsExcludedRPC(requestId string, url string) *iAccessControlUseCaseIsExcludedRPCCall {
	return &iAccessControlUseCaseIsExcludedRPCCall{Call: _m.Mock.On("IsExcludedRPC", requestId, url), Parent: _m}
}

func (_m *iAccessControlUseCaseMock) OnIsExcludedRPCRaw(requestId interface{}, url interface{}) *iAccessControlUseCaseIsExcludedRPCCall {
	return &iAccessControlUseCaseIsExcludedRPCCall{Call: _m.Mock.On("IsExcludedRPC", requestId, url), Parent: _m}
}

type iAccessControlUseCaseIsExcludedRPCCall struct {
	*mock.Call
	Parent *iAccessControlUseCaseMock
}

func (_c *iAccessControlUseCaseIsExcludedRPCCall) Panic(msg string) *iAccessControlUseCaseIsExcludedRPCCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *iAccessControlUseCaseIsExcludedRPCCall) Once() *iAccessControlUseCaseIsExcludedRPCCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *iAccessControlUseCaseIsExcludedRPCCall) Twice() *iAccessControlUseCaseIsExcludedRPCCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *iAccessControlUseCaseIsExcludedRPCCall) Times(i int) *iAccessControlUseCaseIsExcludedRPCCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *iAccessControlUseCaseIsExcludedRPCCall) WaitUntil(w <-chan time.Time) *iAccessControlUseCaseIsExcludedRPCCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *iAccessControlUseCaseIsExcludedRPCCall) After(d time.Duration) *iAccessControlUseCaseIsExcludedRPCCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *iAccessControlUseCaseIsExcludedRPCCall) Run(fn func(args mock.Arguments)) *iAccessControlUseCaseIsExcludedRPCCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *iAccessControlUseCaseIsExcludedRPCCall) Maybe() *iAccessControlUseCaseIsExcludedRPCCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *iAccessControlUseCaseIsExcludedRPCCall) TypedReturns(a bool, b error) *iAccessControlUseCaseIsExcludedRPCCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *iAccessControlUseCaseIsExcludedRPCCall) ReturnsFn(fn func(string, string) (bool, error)) *iAccessControlUseCaseIsExcludedRPCCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *iAccessControlUseCaseIsExcludedRPCCall) TypedRun(fn func(string, string)) *iAccessControlUseCaseIsExcludedRPCCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_requestId := args.String(0)
		_url := args.String(1)
		fn(_requestId, _url)
	})
	return _c
}

func (_c *iAccessControlUseCaseIsExcludedRPCCall) OnIsExcludedHTTP(requestId string, method string, url string) *iAccessControlUseCaseIsExcludedHTTPCall {
	return _c.Parent.OnIsExcludedHTTP(requestId, method, url)
}

func (_c *iAccessControlUseCaseIsExcludedRPCCall) OnIsExcludedRPC(requestId string, url string) *iAccessControlUseCaseIsExcludedRPCCall {
	return _c.Parent.OnIsExcludedRPC(requestId, url)
}

func (_c *iAccessControlUseCaseIsExcludedRPCCall) OnIsHasHTTPAccess(requestId string, role string, httpMethod string, httpUrl string) *iAccessControlUseCaseIsHasHTTPAccessCall {
	return _c.Parent.OnIsHasHTTPAccess(requestId, role, httpMethod, httpUrl)
}

func (_c *iAccessControlUseCaseIsExcludedRPCCall) OnIsHasRPCAccess(requestId string, role string, fullMethodName string) *iAccessControlUseCaseIsHasRPCAccessCall {
	return _c.Parent.OnIsHasRPCAccess(requestId, role, fullMethodName)
}

func (_c *iAccessControlUseCaseIsExcludedRPCCall) OnIsExcludedHTTPRaw(requestId interface{}, method interface{}, url interface{}) *iAccessControlUseCaseIsExcludedHTTPCall {
	return _c.Parent.OnIsExcludedHTTPRaw(requestId, method, url)
}

func (_c *iAccessControlUseCaseIsExcludedRPCCall) OnIsExcludedRPCRaw(requestId interface{}, url interface{}) *iAccessControlUseCaseIsExcludedRPCCall {
	return _c.Parent.OnIsExcludedRPCRaw(requestId, url)
}

func (_c *iAccessControlUseCaseIsExcludedRPCCall) OnIsHasHTTPAccessRaw(requestId interface{}, role interface{}, httpMethod interface{}, httpUrl interface{}) *iAccessControlUseCaseIsHasHTTPAccessCall {
	return _c.Parent.OnIsHasHTTPAccessRaw(requestId, role, httpMethod, httpUrl)
}

func (_c *iAccessControlUseCaseIsExcludedRPCCall) OnIsHasRPCAccessRaw(requestId interface{}, role interface{}, fullMethodName interface{}) *iAccessControlUseCaseIsHasRPCAccessCall {
	return _c.Parent.OnIsHasRPCAccessRaw(requestId, role, fullMethodName)
}

func (_m *iAccessControlUseCaseMock) IsHasHTTPAccess(_ context.Context, requestId string, role string, httpMethod string, httpUrl string) (bool, error) {
	_ret := _m.Called(requestId, role, httpMethod, httpUrl)

	if _rf, ok := _ret.Get(0).(func(string, string, string, string) (bool, error)); ok {
		return _rf(requestId, role, httpMethod, httpUrl)
	}

	_ra0 := _ret.Bool(0)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *iAccessControlUseCaseMock) OnIsHasHTTPAccess(requestId string, role string, httpMethod string, httpUrl string) *iAccessControlUseCaseIsHasHTTPAccessCall {
	return &iAccessControlUseCaseIsHasHTTPAccessCall{Call: _m.Mock.On("IsHasHTTPAccess", requestId, role, httpMethod, httpUrl), Parent: _m}
}

func (_m *iAccessControlUseCaseMock) OnIsHasHTTPAccessRaw(requestId interface{}, role interface{}, httpMethod interface{}, httpUrl interface{}) *iAccessControlUseCaseIsHasHTTPAccessCall {
	return &iAccessControlUseCaseIsHasHTTPAccessCall{Call: _m.Mock.On("IsHasHTTPAccess", requestId, role, httpMethod, httpUrl), Parent: _m}
}

type iAccessControlUseCaseIsHasHTTPAccessCall struct {
	*mock.Call
	Parent *iAccessControlUseCaseMock
}

func (_c *iAccessControlUseCaseIsHasHTTPAccessCall) Panic(msg string) *iAccessControlUseCaseIsHasHTTPAccessCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *iAccessControlUseCaseIsHasHTTPAccessCall) Once() *iAccessControlUseCaseIsHasHTTPAccessCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *iAccessControlUseCaseIsHasHTTPAccessCall) Twice() *iAccessControlUseCaseIsHasHTTPAccessCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *iAccessControlUseCaseIsHasHTTPAccessCall) Times(i int) *iAccessControlUseCaseIsHasHTTPAccessCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *iAccessControlUseCaseIsHasHTTPAccessCall) WaitUntil(w <-chan time.Time) *iAccessControlUseCaseIsHasHTTPAccessCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *iAccessControlUseCaseIsHasHTTPAccessCall) After(d time.Duration) *iAccessControlUseCaseIsHasHTTPAccessCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *iAccessControlUseCaseIsHasHTTPAccessCall) Run(fn func(args mock.Arguments)) *iAccessControlUseCaseIsHasHTTPAccessCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *iAccessControlUseCaseIsHasHTTPAccessCall) Maybe() *iAccessControlUseCaseIsHasHTTPAccessCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *iAccessControlUseCaseIsHasHTTPAccessCall) TypedReturns(a bool, b error) *iAccessControlUseCaseIsHasHTTPAccessCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *iAccessControlUseCaseIsHasHTTPAccessCall) ReturnsFn(fn func(string, string, string, string) (bool, error)) *iAccessControlUseCaseIsHasHTTPAccessCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *iAccessControlUseCaseIsHasHTTPAccessCall) TypedRun(fn func(string, string, string, string)) *iAccessControlUseCaseIsHasHTTPAccessCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_requestId := args.String(0)
		_role := args.String(1)
		_httpMethod := args.String(2)
		_httpUrl := args.String(3)
		fn(_requestId, _role, _httpMethod, _httpUrl)
	})
	return _c
}

func (_c *iAccessControlUseCaseIsHasHTTPAccessCall) OnIsExcludedHTTP(requestId string, method string, url string) *iAccessControlUseCaseIsExcludedHTTPCall {
	return _c.Parent.OnIsExcludedHTTP(requestId, method, url)
}

func (_c *iAccessControlUseCaseIsHasHTTPAccessCall) OnIsExcludedRPC(requestId string, url string) *iAccessControlUseCaseIsExcludedRPCCall {
	return _c.Parent.OnIsExcludedRPC(requestId, url)
}

func (_c *iAccessControlUseCaseIsHasHTTPAccessCall) OnIsHasHTTPAccess(requestId string, role string, httpMethod string, httpUrl string) *iAccessControlUseCaseIsHasHTTPAccessCall {
	return _c.Parent.OnIsHasHTTPAccess(requestId, role, httpMethod, httpUrl)
}

func (_c *iAccessControlUseCaseIsHasHTTPAccessCall) OnIsHasRPCAccess(requestId string, role string, fullMethodName string) *iAccessControlUseCaseIsHasRPCAccessCall {
	return _c.Parent.OnIsHasRPCAccess(requestId, role, fullMethodName)
}

func (_c *iAccessControlUseCaseIsHasHTTPAccessCall) OnIsExcludedHTTPRaw(requestId interface{}, method interface{}, url interface{}) *iAccessControlUseCaseIsExcludedHTTPCall {
	return _c.Parent.OnIsExcludedHTTPRaw(requestId, method, url)
}

func (_c *iAccessControlUseCaseIsHasHTTPAccessCall) OnIsExcludedRPCRaw(requestId interface{}, url interface{}) *iAccessControlUseCaseIsExcludedRPCCall {
	return _c.Parent.OnIsExcludedRPCRaw(requestId, url)
}

func (_c *iAccessControlUseCaseIsHasHTTPAccessCall) OnIsHasHTTPAccessRaw(requestId interface{}, role interface{}, httpMethod interface{}, httpUrl interface{}) *iAccessControlUseCaseIsHasHTTPAccessCall {
	return _c.Parent.OnIsHasHTTPAccessRaw(requestId, role, httpMethod, httpUrl)
}

func (_c *iAccessControlUseCaseIsHasHTTPAccessCall) OnIsHasRPCAccessRaw(requestId interface{}, role interface{}, fullMethodName interface{}) *iAccessControlUseCaseIsHasRPCAccessCall {
	return _c.Parent.OnIsHasRPCAccessRaw(requestId, role, fullMethodName)
}

func (_m *iAccessControlUseCaseMock) IsHasRPCAccess(_ context.Context, requestId string, role string, fullMethodName string) (bool, error) {
	_ret := _m.Called(requestId, role, fullMethodName)

	if _rf, ok := _ret.Get(0).(func(string, string, string) (bool, error)); ok {
		return _rf(requestId, role, fullMethodName)
	}

	_ra0 := _ret.Bool(0)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *iAccessControlUseCaseMock) OnIsHasRPCAccess(requestId string, role string, fullMethodName string) *iAccessControlUseCaseIsHasRPCAccessCall {
	return &iAccessControlUseCaseIsHasRPCAccessCall{Call: _m.Mock.On("IsHasRPCAccess", requestId, role, fullMethodName), Parent: _m}
}

func (_m *iAccessControlUseCaseMock) OnIsHasRPCAccessRaw(requestId interface{}, role interface{}, fullMethodName interface{}) *iAccessControlUseCaseIsHasRPCAccessCall {
	return &iAccessControlUseCaseIsHasRPCAccessCall{Call: _m.Mock.On("IsHasRPCAccess", requestId, role, fullMethodName), Parent: _m}
}

type iAccessControlUseCaseIsHasRPCAccessCall struct {
	*mock.Call
	Parent *iAccessControlUseCaseMock
}

func (_c *iAccessControlUseCaseIsHasRPCAccessCall) Panic(msg string) *iAccessControlUseCaseIsHasRPCAccessCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *iAccessControlUseCaseIsHasRPCAccessCall) Once() *iAccessControlUseCaseIsHasRPCAccessCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *iAccessControlUseCaseIsHasRPCAccessCall) Twice() *iAccessControlUseCaseIsHasRPCAccessCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *iAccessControlUseCaseIsHasRPCAccessCall) Times(i int) *iAccessControlUseCaseIsHasRPCAccessCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *iAccessControlUseCaseIsHasRPCAccessCall) WaitUntil(w <-chan time.Time) *iAccessControlUseCaseIsHasRPCAccessCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *iAccessControlUseCaseIsHasRPCAccessCall) After(d time.Duration) *iAccessControlUseCaseIsHasRPCAccessCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *iAccessControlUseCaseIsHasRPCAccessCall) Run(fn func(args mock.Arguments)) *iAccessControlUseCaseIsHasRPCAccessCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *iAccessControlUseCaseIsHasRPCAccessCall) Maybe() *iAccessControlUseCaseIsHasRPCAccessCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *iAccessControlUseCaseIsHasRPCAccessCall) TypedReturns(a bool, b error) *iAccessControlUseCaseIsHasRPCAccessCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *iAccessControlUseCaseIsHasRPCAccessCall) ReturnsFn(fn func(string, string, string) (bool, error)) *iAccessControlUseCaseIsHasRPCAccessCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *iAccessControlUseCaseIsHasRPCAccessCall) TypedRun(fn func(string, string, string)) *iAccessControlUseCaseIsHasRPCAccessCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_requestId := args.String(0)
		_role := args.String(1)
		_fullMethodName := args.String(2)
		fn(_requestId, _role, _fullMethodName)
	})
	return _c
}

func (_c *iAccessControlUseCaseIsHasRPCAccessCall) OnIsExcludedHTTP(requestId string, method string, url string) *iAccessControlUseCaseIsExcludedHTTPCall {
	return _c.Parent.OnIsExcludedHTTP(requestId, method, url)
}

func (_c *iAccessControlUseCaseIsHasRPCAccessCall) OnIsExcludedRPC(requestId string, url string) *iAccessControlUseCaseIsExcludedRPCCall {
	return _c.Parent.OnIsExcludedRPC(requestId, url)
}

func (_c *iAccessControlUseCaseIsHasRPCAccessCall) OnIsHasHTTPAccess(requestId string, role string, httpMethod string, httpUrl string) *iAccessControlUseCaseIsHasHTTPAccessCall {
	return _c.Parent.OnIsHasHTTPAccess(requestId, role, httpMethod, httpUrl)
}

func (_c *iAccessControlUseCaseIsHasRPCAccessCall) OnIsHasRPCAccess(requestId string, role string, fullMethodName string) *iAccessControlUseCaseIsHasRPCAccessCall {
	return _c.Parent.OnIsHasRPCAccess(requestId, role, fullMethodName)
}

func (_c *iAccessControlUseCaseIsHasRPCAccessCall) OnIsExcludedHTTPRaw(requestId interface{}, method interface{}, url interface{}) *iAccessControlUseCaseIsExcludedHTTPCall {
	return _c.Parent.OnIsExcludedHTTPRaw(requestId, method, url)
}

func (_c *iAccessControlUseCaseIsHasRPCAccessCall) OnIsExcludedRPCRaw(requestId interface{}, url interface{}) *iAccessControlUseCaseIsExcludedRPCCall {
	return _c.Parent.OnIsExcludedRPCRaw(requestId, url)
}

func (_c *iAccessControlUseCaseIsHasRPCAccessCall) OnIsHasHTTPAccessRaw(requestId interface{}, role interface{}, httpMethod interface{}, httpUrl interface{}) *iAccessControlUseCaseIsHasHTTPAccessCall {
	return _c.Parent.OnIsHasHTTPAccessRaw(requestId, role, httpMethod, httpUrl)
}

func (_c *iAccessControlUseCaseIsHasRPCAccessCall) OnIsHasRPCAccessRaw(requestId interface{}, role interface{}, fullMethodName interface{}) *iAccessControlUseCaseIsHasRPCAccessCall {
	return _c.Parent.OnIsHasRPCAccessRaw(requestId, role, fullMethodName)
}
