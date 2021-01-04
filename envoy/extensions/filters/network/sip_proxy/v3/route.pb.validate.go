// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/network/sip_proxy/v3/route.proto

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

// define the regex for a UUID once up-front
var _route_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on RouteConfiguration with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RouteConfiguration) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	for idx, item := range m.GetRoutes() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RouteConfigurationValidationError{
					field:  fmt.Sprintf("Routes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// RouteConfigurationValidationError is the validation error returned by
// RouteConfiguration.Validate if the designated constraints aren't met.
type RouteConfigurationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouteConfigurationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouteConfigurationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouteConfigurationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouteConfigurationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouteConfigurationValidationError) ErrorName() string {
	return "RouteConfigurationValidationError"
}

// Error satisfies the builtin error interface
func (e RouteConfigurationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRouteConfiguration.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouteConfigurationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouteConfigurationValidationError{}

// Validate checks the field values on Route with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Route) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetMatch() == nil {
		return RouteValidationError{
			field:  "Match",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetMatch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RouteValidationError{
				field:  "Match",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetRoute() == nil {
		return RouteValidationError{
			field:  "Route",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetRoute()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RouteValidationError{
				field:  "Route",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// RouteValidationError is the validation error returned by Route.Validate if
// the designated constraints aren't met.
type RouteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouteValidationError) ErrorName() string { return "RouteValidationError" }

// Error satisfies the builtin error interface
func (e RouteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoute.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouteValidationError{}

// Validate checks the field values on RouteMatch with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *RouteMatch) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Invert

	for idx, item := range m.GetHeaders() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RouteMatchValidationError{
					field:  fmt.Sprintf("Headers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	switch m.MatchSpecifier.(type) {

	case *RouteMatch_Domain:
		// no validation rules for Domain

	default:
		return RouteMatchValidationError{
			field:  "MatchSpecifier",
			reason: "value is required",
		}

	}

	return nil
}

// RouteMatchValidationError is the validation error returned by
// RouteMatch.Validate if the designated constraints aren't met.
type RouteMatchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouteMatchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouteMatchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouteMatchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouteMatchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouteMatchValidationError) ErrorName() string { return "RouteMatchValidationError" }

// Error satisfies the builtin error interface
func (e RouteMatchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRouteMatch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouteMatchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouteMatchValidationError{}

// Validate checks the field values on RouteAction with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *RouteAction) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetMetadataMatch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RouteActionValidationError{
				field:  "MetadataMatch",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetRateLimits() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RouteActionValidationError{
					field:  fmt.Sprintf("RateLimits[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for StripServiceName

	switch m.ClusterSpecifier.(type) {

	case *RouteAction_Cluster:

		if utf8.RuneCountInString(m.GetCluster()) < 1 {
			return RouteActionValidationError{
				field:  "Cluster",
				reason: "value length must be at least 1 runes",
			}
		}

	case *RouteAction_ClusterHeader:

		if utf8.RuneCountInString(m.GetClusterHeader()) < 1 {
			return RouteActionValidationError{
				field:  "ClusterHeader",
				reason: "value length must be at least 1 runes",
			}
		}

		if !_RouteAction_ClusterHeader_Pattern.MatchString(m.GetClusterHeader()) {
			return RouteActionValidationError{
				field:  "ClusterHeader",
				reason: "value does not match regex pattern \"^[^\\x00\\n\\r]*$\"",
			}
		}

	default:
		return RouteActionValidationError{
			field:  "ClusterSpecifier",
			reason: "value is required",
		}

	}

	return nil
}

// RouteActionValidationError is the validation error returned by
// RouteAction.Validate if the designated constraints aren't met.
type RouteActionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouteActionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouteActionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouteActionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouteActionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouteActionValidationError) ErrorName() string { return "RouteActionValidationError" }

// Error satisfies the builtin error interface
func (e RouteActionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRouteAction.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouteActionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouteActionValidationError{}

var _RouteAction_ClusterHeader_Pattern = regexp.MustCompile("^[^\x00\n\r]*$")
