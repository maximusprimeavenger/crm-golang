// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api-gateway/grpc/item-grpc/item.proto

package proto

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

// Validate checks the field values on CreateItemRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateItemRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateItemRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateItemRequestMultiError, or nil if none found.
func (m *CreateItemRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateItemRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetItem()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateItemRequestValidationError{
					field:  "Item",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateItemRequestValidationError{
					field:  "Item",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetItem()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateItemRequestValidationError{
				field:  "Item",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateItemRequestMultiError(errors)
	}

	return nil
}

// CreateItemRequestMultiError is an error wrapping multiple validation errors
// returned by CreateItemRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateItemRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateItemRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateItemRequestMultiError) AllErrors() []error { return m }

// CreateItemRequestValidationError is the validation error returned by
// CreateItemRequest.Validate if the designated constraints aren't met.
type CreateItemRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateItemRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateItemRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateItemRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateItemRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateItemRequestValidationError) ErrorName() string {
	return "CreateItemRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateItemRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateItemRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateItemRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateItemRequestValidationError{}

// Validate checks the field values on CreateItemResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateItemResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateItemResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateItemResponseMultiError, or nil if none found.
func (m *CreateItemResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateItemResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetMessage()) < 5 {
		err := CreateItemResponseValidationError{
			field:  "Message",
			reason: "value length must be at least 5 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateItemResponseValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateItemResponseValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateItemResponseValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateItemResponseMultiError(errors)
	}

	return nil
}

// CreateItemResponseMultiError is an error wrapping multiple validation errors
// returned by CreateItemResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateItemResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateItemResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateItemResponseMultiError) AllErrors() []error { return m }

// CreateItemResponseValidationError is the validation error returned by
// CreateItemResponse.Validate if the designated constraints aren't met.
type CreateItemResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateItemResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateItemResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateItemResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateItemResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateItemResponseValidationError) ErrorName() string {
	return "CreateItemResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateItemResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateItemResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateItemResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateItemResponseValidationError{}

// Validate checks the field values on GetItemRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetItemRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetItemRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetItemRequestMultiError,
// or nil if none found.
func (m *GetItemRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetItemRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := GetItemRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetItemRequestMultiError(errors)
	}

	return nil
}

// GetItemRequestMultiError is an error wrapping multiple validation errors
// returned by GetItemRequest.ValidateAll() if the designated constraints
// aren't met.
type GetItemRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetItemRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetItemRequestMultiError) AllErrors() []error { return m }

// GetItemRequestValidationError is the validation error returned by
// GetItemRequest.Validate if the designated constraints aren't met.
type GetItemRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetItemRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetItemRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetItemRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetItemRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetItemRequestValidationError) ErrorName() string { return "GetItemRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetItemRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetItemRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetItemRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetItemRequestValidationError{}

// Validate checks the field values on GetItemResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetItemResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetItemResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetItemResponseMultiError, or nil if none found.
func (m *GetItemResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetItemResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetItem()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetItemResponseValidationError{
					field:  "Item",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetItemResponseValidationError{
					field:  "Item",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetItem()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetItemResponseValidationError{
				field:  "Item",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetItemResponseMultiError(errors)
	}

	return nil
}

// GetItemResponseMultiError is an error wrapping multiple validation errors
// returned by GetItemResponse.ValidateAll() if the designated constraints
// aren't met.
type GetItemResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetItemResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetItemResponseMultiError) AllErrors() []error { return m }

// GetItemResponseValidationError is the validation error returned by
// GetItemResponse.Validate if the designated constraints aren't met.
type GetItemResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetItemResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetItemResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetItemResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetItemResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetItemResponseValidationError) ErrorName() string { return "GetItemResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetItemResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetItemResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetItemResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetItemResponseValidationError{}

// Validate checks the field values on GetItemsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetItemsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetItemsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetItemsRequestMultiError, or nil if none found.
func (m *GetItemsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetItemsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetEmpty()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetItemsRequestValidationError{
					field:  "Empty",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetItemsRequestValidationError{
					field:  "Empty",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEmpty()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetItemsRequestValidationError{
				field:  "Empty",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetItemsRequestMultiError(errors)
	}

	return nil
}

// GetItemsRequestMultiError is an error wrapping multiple validation errors
// returned by GetItemsRequest.ValidateAll() if the designated constraints
// aren't met.
type GetItemsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetItemsRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetItemsRequestMultiError) AllErrors() []error { return m }

// GetItemsRequestValidationError is the validation error returned by
// GetItemsRequest.Validate if the designated constraints aren't met.
type GetItemsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetItemsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetItemsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetItemsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetItemsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetItemsRequestValidationError) ErrorName() string { return "GetItemsRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetItemsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetItemsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetItemsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetItemsRequestValidationError{}

// Validate checks the field values on GetItemsResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetItemsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetItemsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetItemsResponseMultiError, or nil if none found.
func (m *GetItemsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetItemsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetItemsResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetItemsResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetItemsResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetItemsResponseMultiError(errors)
	}

	return nil
}

// GetItemsResponseMultiError is an error wrapping multiple validation errors
// returned by GetItemsResponse.ValidateAll() if the designated constraints
// aren't met.
type GetItemsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetItemsResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetItemsResponseMultiError) AllErrors() []error { return m }

// GetItemsResponseValidationError is the validation error returned by
// GetItemsResponse.Validate if the designated constraints aren't met.
type GetItemsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetItemsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetItemsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetItemsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetItemsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetItemsResponseValidationError) ErrorName() string { return "GetItemsResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetItemsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetItemsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetItemsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetItemsResponseValidationError{}

// Validate checks the field values on PutItemRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PutItemRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PutItemRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PutItemRequestMultiError,
// or nil if none found.
func (m *PutItemRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PutItemRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := PutItemRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetItem()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PutItemRequestValidationError{
					field:  "Item",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PutItemRequestValidationError{
					field:  "Item",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetItem()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PutItemRequestValidationError{
				field:  "Item",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PutItemRequestMultiError(errors)
	}

	return nil
}

// PutItemRequestMultiError is an error wrapping multiple validation errors
// returned by PutItemRequest.ValidateAll() if the designated constraints
// aren't met.
type PutItemRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PutItemRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PutItemRequestMultiError) AllErrors() []error { return m }

// PutItemRequestValidationError is the validation error returned by
// PutItemRequest.Validate if the designated constraints aren't met.
type PutItemRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PutItemRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PutItemRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PutItemRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PutItemRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PutItemRequestValidationError) ErrorName() string { return "PutItemRequestValidationError" }

// Error satisfies the builtin error interface
func (e PutItemRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPutItemRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PutItemRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PutItemRequestValidationError{}

// Validate checks the field values on PutItemResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PutItemResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PutItemResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PutItemResponseMultiError, or nil if none found.
func (m *PutItemResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PutItemResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetItem()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PutItemResponseValidationError{
					field:  "Item",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PutItemResponseValidationError{
					field:  "Item",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetItem()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PutItemResponseValidationError{
				field:  "Item",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PutItemResponseValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PutItemResponseValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PutItemResponseValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PutItemResponseValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PutItemResponseValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PutItemResponseValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PutItemResponseMultiError(errors)
	}

	return nil
}

// PutItemResponseMultiError is an error wrapping multiple validation errors
// returned by PutItemResponse.ValidateAll() if the designated constraints
// aren't met.
type PutItemResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PutItemResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PutItemResponseMultiError) AllErrors() []error { return m }

// PutItemResponseValidationError is the validation error returned by
// PutItemResponse.Validate if the designated constraints aren't met.
type PutItemResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PutItemResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PutItemResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PutItemResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PutItemResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PutItemResponseValidationError) ErrorName() string { return "PutItemResponseValidationError" }

// Error satisfies the builtin error interface
func (e PutItemResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPutItemResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PutItemResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PutItemResponseValidationError{}

// Validate checks the field values on DeleteItemRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteItemRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteItemRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteItemRequestMultiError, or nil if none found.
func (m *DeleteItemRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteItemRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := DeleteItemRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteItemRequestMultiError(errors)
	}

	return nil
}

// DeleteItemRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteItemRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteItemRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteItemRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteItemRequestMultiError) AllErrors() []error { return m }

// DeleteItemRequestValidationError is the validation error returned by
// DeleteItemRequest.Validate if the designated constraints aren't met.
type DeleteItemRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteItemRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteItemRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteItemRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteItemRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteItemRequestValidationError) ErrorName() string {
	return "DeleteItemRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteItemRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteItemRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteItemRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteItemRequestValidationError{}

// Validate checks the field values on DeleteItemResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteItemResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteItemResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteItemResponseMultiError, or nil if none found.
func (m *DeleteItemResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteItemResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return DeleteItemResponseMultiError(errors)
	}

	return nil
}

// DeleteItemResponseMultiError is an error wrapping multiple validation errors
// returned by DeleteItemResponse.ValidateAll() if the designated constraints
// aren't met.
type DeleteItemResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteItemResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteItemResponseMultiError) AllErrors() []error { return m }

// DeleteItemResponseValidationError is the validation error returned by
// DeleteItemResponse.Validate if the designated constraints aren't met.
type DeleteItemResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteItemResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteItemResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteItemResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteItemResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteItemResponseValidationError) ErrorName() string {
	return "DeleteItemResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteItemResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteItemResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteItemResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteItemResponseValidationError{}
