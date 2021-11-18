// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: aws/kinesis/v1/kinesis.proto

package kinesisv1

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

// Validate checks the field values on GetStreamRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetStreamRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetStreamRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetStreamRequestMultiError, or nil if none found.
func (m *GetStreamRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetStreamRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetStreamName()) < 1 {
		err := GetStreamRequestValidationError{
			field:  "StreamName",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetRegion()) < 1 {
		err := GetStreamRequestValidationError{
			field:  "Region",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetAccount()) < 1 {
		err := GetStreamRequestValidationError{
			field:  "Account",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetStreamRequestMultiError(errors)
	}
	return nil
}

// GetStreamRequestMultiError is an error wrapping multiple validation errors
// returned by GetStreamRequest.ValidateAll() if the designated constraints
// aren't met.
type GetStreamRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetStreamRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetStreamRequestMultiError) AllErrors() []error { return m }

// GetStreamRequestValidationError is the validation error returned by
// GetStreamRequest.Validate if the designated constraints aren't met.
type GetStreamRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetStreamRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetStreamRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetStreamRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetStreamRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetStreamRequestValidationError) ErrorName() string { return "GetStreamRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetStreamRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetStreamRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetStreamRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetStreamRequestValidationError{}

// Validate checks the field values on GetStreamResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetStreamResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetStreamResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetStreamResponseMultiError, or nil if none found.
func (m *GetStreamResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetStreamResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStream()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetStreamResponseValidationError{
					field:  "Stream",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetStreamResponseValidationError{
					field:  "Stream",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStream()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetStreamResponseValidationError{
				field:  "Stream",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetStreamResponseMultiError(errors)
	}
	return nil
}

// GetStreamResponseMultiError is an error wrapping multiple validation errors
// returned by GetStreamResponse.ValidateAll() if the designated constraints
// aren't met.
type GetStreamResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetStreamResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetStreamResponseMultiError) AllErrors() []error { return m }

// GetStreamResponseValidationError is the validation error returned by
// GetStreamResponse.Validate if the designated constraints aren't met.
type GetStreamResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetStreamResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetStreamResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetStreamResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetStreamResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetStreamResponseValidationError) ErrorName() string {
	return "GetStreamResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetStreamResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetStreamResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetStreamResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetStreamResponseValidationError{}

// Validate checks the field values on UpdateShardCountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateShardCountRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateShardCountRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateShardCountRequestMultiError, or nil if none found.
func (m *UpdateShardCountRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateShardCountRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetStreamName()) < 1 {
		err := UpdateShardCountRequestValidationError{
			field:  "StreamName",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetRegion()) < 1 {
		err := UpdateShardCountRequestValidationError{
			field:  "Region",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for TargetShardCount

	if len(m.GetAccount()) < 1 {
		err := UpdateShardCountRequestValidationError{
			field:  "Account",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UpdateShardCountRequestMultiError(errors)
	}
	return nil
}

// UpdateShardCountRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateShardCountRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateShardCountRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateShardCountRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateShardCountRequestMultiError) AllErrors() []error { return m }

// UpdateShardCountRequestValidationError is the validation error returned by
// UpdateShardCountRequest.Validate if the designated constraints aren't met.
type UpdateShardCountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateShardCountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateShardCountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateShardCountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateShardCountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateShardCountRequestValidationError) ErrorName() string {
	return "UpdateShardCountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateShardCountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateShardCountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateShardCountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateShardCountRequestValidationError{}

// Validate checks the field values on UpdateShardCountResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateShardCountResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateShardCountResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateShardCountResponseMultiError, or nil if none found.
func (m *UpdateShardCountResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateShardCountResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateShardCountResponseMultiError(errors)
	}
	return nil
}

// UpdateShardCountResponseMultiError is an error wrapping multiple validation
// errors returned by UpdateShardCountResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateShardCountResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateShardCountResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateShardCountResponseMultiError) AllErrors() []error { return m }

// UpdateShardCountResponseValidationError is the validation error returned by
// UpdateShardCountResponse.Validate if the designated constraints aren't met.
type UpdateShardCountResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateShardCountResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateShardCountResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateShardCountResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateShardCountResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateShardCountResponseValidationError) ErrorName() string {
	return "UpdateShardCountResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateShardCountResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateShardCountResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateShardCountResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateShardCountResponseValidationError{}

// Validate checks the field values on Stream with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Stream) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Stream with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in StreamMultiError, or nil if none found.
func (m *Stream) ValidateAll() error {
	return m.validate(true)
}

func (m *Stream) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StreamName

	// no validation rules for Region

	// no validation rules for CurrentShardCount

	// no validation rules for Account

	if len(errors) > 0 {
		return StreamMultiError(errors)
	}
	return nil
}

// StreamMultiError is an error wrapping multiple validation errors returned by
// Stream.ValidateAll() if the designated constraints aren't met.
type StreamMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StreamMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StreamMultiError) AllErrors() []error { return m }

// StreamValidationError is the validation error returned by Stream.Validate if
// the designated constraints aren't met.
type StreamValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamValidationError) ErrorName() string { return "StreamValidationError" }

// Error satisfies the builtin error interface
func (e StreamValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStream.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamValidationError{}
