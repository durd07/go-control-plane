// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/network/ext_authz/v4alpha/ext_authz.proto

package envoy_extensions_filters_network_ext_authz_v4alpha

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

	v4alpha "github.com/durd07/go-control-plane/envoy/config/core/v4alpha"
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

	_ = v4alpha.ApiVersion(0)
)

// define the regex for a UUID once up-front
var _ext_authz_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ExtAuthz with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ExtAuthz) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetStatPrefix()) < 1 {
		return ExtAuthzValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 runes",
		}
	}

	if v, ok := interface{}(m.GetGrpcService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExtAuthzValidationError{
				field:  "GrpcService",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for FailureModeAllow

	// no validation rules for IncludePeerCertificate

	if _, ok := v4alpha.ApiVersion_name[int32(m.GetTransportApiVersion())]; !ok {
		return ExtAuthzValidationError{
			field:  "TransportApiVersion",
			reason: "value must be one of the defined enum values",
		}
	}

	if v, ok := interface{}(m.GetFilterEnabledMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExtAuthzValidationError{
				field:  "FilterEnabledMetadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ExtAuthzValidationError is the validation error returned by
// ExtAuthz.Validate if the designated constraints aren't met.
type ExtAuthzValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExtAuthzValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExtAuthzValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExtAuthzValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExtAuthzValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExtAuthzValidationError) ErrorName() string { return "ExtAuthzValidationError" }

// Error satisfies the builtin error interface
func (e ExtAuthzValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExtAuthz.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExtAuthzValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExtAuthzValidationError{}
