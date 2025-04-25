package utils

import "errors"

var (
	ErrNotFound            = errors.New("not found")
	ErrBadRequest          = errors.New("bad Request")
	ErrQueryRead           = errors.New("error while querying storage")
	ErrQueryTxBegin        = errors.New("error while preparing repository transaction")
	ErrQueryTxInsert       = errors.New("error while inserting record")
	ErrQueryTxRollback     = errors.New("error while cancelling record")
	ErrQueryTxUpdate       = errors.New("error while updating record")
	ErrQueryTxCommit       = errors.New("error while committing record changes")
	ErrInternalServerError = errors.New("internal Server Error")
)

type Causer interface {
	Cause() error
}

type ErrorWithCode struct {
	code  int
	cause error
}

func (err *ErrorWithCode) Error() string {
	return err.Cause().Error()
}

func (err *ErrorWithCode) Code() int {
	return err.code
}

func (err *ErrorWithCode) Cause() error {
	return err.cause
}
