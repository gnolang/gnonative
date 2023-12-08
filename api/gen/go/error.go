package _go

import (
	"fmt"
	"io"

	"connectrpc.com/connect"
	"golang.org/x/xerrors"
)

// WithCode defines an error that can be used by helpers of this package.
type WithCode interface {
	error
	Code() ErrCode
	Grpc() error
}

// Codes returns a list of wrapped codes
func Codes(err error) []ErrCode {
	if err == nil {
		return nil
	}

	codes := []ErrCode{}

	if code := Code(err); code != -1 {
		codes = []ErrCode{code}
	}
	if cause := genericCause(err); cause != nil {
		causeCodes := Codes(cause)
		if len(causeCodes) > 0 {
			codes = append(codes, Codes(cause)...)
		}
	}

	return codes
}

// Has returns true if one of the error is or contains (wraps) an expected errcode
func Has(err error, code WithCode) bool {
	codeCode := code.Code()
	for _, otherCode := range Codes(err) {
		if otherCode == codeCode {
			return true
		}
	}
	return false
}

// Is returns true if the top-level error (not the FirstCode) is actually an ErrCode of the same value
func Is(err error, code WithCode) bool {
	return Code(err) == code.Code()
}

// Code returns the code of the actual error without trying to unwrap it, or -1.
func Code(err error) ErrCode {
	if err == nil {
		return -1
	}

	if typed, ok := err.(WithCode); ok {
		return typed.Code()
	}

	return -1
}

// LastCode walks the passed error and returns the code of the latest ErrCode, or -1.
func LastCode(err error) ErrCode {
	if err == nil {
		return -1
	}

	if cause := genericCause(err); cause != nil {
		if ret := LastCode(cause); ret != -1 {
			return ret
		}
	}

	return Code(err)
}

// FirstCode walks the passed error and returns the code of the first ErrCode met, or -1.
func FirstCode(err error) ErrCode {
	if err == nil {
		return -1
	}

	if code := Code(err); code != -1 {
		return code
	}

	if cause := genericCause(err); cause != nil {
		return FirstCode(cause)
	}

	return -1
}

func genericCause(err error) error {
	type causer interface{ Cause() error }
	type wrapper interface{ Unwrap() error }

	if causer, ok := err.(causer); ok {
		return causer.Cause()
	}

	if wrapper, ok := err.(wrapper); ok {
		return wrapper.Unwrap()
	}

	return nil
}

//
// Error
//

func (e ErrCode) Error() string {
	name, ok := ErrCode_name[int32(e)]
	if ok {
		return fmt.Sprintf("%s(#%d)", name, e)
	}
	return fmt.Sprintf("UNKNOWN_ERRCODE(#%d)", e)
}

func (e ErrCode) Code() ErrCode {
	return e
}

func (e ErrCode) Wrap(inner error) WithCode {
	return wrappedError{
		code:  e,
		inner: inner,
		frame: xerrors.Caller(1),
	}
}

// craft an connectRPC error with the rpc code error in the error details
func (e ErrCode) Grpc() error {
	err := connect.NewError(
		connect.CodeUnknown,
		e,
	)
	if detail, detailErr := connect.NewErrorDetail(&ErrDetails{
		Codes: Codes(e),
	}); detailErr == nil {
		err.AddDetail(detail)
	}
	return err
}

//
// ConfigurableError
//

type wrappedError struct {
	code  ErrCode
	inner error
	frame xerrors.Frame
}

func (e wrappedError) Error() string {
	return fmt.Sprintf("%s: %v", e.code, e.inner)
}

func (e wrappedError) Code() ErrCode {
	return e.code
}

// craft an connectRPC error with the rpc code error in the error details
func (e wrappedError) Grpc() error {
	err := connect.NewError(
		connect.CodeUnknown,
		e,
	)

	if detail, detailErr := connect.NewErrorDetail(&ErrDetails{
		Codes: Codes(e),
	}); detailErr == nil {
		err.AddDetail(detail)
	}
	return err
}

// Cause returns the inner error (github.com/pkg/errors)
func (e wrappedError) Cause() error {
	return e.inner
}

// Unwrap returns the inner error (go1.13)
func (e wrappedError) Unwrap() error {
	return e.inner
}

func (e wrappedError) Format(f fmt.State, c rune) {
	xerrors.FormatError(e, f, c)
	if f.Flag('+') {
		_, _ = io.WriteString(f, "\n")
		if sub := genericCause(e); sub != nil {
			if typed, ok := sub.(wrappedError); ok {
				sub = lightWrappedError{wrappedError: typed}
			}
			formatter, ok := sub.(fmt.Formatter)
			if ok {
				formatter.Format(f, c)
			}
		}
	}
}

func (e wrappedError) FormatError(p xerrors.Printer) error {
	p.Print(e.Error())
	if p.Detail() {
		e.frame.Format(p)
	}
	return nil
}

//
// light wrapper (used to make prettier (less verbose) stacks)
//

type lightWrappedError struct {
	wrappedError
	deepness int
}

func (e lightWrappedError) Error() string { return "" }

func (e lightWrappedError) Format(f fmt.State, c rune) {
	xerrors.FormatError(e, f, c)
	if f.Flag('+') {
		_, _ = io.WriteString(f, "\n")
		if sub := genericCause(e); sub != nil {
			if typed, ok := sub.(wrappedError); ok {
				sub = lightWrappedError{wrappedError: typed, deepness: e.deepness + 1}
			}
			formatter, ok := sub.(fmt.Formatter)
			if ok {
				formatter.Format(f, c)
			}
		}
	}
}

func (e lightWrappedError) FormatError(p xerrors.Printer) error {
	p.Printf("#%d", e.deepness+1)
	e.frame.Format(p)
	return nil
}
