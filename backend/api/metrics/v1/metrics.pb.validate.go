// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: metrics/v1/metrics.proto

package metricsv1

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

// Validate checks the field values on Query with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Query) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Query with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in QueryMultiError, or nil if none found.
func (m *Query) ValidateAll() error {
	return m.validate(true)
}

func (m *Query) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetExpression()) < 1 {
		err := QueryValidationError{
			field:  "Expression",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetStartTimeMs() < 0 {
		err := QueryValidationError{
			field:  "StartTimeMs",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetEndTimeMs() < 0 {
		err := QueryValidationError{
			field:  "EndTimeMs",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for StepMs

	if len(errors) > 0 {
		return QueryMultiError(errors)
	}

	return nil
}

// QueryMultiError is an error wrapping multiple validation errors returned by
// Query.ValidateAll() if the designated constraints aren't met.
type QueryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueryMultiError) AllErrors() []error { return m }

// QueryValidationError is the validation error returned by Query.Validate if
// the designated constraints aren't met.
type QueryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryValidationError) ErrorName() string { return "QueryValidationError" }

// Error satisfies the builtin error interface
func (e QueryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuery.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryValidationError{}

// Validate checks the field values on GetMetricsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetMetricsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetMetricsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetMetricsRequestMultiError, or nil if none found.
func (m *GetMetricsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetMetricsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetMetricQueries()) < 1 {
		err := GetMetricsRequestValidationError{
			field:  "MetricQueries",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetMetricQueries() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetMetricsRequestValidationError{
						field:  fmt.Sprintf("MetricQueries[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetMetricsRequestValidationError{
						field:  fmt.Sprintf("MetricQueries[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetMetricsRequestValidationError{
					field:  fmt.Sprintf("MetricQueries[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetMetricsRequestMultiError(errors)
	}

	return nil
}

// GetMetricsRequestMultiError is an error wrapping multiple validation errors
// returned by GetMetricsRequest.ValidateAll() if the designated constraints
// aren't met.
type GetMetricsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetMetricsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetMetricsRequestMultiError) AllErrors() []error { return m }

// GetMetricsRequestValidationError is the validation error returned by
// GetMetricsRequest.Validate if the designated constraints aren't met.
type GetMetricsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMetricsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMetricsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMetricsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMetricsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMetricsRequestValidationError) ErrorName() string {
	return "GetMetricsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetMetricsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMetricsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMetricsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMetricsRequestValidationError{}

// Validate checks the field values on DataPoint with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DataPoint) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DataPoint with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DataPointMultiError, or nil
// if none found.
func (m *DataPoint) ValidateAll() error {
	return m.validate(true)
}

func (m *DataPoint) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Value

	// no validation rules for Timestamp

	if len(errors) > 0 {
		return DataPointMultiError(errors)
	}

	return nil
}

// DataPointMultiError is an error wrapping multiple validation errors returned
// by DataPoint.ValidateAll() if the designated constraints aren't met.
type DataPointMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DataPointMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DataPointMultiError) AllErrors() []error { return m }

// DataPointValidationError is the validation error returned by
// DataPoint.Validate if the designated constraints aren't met.
type DataPointValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DataPointValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DataPointValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DataPointValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DataPointValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DataPointValidationError) ErrorName() string { return "DataPointValidationError" }

// Error satisfies the builtin error interface
func (e DataPointValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDataPoint.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DataPointValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DataPointValidationError{}

// Validate checks the field values on Metrics with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Metrics) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Metrics with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MetricsMultiError, or nil if none found.
func (m *Metrics) ValidateAll() error {
	return m.validate(true)
}

func (m *Metrics) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetDataPoints() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MetricsValidationError{
						field:  fmt.Sprintf("DataPoints[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MetricsValidationError{
						field:  fmt.Sprintf("DataPoints[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MetricsValidationError{
					field:  fmt.Sprintf("DataPoints[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Label

	// no validation rules for Tags

	if len(errors) > 0 {
		return MetricsMultiError(errors)
	}

	return nil
}

// MetricsMultiError is an error wrapping multiple validation errors returned
// by Metrics.ValidateAll() if the designated constraints aren't met.
type MetricsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetricsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetricsMultiError) AllErrors() []error { return m }

// MetricsValidationError is the validation error returned by Metrics.Validate
// if the designated constraints aren't met.
type MetricsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetricsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetricsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetricsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetricsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetricsValidationError) ErrorName() string { return "MetricsValidationError" }

// Error satisfies the builtin error interface
func (e MetricsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetrics.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetricsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetricsValidationError{}

// Validate checks the field values on Result with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Result) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Result with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ResultMultiError, or nil if none found.
func (m *Result) ValidateAll() error {
	return m.validate(true)
}

func (m *Result) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetMetrics() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ResultValidationError{
						field:  fmt.Sprintf("Metrics[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ResultValidationError{
						field:  fmt.Sprintf("Metrics[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResultValidationError{
					field:  fmt.Sprintf("Metrics[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ResultMultiError(errors)
	}

	return nil
}

// ResultMultiError is an error wrapping multiple validation errors returned by
// Result.ValidateAll() if the designated constraints aren't met.
type ResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResultMultiError) AllErrors() []error { return m }

// ResultValidationError is the validation error returned by Result.Validate if
// the designated constraints aren't met.
type ResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResultValidationError) ErrorName() string { return "ResultValidationError" }

// Error satisfies the builtin error interface
func (e ResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResultValidationError{}

// Validate checks the field values on GetMetricsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetMetricsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetMetricsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetMetricsResponseMultiError, or nil if none found.
func (m *GetMetricsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetMetricsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	{
		sorted_keys := make([]string, len(m.GetQueryResults()))
		i := 0
		for key := range m.GetQueryResults() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetQueryResults()[key]
			_ = val

			// no validation rules for QueryResults[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, GetMetricsResponseValidationError{
							field:  fmt.Sprintf("QueryResults[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, GetMetricsResponseValidationError{
							field:  fmt.Sprintf("QueryResults[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return GetMetricsResponseValidationError{
						field:  fmt.Sprintf("QueryResults[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return GetMetricsResponseMultiError(errors)
	}

	return nil
}

// GetMetricsResponseMultiError is an error wrapping multiple validation errors
// returned by GetMetricsResponse.ValidateAll() if the designated constraints
// aren't met.
type GetMetricsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetMetricsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetMetricsResponseMultiError) AllErrors() []error { return m }

// GetMetricsResponseValidationError is the validation error returned by
// GetMetricsResponse.Validate if the designated constraints aren't met.
type GetMetricsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMetricsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMetricsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMetricsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMetricsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMetricsResponseValidationError) ErrorName() string {
	return "GetMetricsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetMetricsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMetricsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMetricsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMetricsResponseValidationError{}
