import { ErrCode } from '../api/vendor/rpc_pb';

/**
 * Default, human-readable text for the error codes gnonative classifies.
 *
 * These describe *what happened* at the chain or transport level — a fact the
 * library knows and every consumer would otherwise have to phrase for itself.
 * They deliberately stop short of *what to do about it*, because the remedy is
 * app-specific: "raise the gas limit in Settings" names a screen only the app
 * knows exists. Apps add that through the `overrides` argument to
 * `describeErrCode`.
 *
 * The wording is deliberately plain and free of product voice, so an app that
 * overrides an entry for tone is replacing something neutral rather than
 * arguing with it.
 *
 * Not every ErrCode appears here. Codes that only a developer can act on
 * (`ErrNotImplemented`, `ErrSerialization`, `ErrRunGRPCServer`, …) are left out
 * on purpose: inventing user-facing text for them would put a reassuring
 * sentence in front of what is really a bug.
 *
 * These are written from the code that raises each one, not from the code's
 * name, and they state only what that code actually covers. Where tm2 raises one
 * code for several situations the wording stays at the level they share: an
 * `ErrOutOfGas` may be the transaction's own gas or the block's, so "raise the
 * gas limit" would be wrong half the time, and belongs to an app that knows
 * which case it is.
 *
 * They also do not attempt the specifics. gno's own error values carry those in
 * the message wrapped at the raise site ("insufficient funds to pay for fees;
 * 100ugnot < 200ugnot"), which gnonative already passes through in the error
 * text — but that text is developer-facing, unstable, and not localisable, so it
 * is context to show *alongside* one of these sentences, never a replacement.
 */
export const DEFAULT_ERR_CODE_MESSAGES: Readonly<Partial<Record<ErrCode, string>>> = Object.freeze({
  // Transport: the request never reached the node.
  [ErrCode.ErrRemoteUnreachable]: 'The node could not be reached.',
  // getGrpcError only reaches this when constructing the RPC client fails, so
  // the address itself is unusable — this is not a reachability failure.
  [ErrCode.ErrSetRemote]: 'That node address could not be used.',

  // Keys and accounts.
  // Any keybase lookup by name or address, not only a signing one.
  [ErrCode.ErrCryptoKeyNotFound]: 'No key with that name or address was found.',
  [ErrCode.ErrDecryptionFailed]: 'The key could not be unlocked. The password may be wrong.',
  [ErrCode.ErrNoActiveAccount]: 'No account has been activated.',
  [ErrCode.ErrKeyNameExists]: 'An account with that name already exists.',
  // Raised for a blank address as well as one whose bech32 cannot be decoded.
  [ErrCode.ErrInvalidAddress]: 'The account address is missing or not valid.',
  [ErrCode.ErrUnknownAddress]: 'That account does not exist on this chain.',
  // Covers a missing key as well as an unrecognised key type, so this does not
  // claim the key is merely wrong for this chain.
  [ErrCode.ErrInvalidPubKey]: 'The public key is missing or of an unsupported type.',

  // Transactions.
  // Deliberately no "raise the gas limit": tm2 also raises this when the block
  // has no gas left ("no block gas left to run tx"), where that advice is wrong.
  [ErrCode.ErrOutOfGas]: 'The transaction ran out of gas.',
  // An overflow while summing gas, not an out-of-range limit.
  [ErrCode.ErrGasOverflow]: 'The transaction gas amounts are too large to add up.',
  // "invalid gas-wanted; got: N block-max-gas: M" — over the block maximum, or
  // otherwise unusable.
  [ErrCode.ErrInvalidGasWanted]: 'The gas limit is not valid for this chain.',
  // Specifically the fee, per "insufficient funds to pay for fees".
  [ErrCode.ErrInsufficientFunds]: 'There are not enough funds to cover the transaction fee.',
  // Two causes: a malformed fee amount, and a fee below the node's gas price.
  [ErrCode.ErrInsufficientFee]: 'The transaction fee is not valid or is too low.',
  // "insufficient account funds; X < Y", but also "send amount must be positive".
  [ErrCode.ErrInsufficientCoins]:
    'The amount is not positive, or the account does not hold enough.',
  // `Coins.IsValid()` covers ordering and denomination as well as amount, so
  // this must not be described as an amount problem alone.
  [ErrCode.ErrInvalidCoins]: 'The coin amount, denomination or ordering is not valid.',
  // Raised for signature and session-authority problems. Says nothing about
  // whether anything was broadcast, so it must not claim the chain rejected it.
  [ErrCode.ErrUnauthorized]: 'The transaction is not authorised.',
  // Usually a stale sequence because another transaction landed first, but tm2
  // also raises it for a malformed counter — so state the fact, not the cause.
  [ErrCode.ErrInvalidSequence]: 'The account sequence is not valid.',
  // Raised by the app router for any unrecognised message or query path, not
  // only for a realm function that does not exist.
  [ErrCode.ErrUnknownRequest]: 'The chain did not recognise the request.',
  [ErrCode.ErrTxDecode]: 'The transaction could not be decoded.',
  [ErrCode.ErrMemoTooLarge]: 'The transaction memo is too large.',
  [ErrCode.ErrTooManySignatures]: 'The transaction has more signatures than allowed.',
  [ErrCode.ErrNoSignatures]: 'The transaction is not signed.',

  // Realm calls.
  // "missing package path", "pkgpath must be of a realm", "must not be internal".
  [ErrCode.ErrInvalidPkgPath]: 'The package path is missing or not valid.',
  [ErrCode.ErrInvalidStmt]: 'The statement is not valid.',
  [ErrCode.ErrInvalidExpr]: 'The expression is not valid.',
});

/**
 * The code gnonative embedded in an error message, or undefined.
 *
 * `ErrCode.Error()` renders as `ErrOutOfGas(#211)`, so the code travels inside
 * the text. Recovering it from a plain string — rather than from a live
 * `ConnectError` as `GRPCError.errCode()` does — is what a consumer needs when
 * the error has already been flattened: Redux Toolkit, for instance, serialises
 * a rejected thunk to `{ name, message, stack }` before any middleware sees it,
 * and the original error object is gone by then.
 */
export function errCodeFromMessage(message: string): ErrCode | undefined {
  const match = message?.match(/\(#(\d+)\)/);
  if (!match) return undefined;

  const code = Number(match[1]);
  if (!Number.isInteger(code)) return undefined;

  return code in ErrCode ? (code as ErrCode) : undefined;
}

/**
 * Text for a code, preferring the caller's wording.
 *
 * Returns undefined when neither the overrides nor the defaults cover the code,
 * so a caller can tell "nothing sensible to say" from "a generic default" and
 * fall back to whatever context it has. It never invents a message.
 */
export function describeErrCode(
  code: ErrCode,
  overrides?: Partial<Record<ErrCode, string>>,
): string | undefined {
  return overrides?.[code] ?? DEFAULT_ERR_CODE_MESSAGES[code];
}
