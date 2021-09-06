// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/resource_monitor/fixed_heap/v2alpha/fixed_heap.proto

package envoy_config_resource_monitor_fixed_heap_v2alpha

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

// Validate checks the field values on FixedHeapConfig with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *FixedHeapConfig) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetMaxHeapSizeBytes() <= 0 {
		return FixedHeapConfigValidationError{
			field:  "MaxHeapSizeBytes",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// FixedHeapConfigValidationError is the validation error returned by
// FixedHeapConfig.Validate if the designated constraints aren't met.
type FixedHeapConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FixedHeapConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FixedHeapConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FixedHeapConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FixedHeapConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FixedHeapConfigValidationError) ErrorName() string { return "FixedHeapConfigValidationError" }

// Error satisfies the builtin error interface
func (e FixedHeapConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFixedHeapConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FixedHeapConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FixedHeapConfigValidationError{}
