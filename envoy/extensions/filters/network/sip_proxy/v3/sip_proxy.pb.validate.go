// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/network/sip_proxy/v3/sip_proxy.proto

package envoy_extensions_filters_network_sip_proxy_v3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on SipProxy with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *SipProxy) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetStatPrefix()) < 1 {
		return SipProxyValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 runes",
		}
	}

	if v, ok := interface{}(m.GetRouteConfig()).(interface{ Validate() error }); ok {
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

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SipProxyValidationError{
					field:  fmt.Sprintf("SipFilters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetSettings()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SipProxyValidationError{
				field:  "Settings",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

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
// proto definition for this message. If any rules are violated, an error is returned.
func (m *SipFilter) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return SipFilterValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	switch m.ConfigType.(type) {

	case *SipFilter_TypedConfig:

		if v, ok := interface{}(m.GetTypedConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SipFilterValidationError{
					field:  "TypedConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

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
// violated, an error is returned.
func (m *SipProtocolOptions) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for SessionAffinity

	// no validation rules for RegistrationAffinity

	return nil
}

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

// Validate checks the field values on SipProxy_SipSettings with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SipProxy_SipSettings) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTransactionTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SipProxy_SipSettingsValidationError{
				field:  "TransactionTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

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
