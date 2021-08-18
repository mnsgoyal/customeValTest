// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: github.com/maanasasubrahmanyam-sd/customeValTest/interfaces/test_server/test.proto

package test_server

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

// Validate checks the field values on SearchRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SearchRequest) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetQuery()) > 50 {
		return SearchRequestValidationError{
			field:  "Query",
			reason: "value length must be at most 50 bytes",
		}
	}

	if !_SearchRequest_Query_Pattern.MatchString(m.GetQuery()) {
		return SearchRequestValidationError{
			field:  "Query",
			reason: "value does not match regex pattern \"([A-Za-z]+) ([A-Za-z]+)*$\"",
		}
	}

	if err := m._validateEmail(m.GetEmailId()); err != nil {
		return SearchRequestValidationError{
			field:  "EmailId",
			reason: "value must be a valid email address",
			cause:  err,
		}
	}

	// no validation rules for Name

	// no validation rules for ResultPerPage

	return nil
}

func (m *SearchRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *SearchRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// SearchRequestValidationError is the validation error returned by
// SearchRequest.Validate if the designated constraints aren't met.
type SearchRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchRequestValidationError) ErrorName() string { return "SearchRequestValidationError" }

// Error satisfies the builtin error interface
func (e SearchRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchRequestValidationError{}

var _SearchRequest_Query_Pattern = regexp.MustCompile("([A-Za-z]+) ([A-Za-z]+)*$")

// Validate checks the field values on SearchResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SearchResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for SearchResponse

	// no validation rules for PageNumber

	// no validation rules for TotalPages

	return nil
}

// SearchResponseValidationError is the validation error returned by
// SearchResponse.Validate if the designated constraints aren't met.
type SearchResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchResponseValidationError) ErrorName() string { return "SearchResponseValidationError" }

// Error satisfies the builtin error interface
func (e SearchResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchResponseValidationError{}
