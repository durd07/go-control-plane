// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/http/health_check/v2/health_check.proto

package health_checkv2

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

// Validate checks the field values on HealthCheck with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HealthCheck) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HealthCheck with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in HealthCheckMultiError, or
// nil if none found.
func (m *HealthCheck) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthCheck) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPassThroughMode() == nil {
		err := HealthCheckValidationError{
			field:  "PassThroughMode",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetPassThroughMode()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HealthCheckValidationError{
					field:  "PassThroughMode",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HealthCheckValidationError{
					field:  "PassThroughMode",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPassThroughMode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheckValidationError{
				field:  "PassThroughMode",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCacheTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HealthCheckValidationError{
					field:  "CacheTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HealthCheckValidationError{
					field:  "CacheTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCacheTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheckValidationError{
				field:  "CacheTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	{
		sorted_keys := make([]string, len(m.GetClusterMinHealthyPercentages()))
		i := 0
		for key := range m.GetClusterMinHealthyPercentages() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetClusterMinHealthyPercentages()[key]
			_ = val

			// no validation rules for ClusterMinHealthyPercentages[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, HealthCheckValidationError{
							field:  fmt.Sprintf("ClusterMinHealthyPercentages[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, HealthCheckValidationError{
							field:  fmt.Sprintf("ClusterMinHealthyPercentages[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return HealthCheckValidationError{
						field:  fmt.Sprintf("ClusterMinHealthyPercentages[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	for idx, item := range m.GetHeaders() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HealthCheckValidationError{
						field:  fmt.Sprintf("Headers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HealthCheckValidationError{
						field:  fmt.Sprintf("Headers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckValidationError{
					field:  fmt.Sprintf("Headers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return HealthCheckMultiError(errors)
	}
	return nil
}

// HealthCheckMultiError is an error wrapping multiple validation errors
// returned by HealthCheck.ValidateAll() if the designated constraints aren't met.
type HealthCheckMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthCheckMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthCheckMultiError) AllErrors() []error { return m }

// HealthCheckValidationError is the validation error returned by
// HealthCheck.Validate if the designated constraints aren't met.
type HealthCheckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckValidationError) ErrorName() string { return "HealthCheckValidationError" }

// Error satisfies the builtin error interface
func (e HealthCheckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckValidationError{}
