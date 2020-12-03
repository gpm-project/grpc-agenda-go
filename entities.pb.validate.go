// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: agenda/entities.proto

package grpc_agenda_go

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
var _entities_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Person with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Person) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Email

	return nil
}

// PersonValidationError is the validation error returned by Person.Validate if
// the designated constraints aren't met.
type PersonValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PersonValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PersonValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PersonValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PersonValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PersonValidationError) ErrorName() string { return "PersonValidationError" }

// Error satisfies the builtin error interface
func (e PersonValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPerson.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PersonValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PersonValidationError{}

// Validate checks the field values on PersonList with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *PersonList) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetPersons() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PersonListValidationError{
					field:  fmt.Sprintf("Persons[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PersonListValidationError is the validation error returned by
// PersonList.Validate if the designated constraints aren't met.
type PersonListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PersonListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PersonListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PersonListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PersonListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PersonListValidationError) ErrorName() string { return "PersonListValidationError" }

// Error satisfies the builtin error interface
func (e PersonListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPersonList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PersonListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PersonListValidationError{}

// Validate checks the field values on AddPersonRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *AddPersonRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	if err := m._validateEmail(m.GetEmail()); err != nil {
		return AddPersonRequestValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
	}

	return nil
}

func (m *AddPersonRequest) _validateHostname(host string) error {
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

func (m *AddPersonRequest) _validateEmail(addr string) error {
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

// AddPersonRequestValidationError is the validation error returned by
// AddPersonRequest.Validate if the designated constraints aren't met.
type AddPersonRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddPersonRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddPersonRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddPersonRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddPersonRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddPersonRequestValidationError) ErrorName() string { return "AddPersonRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddPersonRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddPersonRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddPersonRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddPersonRequestValidationError{}

// Validate checks the field values on ListPersonRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListPersonRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListPersonRequestValidationError is the validation error returned by
// ListPersonRequest.Validate if the designated constraints aren't met.
type ListPersonRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPersonRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPersonRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPersonRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPersonRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPersonRequestValidationError) ErrorName() string {
	return "ListPersonRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListPersonRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPersonRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPersonRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPersonRequestValidationError{}