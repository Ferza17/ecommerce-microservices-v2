// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/commerce/model.proto

package commerce

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on CartItem with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CartItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CartItem with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CartItemMultiError, or nil
// if none found.
func (m *CartItem) ValidateAll() error {
	return m.validate(true)
}

func (m *CartItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for ProductId

	// no validation rules for UserId

	// no validation rules for Qty

	// no validation rules for Price

	if all {
		switch v := interface{}(m.GetCratedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CartItemValidationError{
					field:  "CratedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CartItemValidationError{
					field:  "CratedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCratedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CartItemValidationError{
				field:  "CratedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CartItemValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CartItemValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CartItemValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CartItemMultiError(errors)
	}

	return nil
}

// CartItemMultiError is an error wrapping multiple validation errors returned
// by CartItem.ValidateAll() if the designated constraints aren't met.
type CartItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CartItemMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CartItemMultiError) AllErrors() []error { return m }

// CartItemValidationError is the validation error returned by
// CartItem.Validate if the designated constraints aren't met.
type CartItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CartItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CartItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CartItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CartItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CartItemValidationError) ErrorName() string { return "CartItemValidationError" }

// Error satisfies the builtin error interface
func (e CartItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCartItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CartItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CartItemValidationError{}

// Validate checks the field values on WishlistItem with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *WishlistItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WishlistItem with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in WishlistItemMultiError, or
// nil if none found.
func (m *WishlistItem) ValidateAll() error {
	return m.validate(true)
}

func (m *WishlistItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for ProductId

	// no validation rules for UserId

	if len(errors) > 0 {
		return WishlistItemMultiError(errors)
	}

	return nil
}

// WishlistItemMultiError is an error wrapping multiple validation errors
// returned by WishlistItem.ValidateAll() if the designated constraints aren't met.
type WishlistItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WishlistItemMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WishlistItemMultiError) AllErrors() []error { return m }

// WishlistItemValidationError is the validation error returned by
// WishlistItem.Validate if the designated constraints aren't met.
type WishlistItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WishlistItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WishlistItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WishlistItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WishlistItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WishlistItemValidationError) ErrorName() string { return "WishlistItemValidationError" }

// Error satisfies the builtin error interface
func (e WishlistItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWishlistItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WishlistItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WishlistItemValidationError{}
