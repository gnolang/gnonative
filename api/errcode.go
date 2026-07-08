package api

import (
	"fmt"
	"io"

	"golang.org/x/xerrors"
)

// ErrCode defines errors for the gnonative API functions. These are converted
// from the Go error types returned by gnoclient. The names and values match
// api/rpc.proto (kept generated-style, e.g. ErrCode_ErrSetRemote).
type ErrCode int32

const (
	// Undefined is the default value. It should never be set manually
	ErrCode_Undefined ErrCode = 0
	// TODO indicates that you plan to create an error later
	ErrCode_TODO ErrCode = 1
	// ErrNotImplemented indicates that a method is not implemented yet
	ErrCode_ErrNotImplemented ErrCode = 2
	// ErrInternal indicates an unknown error (without Code), i.e. in gRPC
	ErrCode_ErrInternal             ErrCode = 3
	ErrCode_ErrInvalidInput         ErrCode = 100
	ErrCode_ErrBridgeInterrupted    ErrCode = 101
	ErrCode_ErrMissingInput         ErrCode = 102
	ErrCode_ErrSerialization        ErrCode = 103
	ErrCode_ErrDeserialization      ErrCode = 104
	ErrCode_ErrInitService          ErrCode = 105
	ErrCode_ErrSetRemote            ErrCode = 106
	ErrCode_ErrKeyNameExists        ErrCode = 107
	ErrCode_ErrCryptoKeyTypeUnknown ErrCode = 150
	// ErrCryptoKeyNotFound indicates that the doesn't exist in the keybase
	ErrCode_ErrCryptoKeyNotFound ErrCode = 151
	// ErrNoActiveAccount indicates that no account with the given address has been activated with ActivateAccount
	ErrCode_ErrNoActiveAccount ErrCode = 152
	ErrCode_ErrRunGRPCServer   ErrCode = 153
	// ErrDecryptionFailed indicates a decryption failure including a wrong password
	ErrCode_ErrDecryptionFailed ErrCode = 154
	ErrCode_ErrTxDecode         ErrCode = 200
	ErrCode_ErrInvalidSequence  ErrCode = 201
	ErrCode_ErrUnauthorized     ErrCode = 202
	// ErrInsufficientFunds indicates that there are insufficient funds to pay for fees
	ErrCode_ErrInsufficientFunds ErrCode = 203
	// ErrUnknownRequest indicates that the path of a realm function call is unrecognized
	ErrCode_ErrUnknownRequest ErrCode = 204
	// ErrInvalidAddress indicates that an account address is blank or the bech32 can't be decoded
	ErrCode_ErrInvalidAddress ErrCode = 205
	// ErrUnknownAddress indicates that the address is unknown on the blockchain
	ErrCode_ErrUnknownAddress ErrCode = 206
	// ErrInvalidPubKey indicates that the public key was not found or has an invalid algorithm or format
	ErrCode_ErrInvalidPubKey ErrCode = 207
	// ErrInsufficientCoins indicates that the transaction has insufficient account funds to send
	ErrCode_ErrInsufficientCoins ErrCode = 208
	// ErrInvalidCoins indicates that the transaction Coins are not sorted, or don't have a
	// positive amount, or the coin Denom contains upper case characters
	ErrCode_ErrInvalidCoins ErrCode = 209
	// ErrInvalidGasWanted indicates that the transaction gas wanted is too large or otherwise invalid
	ErrCode_ErrInvalidGasWanted ErrCode = 210
	// ErrOutOfGas indicates that the transaction doesn't have enough gas
	ErrCode_ErrOutOfGas ErrCode = 211
	// ErrMemoTooLarge indicates that the transaction memo is too large
	ErrCode_ErrMemoTooLarge ErrCode = 212
	// ErrInsufficientFee indicates that the gas fee is insufficient
	ErrCode_ErrInsufficientFee ErrCode = 213
	// ErrTooManySignatures indicates that the transaction has too many signatures
	ErrCode_ErrTooManySignatures ErrCode = 214
	// ErrNoSignatures indicates that the transaction has no signatures
	ErrCode_ErrNoSignatures ErrCode = 215
	// ErrGasOverflow indicates that an action results in a gas consumption unsigned integer overflow
	ErrCode_ErrGasOverflow ErrCode = 216
	// ErrInvalidPkgPath indicates that the package path is not recognized.
	ErrCode_ErrInvalidPkgPath ErrCode = 217
	ErrCode_ErrInvalidStmt    ErrCode = 218
	ErrCode_ErrInvalidExpr    ErrCode = 219
)

// Enum value maps for ErrCode.
var (
	ErrCode_name = map[int32]string{
		0:   "Undefined",
		1:   "TODO",
		2:   "ErrNotImplemented",
		3:   "ErrInternal",
		100: "ErrInvalidInput",
		101: "ErrBridgeInterrupted",
		102: "ErrMissingInput",
		103: "ErrSerialization",
		104: "ErrDeserialization",
		105: "ErrInitService",
		106: "ErrSetRemote",
		107: "ErrKeyNameExists",
		150: "ErrCryptoKeyTypeUnknown",
		151: "ErrCryptoKeyNotFound",
		152: "ErrNoActiveAccount",
		153: "ErrRunGRPCServer",
		154: "ErrDecryptionFailed",
		200: "ErrTxDecode",
		201: "ErrInvalidSequence",
		202: "ErrUnauthorized",
		203: "ErrInsufficientFunds",
		204: "ErrUnknownRequest",
		205: "ErrInvalidAddress",
		206: "ErrUnknownAddress",
		207: "ErrInvalidPubKey",
		208: "ErrInsufficientCoins",
		209: "ErrInvalidCoins",
		210: "ErrInvalidGasWanted",
		211: "ErrOutOfGas",
		212: "ErrMemoTooLarge",
		213: "ErrInsufficientFee",
		214: "ErrTooManySignatures",
		215: "ErrNoSignatures",
		216: "ErrGasOverflow",
		217: "ErrInvalidPkgPath",
		218: "ErrInvalidStmt",
		219: "ErrInvalidExpr",
	}
	ErrCode_value = map[string]int32{
		"Undefined":               0,
		"TODO":                    1,
		"ErrNotImplemented":       2,
		"ErrInternal":             3,
		"ErrInvalidInput":         100,
		"ErrBridgeInterrupted":    101,
		"ErrMissingInput":         102,
		"ErrSerialization":        103,
		"ErrDeserialization":      104,
		"ErrInitService":          105,
		"ErrSetRemote":            106,
		"ErrKeyNameExists":        107,
		"ErrCryptoKeyTypeUnknown": 150,
		"ErrCryptoKeyNotFound":    151,
		"ErrNoActiveAccount":      152,
		"ErrRunGRPCServer":        153,
		"ErrDecryptionFailed":     154,
		"ErrTxDecode":             200,
		"ErrInvalidSequence":      201,
		"ErrUnauthorized":         202,
		"ErrInsufficientFunds":    203,
		"ErrUnknownRequest":       204,
		"ErrInvalidAddress":       205,
		"ErrUnknownAddress":       206,
		"ErrInvalidPubKey":        207,
		"ErrInsufficientCoins":    208,
		"ErrInvalidCoins":         209,
		"ErrInvalidGasWanted":     210,
		"ErrOutOfGas":             211,
		"ErrMemoTooLarge":         212,
		"ErrInsufficientFee":      213,
		"ErrTooManySignatures":    214,
		"ErrNoSignatures":         215,
		"ErrGasOverflow":          216,
		"ErrInvalidPkgPath":       217,
		"ErrInvalidStmt":          218,
		"ErrInvalidExpr":          219,
	}
)

// WithCode defines an error that can be used by helpers of this package.
type WithCode interface {
	error
	Code() ErrCode
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
