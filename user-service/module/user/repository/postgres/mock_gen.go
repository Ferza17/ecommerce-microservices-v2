// Code generated by mocktail; DO NOT EDIT.

package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// iUserPostgresqlRepositoryMock mock of IUserPostgresqlRepository.
type iUserPostgresqlRepositoryMock struct{ mock.Mock }

// NewIUserPostgresqlRepositoryMock creates a new iUserPostgresqlRepositoryMock.
func NewIUserPostgresqlRepositoryMock(tb testing.TB) *iUserPostgresqlRepositoryMock {
	tb.Helper()

	m := &iUserPostgresqlRepositoryMock{}
	m.Mock.Test(tb)

	tb.Cleanup(func() { m.AssertExpectations(tb) })

	return m
}

func (_m *iUserPostgresqlRepositoryMock) CreateUser(_ context.Context, requestId string, req *orm.User, tx *gorm.DB) (*orm.User, error) {
	_ret := _m.Called(requestId, req, tx)

	if _rf, ok := _ret.Get(0).(func(string, *orm.User, *gorm.DB) (*orm.User, error)); ok {
		return _rf(requestId, req, tx)
	}

	_ra0, _ := _ret.Get(0).(*orm.User)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *iUserPostgresqlRepositoryMock) OnCreateUser(requestId string, req *orm.User, tx *gorm.DB) *iUserPostgresqlRepositoryCreateUserCall {
	return &iUserPostgresqlRepositoryCreateUserCall{Call: _m.Mock.On("CreateUser", requestId, req, tx), Parent: _m}
}

func (_m *iUserPostgresqlRepositoryMock) OnCreateUserRaw(requestId interface{}, req interface{}, tx interface{}) *iUserPostgresqlRepositoryCreateUserCall {
	return &iUserPostgresqlRepositoryCreateUserCall{Call: _m.Mock.On("CreateUser", requestId, req, tx), Parent: _m}
}

type iUserPostgresqlRepositoryCreateUserCall struct {
	*mock.Call
	Parent *iUserPostgresqlRepositoryMock
}

func (_c *iUserPostgresqlRepositoryCreateUserCall) Panic(msg string) *iUserPostgresqlRepositoryCreateUserCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *iUserPostgresqlRepositoryCreateUserCall) Once() *iUserPostgresqlRepositoryCreateUserCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *iUserPostgresqlRepositoryCreateUserCall) Twice() *iUserPostgresqlRepositoryCreateUserCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *iUserPostgresqlRepositoryCreateUserCall) Times(i int) *iUserPostgresqlRepositoryCreateUserCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *iUserPostgresqlRepositoryCreateUserCall) WaitUntil(w <-chan time.Time) *iUserPostgresqlRepositoryCreateUserCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *iUserPostgresqlRepositoryCreateUserCall) After(d time.Duration) *iUserPostgresqlRepositoryCreateUserCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *iUserPostgresqlRepositoryCreateUserCall) Run(fn func(args mock.Arguments)) *iUserPostgresqlRepositoryCreateUserCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *iUserPostgresqlRepositoryCreateUserCall) Maybe() *iUserPostgresqlRepositoryCreateUserCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *iUserPostgresqlRepositoryCreateUserCall) TypedReturns(a *orm.User, b error) *iUserPostgresqlRepositoryCreateUserCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *iUserPostgresqlRepositoryCreateUserCall) ReturnsFn(fn func(string, *orm.User, *gorm.DB) (*orm.User, error)) *iUserPostgresqlRepositoryCreateUserCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *iUserPostgresqlRepositoryCreateUserCall) TypedRun(fn func(string, *orm.User, *gorm.DB)) *iUserPostgresqlRepositoryCreateUserCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_requestId := args.String(0)
		_req, _ := args.Get(1).(*orm.User)
		_tx, _ := args.Get(2).(*gorm.DB)
		fn(_requestId, _req, _tx)
	})
	return _c
}

func (_c *iUserPostgresqlRepositoryCreateUserCall) OnCreateUser(requestId string, req *orm.User, tx *gorm.DB) *iUserPostgresqlRepositoryCreateUserCall {
	return _c.Parent.OnCreateUser(requestId, req, tx)
}

func (_c *iUserPostgresqlRepositoryCreateUserCall) OnFindUserByEmail(requestId string, email string, tx *gorm.DB) *iUserPostgresqlRepositoryFindUserByEmailCall {
	return _c.Parent.OnFindUserByEmail(requestId, email, tx)
}

func (_c *iUserPostgresqlRepositoryCreateUserCall) OnFindUserById(requestId string, id string, tx *gorm.DB) *iUserPostgresqlRepositoryFindUserByIdCall {
	return _c.Parent.OnFindUserById(requestId, id, tx)
}

func (_c *iUserPostgresqlRepositoryCreateUserCall) OnUpdateUserById(requestId string, req *orm.User, tx *gorm.DB) *iUserPostgresqlRepositoryUpdateUserByIdCall {
	return _c.Parent.OnUpdateUserById(requestId, req, tx)
}

func (_c *iUserPostgresqlRepositoryCreateUserCall) OnCreateUserRaw(requestId interface{}, req interface{}, tx interface{}) *iUserPostgresqlRepositoryCreateUserCall {
	return _c.Parent.OnCreateUserRaw(requestId, req, tx)
}

func (_c *iUserPostgresqlRepositoryCreateUserCall) OnFindUserByEmailRaw(requestId interface{}, email interface{}, tx interface{}) *iUserPostgresqlRepositoryFindUserByEmailCall {
	return _c.Parent.OnFindUserByEmailRaw(requestId, email, tx)
}

func (_c *iUserPostgresqlRepositoryCreateUserCall) OnFindUserByIdRaw(requestId interface{}, id interface{}, tx interface{}) *iUserPostgresqlRepositoryFindUserByIdCall {
	return _c.Parent.OnFindUserByIdRaw(requestId, id, tx)
}

func (_c *iUserPostgresqlRepositoryCreateUserCall) OnUpdateUserByIdRaw(requestId interface{}, req interface{}, tx interface{}) *iUserPostgresqlRepositoryUpdateUserByIdCall {
	return _c.Parent.OnUpdateUserByIdRaw(requestId, req, tx)
}

func (_m *iUserPostgresqlRepositoryMock) FindUserByEmail(_ context.Context, requestId string, email string, tx *gorm.DB) (*orm.User, error) {
	_ret := _m.Called(requestId, email, tx)

	if _rf, ok := _ret.Get(0).(func(string, string, *gorm.DB) (*orm.User, error)); ok {
		return _rf(requestId, email, tx)
	}

	_ra0, _ := _ret.Get(0).(*orm.User)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *iUserPostgresqlRepositoryMock) OnFindUserByEmail(requestId string, email string, tx *gorm.DB) *iUserPostgresqlRepositoryFindUserByEmailCall {
	return &iUserPostgresqlRepositoryFindUserByEmailCall{Call: _m.Mock.On("FindUserByEmail", requestId, email, tx), Parent: _m}
}

func (_m *iUserPostgresqlRepositoryMock) OnFindUserByEmailRaw(requestId interface{}, email interface{}, tx interface{}) *iUserPostgresqlRepositoryFindUserByEmailCall {
	return &iUserPostgresqlRepositoryFindUserByEmailCall{Call: _m.Mock.On("FindUserByEmail", requestId, email, tx), Parent: _m}
}

type iUserPostgresqlRepositoryFindUserByEmailCall struct {
	*mock.Call
	Parent *iUserPostgresqlRepositoryMock
}

func (_c *iUserPostgresqlRepositoryFindUserByEmailCall) Panic(msg string) *iUserPostgresqlRepositoryFindUserByEmailCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *iUserPostgresqlRepositoryFindUserByEmailCall) Once() *iUserPostgresqlRepositoryFindUserByEmailCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *iUserPostgresqlRepositoryFindUserByEmailCall) Twice() *iUserPostgresqlRepositoryFindUserByEmailCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *iUserPostgresqlRepositoryFindUserByEmailCall) Times(i int) *iUserPostgresqlRepositoryFindUserByEmailCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *iUserPostgresqlRepositoryFindUserByEmailCall) WaitUntil(w <-chan time.Time) *iUserPostgresqlRepositoryFindUserByEmailCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *iUserPostgresqlRepositoryFindUserByEmailCall) After(d time.Duration) *iUserPostgresqlRepositoryFindUserByEmailCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *iUserPostgresqlRepositoryFindUserByEmailCall) Run(fn func(args mock.Arguments)) *iUserPostgresqlRepositoryFindUserByEmailCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *iUserPostgresqlRepositoryFindUserByEmailCall) Maybe() *iUserPostgresqlRepositoryFindUserByEmailCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *iUserPostgresqlRepositoryFindUserByEmailCall) TypedReturns(a *orm.User, b error) *iUserPostgresqlRepositoryFindUserByEmailCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *iUserPostgresqlRepositoryFindUserByEmailCall) ReturnsFn(fn func(string, string, *gorm.DB) (*orm.User, error)) *iUserPostgresqlRepositoryFindUserByEmailCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *iUserPostgresqlRepositoryFindUserByEmailCall) TypedRun(fn func(string, string, *gorm.DB)) *iUserPostgresqlRepositoryFindUserByEmailCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_requestId := args.String(0)
		_email := args.String(1)
		_tx, _ := args.Get(2).(*gorm.DB)
		fn(_requestId, _email, _tx)
	})
	return _c
}

func (_c *iUserPostgresqlRepositoryFindUserByEmailCall) OnCreateUser(requestId string, req *orm.User, tx *gorm.DB) *iUserPostgresqlRepositoryCreateUserCall {
	return _c.Parent.OnCreateUser(requestId, req, tx)
}

func (_c *iUserPostgresqlRepositoryFindUserByEmailCall) OnFindUserByEmail(requestId string, email string, tx *gorm.DB) *iUserPostgresqlRepositoryFindUserByEmailCall {
	return _c.Parent.OnFindUserByEmail(requestId, email, tx)
}

func (_c *iUserPostgresqlRepositoryFindUserByEmailCall) OnFindUserById(requestId string, id string, tx *gorm.DB) *iUserPostgresqlRepositoryFindUserByIdCall {
	return _c.Parent.OnFindUserById(requestId, id, tx)
}

func (_c *iUserPostgresqlRepositoryFindUserByEmailCall) OnUpdateUserById(requestId string, req *orm.User, tx *gorm.DB) *iUserPostgresqlRepositoryUpdateUserByIdCall {
	return _c.Parent.OnUpdateUserById(requestId, req, tx)
}

func (_c *iUserPostgresqlRepositoryFindUserByEmailCall) OnCreateUserRaw(requestId interface{}, req interface{}, tx interface{}) *iUserPostgresqlRepositoryCreateUserCall {
	return _c.Parent.OnCreateUserRaw(requestId, req, tx)
}

func (_c *iUserPostgresqlRepositoryFindUserByEmailCall) OnFindUserByEmailRaw(requestId interface{}, email interface{}, tx interface{}) *iUserPostgresqlRepositoryFindUserByEmailCall {
	return _c.Parent.OnFindUserByEmailRaw(requestId, email, tx)
}

func (_c *iUserPostgresqlRepositoryFindUserByEmailCall) OnFindUserByIdRaw(requestId interface{}, id interface{}, tx interface{}) *iUserPostgresqlRepositoryFindUserByIdCall {
	return _c.Parent.OnFindUserByIdRaw(requestId, id, tx)
}

func (_c *iUserPostgresqlRepositoryFindUserByEmailCall) OnUpdateUserByIdRaw(requestId interface{}, req interface{}, tx interface{}) *iUserPostgresqlRepositoryUpdateUserByIdCall {
	return _c.Parent.OnUpdateUserByIdRaw(requestId, req, tx)
}

func (_m *iUserPostgresqlRepositoryMock) FindUserById(_ context.Context, requestId string, id string, tx *gorm.DB) (*orm.User, error) {
	_ret := _m.Called(requestId, id, tx)

	if _rf, ok := _ret.Get(0).(func(string, string, *gorm.DB) (*orm.User, error)); ok {
		return _rf(requestId, id, tx)
	}

	_ra0, _ := _ret.Get(0).(*orm.User)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *iUserPostgresqlRepositoryMock) OnFindUserById(requestId string, id string, tx *gorm.DB) *iUserPostgresqlRepositoryFindUserByIdCall {
	return &iUserPostgresqlRepositoryFindUserByIdCall{Call: _m.Mock.On("FindUserById", requestId, id, tx), Parent: _m}
}

func (_m *iUserPostgresqlRepositoryMock) OnFindUserByIdRaw(requestId interface{}, id interface{}, tx interface{}) *iUserPostgresqlRepositoryFindUserByIdCall {
	return &iUserPostgresqlRepositoryFindUserByIdCall{Call: _m.Mock.On("FindUserById", requestId, id, tx), Parent: _m}
}

type iUserPostgresqlRepositoryFindUserByIdCall struct {
	*mock.Call
	Parent *iUserPostgresqlRepositoryMock
}

func (_c *iUserPostgresqlRepositoryFindUserByIdCall) Panic(msg string) *iUserPostgresqlRepositoryFindUserByIdCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *iUserPostgresqlRepositoryFindUserByIdCall) Once() *iUserPostgresqlRepositoryFindUserByIdCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *iUserPostgresqlRepositoryFindUserByIdCall) Twice() *iUserPostgresqlRepositoryFindUserByIdCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *iUserPostgresqlRepositoryFindUserByIdCall) Times(i int) *iUserPostgresqlRepositoryFindUserByIdCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *iUserPostgresqlRepositoryFindUserByIdCall) WaitUntil(w <-chan time.Time) *iUserPostgresqlRepositoryFindUserByIdCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *iUserPostgresqlRepositoryFindUserByIdCall) After(d time.Duration) *iUserPostgresqlRepositoryFindUserByIdCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *iUserPostgresqlRepositoryFindUserByIdCall) Run(fn func(args mock.Arguments)) *iUserPostgresqlRepositoryFindUserByIdCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *iUserPostgresqlRepositoryFindUserByIdCall) Maybe() *iUserPostgresqlRepositoryFindUserByIdCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *iUserPostgresqlRepositoryFindUserByIdCall) TypedReturns(a *orm.User, b error) *iUserPostgresqlRepositoryFindUserByIdCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *iUserPostgresqlRepositoryFindUserByIdCall) ReturnsFn(fn func(string, string, *gorm.DB) (*orm.User, error)) *iUserPostgresqlRepositoryFindUserByIdCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *iUserPostgresqlRepositoryFindUserByIdCall) TypedRun(fn func(string, string, *gorm.DB)) *iUserPostgresqlRepositoryFindUserByIdCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_requestId := args.String(0)
		_id := args.String(1)
		_tx, _ := args.Get(2).(*gorm.DB)
		fn(_requestId, _id, _tx)
	})
	return _c
}

func (_c *iUserPostgresqlRepositoryFindUserByIdCall) OnCreateUser(requestId string, req *orm.User, tx *gorm.DB) *iUserPostgresqlRepositoryCreateUserCall {
	return _c.Parent.OnCreateUser(requestId, req, tx)
}

func (_c *iUserPostgresqlRepositoryFindUserByIdCall) OnFindUserByEmail(requestId string, email string, tx *gorm.DB) *iUserPostgresqlRepositoryFindUserByEmailCall {
	return _c.Parent.OnFindUserByEmail(requestId, email, tx)
}

func (_c *iUserPostgresqlRepositoryFindUserByIdCall) OnFindUserById(requestId string, id string, tx *gorm.DB) *iUserPostgresqlRepositoryFindUserByIdCall {
	return _c.Parent.OnFindUserById(requestId, id, tx)
}

func (_c *iUserPostgresqlRepositoryFindUserByIdCall) OnUpdateUserById(requestId string, req *orm.User, tx *gorm.DB) *iUserPostgresqlRepositoryUpdateUserByIdCall {
	return _c.Parent.OnUpdateUserById(requestId, req, tx)
}

func (_c *iUserPostgresqlRepositoryFindUserByIdCall) OnCreateUserRaw(requestId interface{}, req interface{}, tx interface{}) *iUserPostgresqlRepositoryCreateUserCall {
	return _c.Parent.OnCreateUserRaw(requestId, req, tx)
}

func (_c *iUserPostgresqlRepositoryFindUserByIdCall) OnFindUserByEmailRaw(requestId interface{}, email interface{}, tx interface{}) *iUserPostgresqlRepositoryFindUserByEmailCall {
	return _c.Parent.OnFindUserByEmailRaw(requestId, email, tx)
}

func (_c *iUserPostgresqlRepositoryFindUserByIdCall) OnFindUserByIdRaw(requestId interface{}, id interface{}, tx interface{}) *iUserPostgresqlRepositoryFindUserByIdCall {
	return _c.Parent.OnFindUserByIdRaw(requestId, id, tx)
}

func (_c *iUserPostgresqlRepositoryFindUserByIdCall) OnUpdateUserByIdRaw(requestId interface{}, req interface{}, tx interface{}) *iUserPostgresqlRepositoryUpdateUserByIdCall {
	return _c.Parent.OnUpdateUserByIdRaw(requestId, req, tx)
}

func (_m *iUserPostgresqlRepositoryMock) UpdateUserById(_ context.Context, requestId string, req *orm.User, tx *gorm.DB) (*orm.User, error) {
	_ret := _m.Called(requestId, req, tx)

	if _rf, ok := _ret.Get(0).(func(string, *orm.User, *gorm.DB) (*orm.User, error)); ok {
		return _rf(requestId, req, tx)
	}

	_ra0, _ := _ret.Get(0).(*orm.User)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *iUserPostgresqlRepositoryMock) OnUpdateUserById(requestId string, req *orm.User, tx *gorm.DB) *iUserPostgresqlRepositoryUpdateUserByIdCall {
	return &iUserPostgresqlRepositoryUpdateUserByIdCall{Call: _m.Mock.On("UpdateUserById", requestId, req, tx), Parent: _m}
}

func (_m *iUserPostgresqlRepositoryMock) OnUpdateUserByIdRaw(requestId interface{}, req interface{}, tx interface{}) *iUserPostgresqlRepositoryUpdateUserByIdCall {
	return &iUserPostgresqlRepositoryUpdateUserByIdCall{Call: _m.Mock.On("UpdateUserById", requestId, req, tx), Parent: _m}
}

type iUserPostgresqlRepositoryUpdateUserByIdCall struct {
	*mock.Call
	Parent *iUserPostgresqlRepositoryMock
}

func (_c *iUserPostgresqlRepositoryUpdateUserByIdCall) Panic(msg string) *iUserPostgresqlRepositoryUpdateUserByIdCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *iUserPostgresqlRepositoryUpdateUserByIdCall) Once() *iUserPostgresqlRepositoryUpdateUserByIdCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *iUserPostgresqlRepositoryUpdateUserByIdCall) Twice() *iUserPostgresqlRepositoryUpdateUserByIdCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *iUserPostgresqlRepositoryUpdateUserByIdCall) Times(i int) *iUserPostgresqlRepositoryUpdateUserByIdCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *iUserPostgresqlRepositoryUpdateUserByIdCall) WaitUntil(w <-chan time.Time) *iUserPostgresqlRepositoryUpdateUserByIdCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *iUserPostgresqlRepositoryUpdateUserByIdCall) After(d time.Duration) *iUserPostgresqlRepositoryUpdateUserByIdCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *iUserPostgresqlRepositoryUpdateUserByIdCall) Run(fn func(args mock.Arguments)) *iUserPostgresqlRepositoryUpdateUserByIdCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *iUserPostgresqlRepositoryUpdateUserByIdCall) Maybe() *iUserPostgresqlRepositoryUpdateUserByIdCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *iUserPostgresqlRepositoryUpdateUserByIdCall) TypedReturns(a *orm.User, b error) *iUserPostgresqlRepositoryUpdateUserByIdCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *iUserPostgresqlRepositoryUpdateUserByIdCall) ReturnsFn(fn func(string, *orm.User, *gorm.DB) (*orm.User, error)) *iUserPostgresqlRepositoryUpdateUserByIdCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *iUserPostgresqlRepositoryUpdateUserByIdCall) TypedRun(fn func(string, *orm.User, *gorm.DB)) *iUserPostgresqlRepositoryUpdateUserByIdCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_requestId := args.String(0)
		_req, _ := args.Get(1).(*orm.User)
		_tx, _ := args.Get(2).(*gorm.DB)
		fn(_requestId, _req, _tx)
	})
	return _c
}

func (_c *iUserPostgresqlRepositoryUpdateUserByIdCall) OnCreateUser(requestId string, req *orm.User, tx *gorm.DB) *iUserPostgresqlRepositoryCreateUserCall {
	return _c.Parent.OnCreateUser(requestId, req, tx)
}

func (_c *iUserPostgresqlRepositoryUpdateUserByIdCall) OnFindUserByEmail(requestId string, email string, tx *gorm.DB) *iUserPostgresqlRepositoryFindUserByEmailCall {
	return _c.Parent.OnFindUserByEmail(requestId, email, tx)
}

func (_c *iUserPostgresqlRepositoryUpdateUserByIdCall) OnFindUserById(requestId string, id string, tx *gorm.DB) *iUserPostgresqlRepositoryFindUserByIdCall {
	return _c.Parent.OnFindUserById(requestId, id, tx)
}

func (_c *iUserPostgresqlRepositoryUpdateUserByIdCall) OnUpdateUserById(requestId string, req *orm.User, tx *gorm.DB) *iUserPostgresqlRepositoryUpdateUserByIdCall {
	return _c.Parent.OnUpdateUserById(requestId, req, tx)
}

func (_c *iUserPostgresqlRepositoryUpdateUserByIdCall) OnCreateUserRaw(requestId interface{}, req interface{}, tx interface{}) *iUserPostgresqlRepositoryCreateUserCall {
	return _c.Parent.OnCreateUserRaw(requestId, req, tx)
}

func (_c *iUserPostgresqlRepositoryUpdateUserByIdCall) OnFindUserByEmailRaw(requestId interface{}, email interface{}, tx interface{}) *iUserPostgresqlRepositoryFindUserByEmailCall {
	return _c.Parent.OnFindUserByEmailRaw(requestId, email, tx)
}

func (_c *iUserPostgresqlRepositoryUpdateUserByIdCall) OnFindUserByIdRaw(requestId interface{}, id interface{}, tx interface{}) *iUserPostgresqlRepositoryFindUserByIdCall {
	return _c.Parent.OnFindUserByIdRaw(requestId, id, tx)
}

func (_c *iUserPostgresqlRepositoryUpdateUserByIdCall) OnUpdateUserByIdRaw(requestId interface{}, req interface{}, tx interface{}) *iUserPostgresqlRepositoryUpdateUserByIdCall {
	return _c.Parent.OnUpdateUserByIdRaw(requestId, req, tx)
}
