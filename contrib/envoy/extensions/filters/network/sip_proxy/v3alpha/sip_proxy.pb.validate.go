// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: contrib/envoy/extensions/filters/network/sip_proxy/v3alpha/sip_proxy.proto

package v3alpha

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

// Validate checks the field values on SipProxy with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SipProxy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SipProxy with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SipProxyMultiError, or nil
// if none found.
func (m *SipProxy) ValidateAll() error {
	return m.validate(true)
}

func (m *SipProxy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetStatPrefix()) < 1 {
		err := SipProxyValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetRouteConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SipProxyValidationError{
					field:  "RouteConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SipProxyValidationError{
					field:  "RouteConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRouteConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SipProxyValidationError{
				field:  "RouteConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetSipFilters() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SipProxyValidationError{
						field:  fmt.Sprintf("SipFilters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SipProxyValidationError{
						field:  fmt.Sprintf("SipFilters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SipProxyValidationError{
					field:  fmt.Sprintf("SipFilters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetSettings()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SipProxyValidationError{
					field:  "Settings",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SipProxyValidationError{
					field:  "Settings",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSettings()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SipProxyValidationError{
				field:  "Settings",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SipProxyMultiError(errors)
	}
	return nil
}

// SipProxyMultiError is an error wrapping multiple validation errors returned
// by SipProxy.ValidateAll() if the designated constraints aren't met.
type SipProxyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SipProxyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SipProxyMultiError) AllErrors() []error { return m }

// SipProxyValidationError is the validation error returned by
// SipProxy.Validate if the designated constraints aren't met.
type SipProxyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SipProxyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SipProxyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SipProxyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SipProxyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SipProxyValidationError) ErrorName() string { return "SipProxyValidationError" }

// Error satisfies the builtin error interface
func (e SipProxyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSipProxy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SipProxyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SipProxyValidationError{}

// Validate checks the field values on SipFilter with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SipFilter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SipFilter with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SipFilterMultiError, or nil
// if none found.
func (m *SipFilter) ValidateAll() error {
	return m.validate(true)
}

func (m *SipFilter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := SipFilterValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	switch m.ConfigType.(type) {

	case *SipFilter_TypedConfig:

		if all {
			switch v := interface{}(m.GetTypedConfig()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SipFilterValidationError{
						field:  "TypedConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SipFilterValidationError{
						field:  "TypedConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetTypedConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SipFilterValidationError{
					field:  "TypedConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SipFilterMultiError(errors)
	}
	return nil
}

// SipFilterMultiError is an error wrapping multiple validation errors returned
// by SipFilter.ValidateAll() if the designated constraints aren't met.
type SipFilterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SipFilterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SipFilterMultiError) AllErrors() []error { return m }

// SipFilterValidationError is the validation error returned by
// SipFilter.Validate if the designated constraints aren't met.
type SipFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SipFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SipFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SipFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SipFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SipFilterValidationError) ErrorName() string { return "SipFilterValidationError" }

// Error satisfies the builtin error interface
func (e SipFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSipFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SipFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SipFilterValidationError{}

// Validate checks the field values on SipProtocolOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SipProtocolOptions) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SipProtocolOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SipProtocolOptionsMultiError, or nil if none found.
func (m *SipProtocolOptions) ValidateAll() error {
	return m.validate(true)
}

func (m *SipProtocolOptions) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SessionAffinity

	// no validation rules for RegistrationAffinity

	if all {
		switch v := interface{}(m.GetCustomizedAffinity()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SipProtocolOptionsValidationError{
					field:  "CustomizedAffinity",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SipProtocolOptionsValidationError{
					field:  "CustomizedAffinity",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCustomizedAffinity()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SipProtocolOptionsValidationError{
				field:  "CustomizedAffinity",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SipProtocolOptionsMultiError(errors)
	}
	return nil
}

// SipProtocolOptionsMultiError is an error wrapping multiple validation errors
// returned by SipProtocolOptions.ValidateAll() if the designated constraints
// aren't met.
type SipProtocolOptionsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SipProtocolOptionsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SipProtocolOptionsMultiError) AllErrors() []error { return m }

// SipProtocolOptionsValidationError is the validation error returned by
// SipProtocolOptions.Validate if the designated constraints aren't met.
type SipProtocolOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SipProtocolOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SipProtocolOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SipProtocolOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SipProtocolOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SipProtocolOptionsValidationError) ErrorName() string {
	return "SipProtocolOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e SipProtocolOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSipProtocolOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SipProtocolOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SipProtocolOptionsValidationError{}

// Validate checks the field values on CustomizedAffinity with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CustomizedAffinity) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CustomizedAffinity with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CustomizedAffinityMultiError, or nil if none found.
func (m *CustomizedAffinity) ValidateAll() error {
	return m.validate(true)
}

func (m *CustomizedAffinity) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetEntries() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CustomizedAffinityValidationError{
						field:  fmt.Sprintf("Entries[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CustomizedAffinityValidationError{
						field:  fmt.Sprintf("Entries[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CustomizedAffinityValidationError{
					field:  fmt.Sprintf("Entries[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for StopLoadBalance

	if len(errors) > 0 {
		return CustomizedAffinityMultiError(errors)
	}
	return nil
}

// CustomizedAffinityMultiError is an error wrapping multiple validation errors
// returned by CustomizedAffinity.ValidateAll() if the designated constraints
// aren't met.
type CustomizedAffinityMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CustomizedAffinityMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CustomizedAffinityMultiError) AllErrors() []error { return m }

// CustomizedAffinityValidationError is the validation error returned by
// CustomizedAffinity.Validate if the designated constraints aren't met.
type CustomizedAffinityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CustomizedAffinityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CustomizedAffinityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CustomizedAffinityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CustomizedAffinityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CustomizedAffinityValidationError) ErrorName() string {
	return "CustomizedAffinityValidationError"
}

// Error satisfies the builtin error interface
func (e CustomizedAffinityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCustomizedAffinity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CustomizedAffinityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CustomizedAffinityValidationError{}

// Validate checks the field values on CustomizedAffinityEntry with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CustomizedAffinityEntry) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CustomizedAffinityEntry with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CustomizedAffinityEntryMultiError, or nil if none found.
func (m *CustomizedAffinityEntry) ValidateAll() error {
	return m.validate(true)
}

func (m *CustomizedAffinityEntry) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for KeyName

	// no validation rules for Subscribe

	// no validation rules for Query

	if all {
		switch v := interface{}(m.GetCache()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CustomizedAffinityEntryValidationError{
					field:  "Cache",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CustomizedAffinityEntryValidationError{
					field:  "Cache",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCache()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CustomizedAffinityEntryValidationError{
				field:  "Cache",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CustomizedAffinityEntryMultiError(errors)
	}
	return nil
}

// CustomizedAffinityEntryMultiError is an error wrapping multiple validation
// errors returned by CustomizedAffinityEntry.ValidateAll() if the designated
// constraints aren't met.
type CustomizedAffinityEntryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CustomizedAffinityEntryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CustomizedAffinityEntryMultiError) AllErrors() []error { return m }

// CustomizedAffinityEntryValidationError is the validation error returned by
// CustomizedAffinityEntry.Validate if the designated constraints aren't met.
type CustomizedAffinityEntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CustomizedAffinityEntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CustomizedAffinityEntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CustomizedAffinityEntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CustomizedAffinityEntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CustomizedAffinityEntryValidationError) ErrorName() string {
	return "CustomizedAffinityEntryValidationError"
}

// Error satisfies the builtin error interface
func (e CustomizedAffinityEntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCustomizedAffinityEntry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CustomizedAffinityEntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CustomizedAffinityEntryValidationError{}

// Validate checks the field values on Cache with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Cache) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Cache with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CacheMultiError, or nil if none found.
func (m *Cache) ValidateAll() error {
	return m.validate(true)
}

func (m *Cache) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for MaxCacheItem

	// no validation rules for AddQueryToCache

	if len(errors) > 0 {
		return CacheMultiError(errors)
	}
	return nil
}

// CacheMultiError is an error wrapping multiple validation errors returned by
// Cache.ValidateAll() if the designated constraints aren't met.
type CacheMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CacheMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CacheMultiError) AllErrors() []error { return m }

// CacheValidationError is the validation error returned by Cache.Validate if
// the designated constraints aren't met.
type CacheValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CacheValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CacheValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CacheValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CacheValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CacheValidationError) ErrorName() string { return "CacheValidationError" }

// Error satisfies the builtin error interface
func (e CacheValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCache.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CacheValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CacheValidationError{}

// Validate checks the field values on LocalService with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LocalService) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LocalService with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LocalServiceMultiError, or
// nil if none found.
func (m *LocalService) ValidateAll() error {
	return m.validate(true)
}

func (m *LocalService) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Domain

	// no validation rules for Parameter

	if len(errors) > 0 {
		return LocalServiceMultiError(errors)
	}
	return nil
}

// LocalServiceMultiError is an error wrapping multiple validation errors
// returned by LocalService.ValidateAll() if the designated constraints aren't met.
type LocalServiceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LocalServiceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LocalServiceMultiError) AllErrors() []error { return m }

// LocalServiceValidationError is the validation error returned by
// LocalService.Validate if the designated constraints aren't met.
type LocalServiceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocalServiceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocalServiceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocalServiceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocalServiceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocalServiceValidationError) ErrorName() string { return "LocalServiceValidationError" }

// Error satisfies the builtin error interface
func (e LocalServiceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocalService.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocalServiceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocalServiceValidationError{}

// Validate checks the field values on SipProxy_SipSettings with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SipProxy_SipSettings) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SipProxy_SipSettings with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SipProxy_SipSettingsMultiError, or nil if none found.
func (m *SipProxy_SipSettings) ValidateAll() error {
	return m.validate(true)
}

func (m *SipProxy_SipSettings) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTransactionTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SipProxy_SipSettingsValidationError{
					field:  "TransactionTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SipProxy_SipSettingsValidationError{
					field:  "TransactionTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTransactionTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SipProxy_SipSettingsValidationError{
				field:  "TransactionTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetLocalServices() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SipProxy_SipSettingsValidationError{
						field:  fmt.Sprintf("LocalServices[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SipProxy_SipSettingsValidationError{
						field:  fmt.Sprintf("LocalServices[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SipProxy_SipSettingsValidationError{
					field:  fmt.Sprintf("LocalServices[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetTraServiceConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SipProxy_SipSettingsValidationError{
					field:  "TraServiceConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SipProxy_SipSettingsValidationError{
					field:  "TraServiceConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTraServiceConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SipProxy_SipSettingsValidationError{
				field:  "TraServiceConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for OperateVia

	if len(errors) > 0 {
		return SipProxy_SipSettingsMultiError(errors)
	}
	return nil
}

// SipProxy_SipSettingsMultiError is an error wrapping multiple validation
// errors returned by SipProxy_SipSettings.ValidateAll() if the designated
// constraints aren't met.
type SipProxy_SipSettingsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SipProxy_SipSettingsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SipProxy_SipSettingsMultiError) AllErrors() []error { return m }

// SipProxy_SipSettingsValidationError is the validation error returned by
// SipProxy_SipSettings.Validate if the designated constraints aren't met.
type SipProxy_SipSettingsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SipProxy_SipSettingsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SipProxy_SipSettingsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SipProxy_SipSettingsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SipProxy_SipSettingsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SipProxy_SipSettingsValidationError) ErrorName() string {
	return "SipProxy_SipSettingsValidationError"
}

// Error satisfies the builtin error interface
func (e SipProxy_SipSettingsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSipProxy_SipSettings.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SipProxy_SipSettingsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SipProxy_SipSettingsValidationError{}
