package service

import (
	"errors"

	abci "github.com/gnolang/gno/tm2/pkg/bft/abci/types"
	api_gen "github.com/gnolang/gnonative/v4/api/gen/go"
)

// errCodeMessages is the default text sent with each code a client may need to
// show someone.
//
// It lives in the service so every language gets it: code and wording travel
// together in ErrDetails. A client may map the code to its own wording, but none
// has to invent text for a code it does not recognise.
//
// Each entry is written from the code that raises it — here, in tm2/pkg/std and
// in gno.land/pkg/sdk/vm — not from the code's name, which produces confidently
// wrong text (see the comments below). The text says what happened and stops
// there: the remedy names screens only an application knows exist.
//
// Codes only a developer can act on (ErrNotImplemented, ErrSerialization, …) are
// deliberately absent — a reassuring sentence in front of a bug is worse than
// none — and an absent code yields an empty message a client can tell apart from
// a real one. ErrChainRejected is absent for the opposite reason: its text
// exists but differs per failure, so messageForError supplies it.
var errCodeMessages = map[api_gen.ErrCode]string{
	// Transport: the request never reached the node.
	api_gen.ErrCode_ErrRemoteUnreachable: "The node could not be reached.",
	// Only when constructing the RPC client fails: the address itself is unusable,
	// which is not a reachability failure.
	api_gen.ErrCode_ErrSetRemote: "That node address could not be used.",

	// Keys and accounts.
	// Any keybase lookup by name or address, not only a signing one.
	api_gen.ErrCode_ErrCryptoKeyNotFound: "No key with that name or address was found.",
	api_gen.ErrCode_ErrDecryptionFailed:  "The key could not be unlocked. The password may be wrong.",
	api_gen.ErrCode_ErrNoActiveAccount:   "No account has been activated.",
	api_gen.ErrCode_ErrKeyNameExists:     "An account with that name already exists.",
	// Raised for a blank address as well as one whose bech32 cannot be decoded.
	api_gen.ErrCode_ErrInvalidAddress: "The account address is missing or not valid.",
	api_gen.ErrCode_ErrUnknownAddress: "That account does not exist on this chain.",
	// Covers a missing key as well as an unrecognised type.
	api_gen.ErrCode_ErrInvalidPubKey: "The public key is missing or of an unsupported type.",

	// Transactions.
	// No "raise the gas limit": tm2 also raises this for "no block gas left to run
	// tx", where that advice is wrong.
	api_gen.ErrCode_ErrOutOfGas: "The transaction ran out of gas.",
	// An overflow while summing gas, not an out-of-range limit.
	api_gen.ErrCode_ErrGasOverflow: "The transaction gas amounts are too large to add up.",
	// "invalid gas-wanted; got: N block-max-gas: M" — over the block maximum.
	api_gen.ErrCode_ErrInvalidGasWanted: "The gas limit is not valid for this chain.",
	// Specifically the fee, per "insufficient funds to pay for fees".
	api_gen.ErrCode_ErrInsufficientFunds: "There are not enough funds to cover the transaction fee.",
	// Two causes: a malformed fee amount, and a fee below the node's gas price.
	api_gen.ErrCode_ErrInsufficientFee: "The transaction fee is not valid or is too low.",
	// "insufficient account funds; X < Y", but also "send amount must be positive".
	api_gen.ErrCode_ErrInsufficientCoins: "The amount is not positive, or the account does not hold enough.",
	// Coins.IsValid() covers ordering and denomination, not only the amount.
	api_gen.ErrCode_ErrInvalidCoins: "The coin amount, denomination or ordering is not valid.",
	// Signature and session-authority problems. Says nothing about what was
	// broadcast, so it must not claim the chain rejected anything.
	api_gen.ErrCode_ErrUnauthorized: "The transaction is not authorised.",
	// Usually a stale sequence, but also a malformed counter — state the fact.
	api_gen.ErrCode_ErrInvalidSequence: "The account sequence is not valid.",
	// The app router, for any unrecognised message or query path.
	api_gen.ErrCode_ErrUnknownRequest:    "The chain did not recognise the request.",
	api_gen.ErrCode_ErrTxDecode:          "The transaction could not be decoded.",
	api_gen.ErrCode_ErrMemoTooLarge:      "The transaction memo is too large.",
	api_gen.ErrCode_ErrTooManySignatures: "The transaction has more signatures than allowed.",
	api_gen.ErrCode_ErrNoSignatures:      "The transaction is not signed.",

	// Realm calls.
	// "missing package path", "pkgpath must be of a realm", "must not be internal".
	api_gen.ErrCode_ErrInvalidPkgPath: "The package path is missing or not valid.",
	api_gen.ErrCode_ErrInvalidStmt:    "The statement is not valid.",
	api_gen.ErrCode_ErrInvalidExpr:    "The expression is not valid.",
}

// messageForError returns the text to send with a classified failure, or "" when
// there is none to give.
//
// ErrChainRejected is the one code whose wording is not in the table: the reason
// the chain refused a request differs per failure, so it is read back out of the
// error where getGrpcError left it. No sentence this library could write is
// worth as much to someone as "thread body is required".
func messageForError(err error, code api_gen.ErrCode) string {
	if code == api_gen.ErrCode_ErrChainRejected {
		var rejected abci.StringError
		if errors.As(err, &rejected) {
			return rejected.Error()
		}
		// Unreachable: the code is only set where that error was found.
		return ""
	}

	return errCodeMessages[code]
}
