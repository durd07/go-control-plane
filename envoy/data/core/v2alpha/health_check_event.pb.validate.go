// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/data/core/v2alpha/health_check_event.proto

package v2alpha

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

// Validate checks the field values on HealthCheckEvent with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *HealthCheckEvent) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HealthCheckEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HealthCheckEventMultiError, or nil if none found.
func (m *HealthCheckEvent) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthCheckEvent) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := HealthCheckerType_name[int32(m.GetHealthCheckerType())]; !ok {
		err := HealthCheckEventValidationError{
			field:  "HealthCheckerType",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetHost()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HealthCheckEventValidationError{
					field:  "Host",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HealthCheckEventValidationError{
					field:  "Host",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHost()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheckEventValidationError{
				field:  "Host",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetClusterName()) < 1 {
		err := HealthCheckEventValidationError{
			field:  "ClusterName",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HealthCheckEventValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HealthCheckEventValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheckEventValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.Event.(type) {

	case *HealthCheckEvent_EjectUnhealthyEvent:

		if all {
			switch v := interface{}(m.GetEjectUnhealthyEvent()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HealthCheckEventValidationError{
						field:  "EjectUnhealthyEvent",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HealthCheckEventValidationError{
						field:  "EjectUnhealthyEvent",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetEjectUnhealthyEvent()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckEventValidationError{
					field:  "EjectUnhealthyEvent",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HealthCheckEvent_AddHealthyEvent:

		if all {
			switch v := interface{}(m.GetAddHealthyEvent()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HealthCheckEventValidationError{
						field:  "AddHealthyEvent",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HealthCheckEventValidationError{
						field:  "AddHealthyEvent",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAddHealthyEvent()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckEventValidationError{
					field:  "AddHealthyEvent",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HealthCheckEvent_HealthCheckFailureEvent:

		if all {
			switch v := interface{}(m.GetHealthCheckFailureEvent()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HealthCheckEventValidationError{
						field:  "HealthCheckFailureEvent",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HealthCheckEventValidationError{
						field:  "HealthCheckFailureEvent",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetHealthCheckFailureEvent()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckEventValidationError{
					field:  "HealthCheckFailureEvent",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HealthCheckEvent_DegradedHealthyHost:

		if all {
			switch v := interface{}(m.GetDegradedHealthyHost()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HealthCheckEventValidationError{
						field:  "DegradedHealthyHost",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HealthCheckEventValidationError{
						field:  "DegradedHealthyHost",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetDegradedHealthyHost()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckEventValidationError{
					field:  "DegradedHealthyHost",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HealthCheckEvent_NoLongerDegradedHost:

		if all {
			switch v := interface{}(m.GetNoLongerDegradedHost()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HealthCheckEventValidationError{
						field:  "NoLongerDegradedHost",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HealthCheckEventValidationError{
						field:  "NoLongerDegradedHost",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetNoLongerDegradedHost()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckEventValidationError{
					field:  "NoLongerDegradedHost",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		err := HealthCheckEventValidationError{
			field:  "Event",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return HealthCheckEventMultiError(errors)
	}
	return nil
}

// HealthCheckEventMultiError is an error wrapping multiple validation errors
// returned by HealthCheckEvent.ValidateAll() if the designated constraints
// aren't met.
type HealthCheckEventMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthCheckEventMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthCheckEventMultiError) AllErrors() []error { return m }

// HealthCheckEventValidationError is the validation error returned by
// HealthCheckEvent.Validate if the designated constraints aren't met.
type HealthCheckEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckEventValidationError) ErrorName() string { return "HealthCheckEventValidationError" }

// Error satisfies the builtin error interface
func (e HealthCheckEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheckEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckEventValidationError{}

// Validate checks the field values on HealthCheckEjectUnhealthy with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HealthCheckEjectUnhealthy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HealthCheckEjectUnhealthy with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HealthCheckEjectUnhealthyMultiError, or nil if none found.
func (m *HealthCheckEjectUnhealthy) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthCheckEjectUnhealthy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := HealthCheckFailureType_name[int32(m.GetFailureType())]; !ok {
		err := HealthCheckEjectUnhealthyValidationError{
			field:  "FailureType",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return HealthCheckEjectUnhealthyMultiError(errors)
	}
	return nil
}

// HealthCheckEjectUnhealthyMultiError is an error wrapping multiple validation
// errors returned by HealthCheckEjectUnhealthy.ValidateAll() if the
// designated constraints aren't met.
type HealthCheckEjectUnhealthyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthCheckEjectUnhealthyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthCheckEjectUnhealthyMultiError) AllErrors() []error { return m }

// HealthCheckEjectUnhealthyValidationError is the validation error returned by
// HealthCheckEjectUnhealthy.Validate if the designated constraints aren't met.
type HealthCheckEjectUnhealthyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckEjectUnhealthyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckEjectUnhealthyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckEjectUnhealthyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckEjectUnhealthyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckEjectUnhealthyValidationError) ErrorName() string {
	return "HealthCheckEjectUnhealthyValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheckEjectUnhealthyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheckEjectUnhealthy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckEjectUnhealthyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckEjectUnhealthyValidationError{}

// Validate checks the field values on HealthCheckAddHealthy with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HealthCheckAddHealthy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HealthCheckAddHealthy with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HealthCheckAddHealthyMultiError, or nil if none found.
func (m *HealthCheckAddHealthy) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthCheckAddHealthy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FirstCheck

	if len(errors) > 0 {
		return HealthCheckAddHealthyMultiError(errors)
	}
	return nil
}

// HealthCheckAddHealthyMultiError is an error wrapping multiple validation
// errors returned by HealthCheckAddHealthy.ValidateAll() if the designated
// constraints aren't met.
type HealthCheckAddHealthyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthCheckAddHealthyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthCheckAddHealthyMultiError) AllErrors() []error { return m }

// HealthCheckAddHealthyValidationError is the validation error returned by
// HealthCheckAddHealthy.Validate if the designated constraints aren't met.
type HealthCheckAddHealthyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckAddHealthyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckAddHealthyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckAddHealthyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckAddHealthyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckAddHealthyValidationError) ErrorName() string {
	return "HealthCheckAddHealthyValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheckAddHealthyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheckAddHealthy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckAddHealthyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckAddHealthyValidationError{}

// Validate checks the field values on HealthCheckFailure with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HealthCheckFailure) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HealthCheckFailure with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HealthCheckFailureMultiError, or nil if none found.
func (m *HealthCheckFailure) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthCheckFailure) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := HealthCheckFailureType_name[int32(m.GetFailureType())]; !ok {
		err := HealthCheckFailureValidationError{
			field:  "FailureType",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for FirstCheck

	if len(errors) > 0 {
		return HealthCheckFailureMultiError(errors)
	}
	return nil
}

// HealthCheckFailureMultiError is an error wrapping multiple validation errors
// returned by HealthCheckFailure.ValidateAll() if the designated constraints
// aren't met.
type HealthCheckFailureMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthCheckFailureMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthCheckFailureMultiError) AllErrors() []error { return m }

// HealthCheckFailureValidationError is the validation error returned by
// HealthCheckFailure.Validate if the designated constraints aren't met.
type HealthCheckFailureValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckFailureValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckFailureValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckFailureValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckFailureValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckFailureValidationError) ErrorName() string {
	return "HealthCheckFailureValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheckFailureValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheckFailure.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckFailureValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckFailureValidationError{}

// Validate checks the field values on DegradedHealthyHost with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DegradedHealthyHost) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DegradedHealthyHost with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DegradedHealthyHostMultiError, or nil if none found.
func (m *DegradedHealthyHost) ValidateAll() error {
	return m.validate(true)
}

func (m *DegradedHealthyHost) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DegradedHealthyHostMultiError(errors)
	}
	return nil
}

// DegradedHealthyHostMultiError is an error wrapping multiple validation
// errors returned by DegradedHealthyHost.ValidateAll() if the designated
// constraints aren't met.
type DegradedHealthyHostMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DegradedHealthyHostMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DegradedHealthyHostMultiError) AllErrors() []error { return m }

// DegradedHealthyHostValidationError is the validation error returned by
// DegradedHealthyHost.Validate if the designated constraints aren't met.
type DegradedHealthyHostValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DegradedHealthyHostValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DegradedHealthyHostValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DegradedHealthyHostValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DegradedHealthyHostValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DegradedHealthyHostValidationError) ErrorName() string {
	return "DegradedHealthyHostValidationError"
}

// Error satisfies the builtin error interface
func (e DegradedHealthyHostValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDegradedHealthyHost.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DegradedHealthyHostValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DegradedHealthyHostValidationError{}

// Validate checks the field values on NoLongerDegradedHost with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *NoLongerDegradedHost) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NoLongerDegradedHost with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// NoLongerDegradedHostMultiError, or nil if none found.
func (m *NoLongerDegradedHost) ValidateAll() error {
	return m.validate(true)
}

func (m *NoLongerDegradedHost) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return NoLongerDegradedHostMultiError(errors)
	}
	return nil
}

// NoLongerDegradedHostMultiError is an error wrapping multiple validation
// errors returned by NoLongerDegradedHost.ValidateAll() if the designated
// constraints aren't met.
type NoLongerDegradedHostMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NoLongerDegradedHostMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NoLongerDegradedHostMultiError) AllErrors() []error { return m }

// NoLongerDegradedHostValidationError is the validation error returned by
// NoLongerDegradedHost.Validate if the designated constraints aren't met.
type NoLongerDegradedHostValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NoLongerDegradedHostValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NoLongerDegradedHostValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NoLongerDegradedHostValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NoLongerDegradedHostValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NoLongerDegradedHostValidationError) ErrorName() string {
	return "NoLongerDegradedHostValidationError"
}

// Error satisfies the builtin error interface
func (e NoLongerDegradedHostValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNoLongerDegradedHost.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NoLongerDegradedHostValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NoLongerDegradedHostValidationError{}
