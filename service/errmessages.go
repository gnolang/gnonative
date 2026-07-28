package service

import (
	api_gen "github.com/gnolang/gnonative/v4/api/gen/go"
)

// errCodeMessages is the default text sent with each code a client may need to
// show someone.
//
// It lives here, in the service, so that every language gets it: the code and
// its wording travel together in ErrDetails, over the schema published to the
// buf registry. A client that wants different wording maps the code itself, but
// no client has to invent text for a code it does not recognise, and none has to
// re-derive what the code means.
//
// Each entry is written from the code that raises it — in this repo, in
// tm2/pkg/std and in gno.land/pkg/sdk/vm — not from the code's name, because
// naming alone produces confidently wrong text. Where one code covers several
// situations the wording stays at the level they share.
//
// The text says what happened and stops there. What to do about it names screens
// only an application knows exist, so it belongs to that application.
//
// Codes only a developer can act on (ErrNotImplemented, ErrSerialization,
// ErrRunGRPCServer, …) are deliberately absent: inventing user-facing text for
// them puts a reassuring sentence in front of what is really a bug. An absent
// code yields an empty message, which a client can tell apart from a real one.
var errCodeMessages = map[api_gen.ErrCode]string{
	// Transport: the request never reached the node.
	api_gen.ErrCode_ErrRemoteUnreachable: "The node could not be reached.",
	// Reached only when constructing the RPC client fails, so the address itself
	// is unusable — this is not a reachability failure.
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
	// Covers a missing key as well as an unrecognised key type, so this does not
	// claim the key is merely wrong for this chain.
	api_gen.ErrCode_ErrInvalidPubKey: "The public key is missing or of an unsupported type.",

	// Transactions.
	// Deliberately no "raise the gas limit": tm2 also raises this when the block
	// has no gas left ("no block gas left to run tx"), where that advice is wrong.
	api_gen.ErrCode_ErrOutOfGas: "The transaction ran out of gas.",
	// An overflow while summing gas, not an out-of-range limit.
	api_gen.ErrCode_ErrGasOverflow: "The transaction gas amounts are too large to add up.",
	// "invalid gas-wanted; got: N block-max-gas: M" — over the block maximum, or
	// otherwise unusable.
	api_gen.ErrCode_ErrInvalidGasWanted: "The gas limit is not valid for this chain.",
	// Specifically the fee, per "insufficient funds to pay for fees".
	api_gen.ErrCode_ErrInsufficientFunds: "There are not enough funds to cover the transaction fee.",
	// Two causes: a malformed fee amount, and a fee below the node's gas price.
	api_gen.ErrCode_ErrInsufficientFee: "The transaction fee is not valid or is too low.",
	// "insufficient account funds; X < Y", but also "send amount must be positive".
	api_gen.ErrCode_ErrInsufficientCoins: "The amount is not positive, or the account does not hold enough.",
	// Coins.IsValid() covers ordering and denomination as well as amount, so this
	// must not be described as an amount problem alone.
	api_gen.ErrCode_ErrInvalidCoins: "The coin amount, denomination or ordering is not valid.",
	// Raised for signature and session-authority problems. Says nothing about
	// whether anything was broadcast, so it must not claim the chain rejected it.
	api_gen.ErrCode_ErrUnauthorized: "The transaction is not authorised.",
	// Usually a stale sequence because another transaction landed first, but tm2
	// also raises it for a malformed counter — so state the fact, not the cause.
	api_gen.ErrCode_ErrInvalidSequence: "The account sequence is not valid.",
	// Raised by the app router for any unrecognised message or query path, not
	// only for a realm function that does not exist.
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

// messageForErrCode returns the default text for a code, or "" when there is
// none to give.
func messageForErrCode(code api_gen.ErrCode) string {
	return errCodeMessages[code]
}
