import { Code, ConnectError } from '@connectrpc/connect';

import { ErrCode, ErrDetailsSchema } from '../api/vendor/rpc_pb';

/**
 * Marker for a bridge rejection carrying a JSON envelope rather than plain text.
 *
 * Distinguishes an envelope from an ordinary message without parsing every
 * rejection as JSON to find out. Must match BridgeErrorPrefix in
 * framework/service/bridge_error.go.
 */
const BRIDGE_ERROR_PREFIX = 'gnonative-error:';

type BridgeErrorEnvelope = {
  detail?: { code?: number; message?: string };
  error?: string;
};

/** The ErrCode a failure carries, with the text sent alongside it. */
export type ErrDetail = {
  code: ErrCode;
  /**
   * English text to show for this failure. Empty when there is none to give.
   *
   * Usually the library's default wording for the code. For ErrChainRejected it
   * is instead the reason the chain itself gave — a realm's panic message, say —
   * because that is what differs per failure and no constant could replace it.
   * Either way it is meant to be shown as it arrives, with no framing to strip.
   */
  message: string;
};

/**
 * Rebuilds a ConnectError, with its ErrDetails, from a bridge rejection.
 *
 * The native bridge can only carry a string to JavaScript, so a connect error's
 * details do not survive it. The Go side therefore re-encodes the codes into the
 * rejection message, and this puts them back where a caller expects them: on a
 * ConnectError, as the same ErrDetails message the schema publishes — reachable
 * with findDetails(ErrDetailsSchema) rather than by recognising a shape in prose.
 *
 * Anything that is not one of these envelopes is returned untouched, so a
 * rejection from elsewhere in the bridge keeps its own message.
 */
export function bridgeErrorToConnectError(error: unknown): unknown {
  const message = messageOf(error);
  if (!message?.startsWith(BRIDGE_ERROR_PREFIX)) return error;

  let envelope: BridgeErrorEnvelope;
  try {
    envelope = JSON.parse(message.slice(BRIDGE_ERROR_PREFIX.length)) as BridgeErrorEnvelope;
  } catch {
    // A malformed envelope must not lose the failure it was describing.
    return error;
  }

  const code = envelope.detail?.code;
  const hasCode = typeof code === 'number' && code in ErrCode;

  return new ConnectError(
    envelope.error ?? message,
    Code.Unknown,
    undefined,
    hasCode
      ? [{ desc: ErrDetailsSchema, value: { code, message: envelope.detail?.message ?? '' } }]
      : undefined,
    error,
  );
}

/**
 * The ErrDetail a failure carries, or undefined when it carries none.
 *
 * Works for both transports: a networked client reads the detail off the wire, a
 * bridge client gets it from bridgeErrorToConnectError. Neither has to look at
 * the message text.
 *
 * `message` is wording meant to be shown. Show it, or branch on `code` and use
 * your own — an app that knows which of its screens fixes the problem can say
 * so, which the library cannot. One code resists that: ErrChainRejected carries
 * the chain's own reason, which no client can improve on and none should try to
 * replace with a constant.
 */
export function errDetailOf(error: unknown): ErrDetail | undefined {
  if (!(error instanceof ConnectError)) return undefined;

  for (const details of error.findDetails(ErrDetailsSchema)) {
    if (details.code in ErrCode) {
      return { code: details.code as ErrCode, message: details.message };
    }
  }
  return undefined;
}

function messageOf(error: unknown): string | undefined {
  if (typeof error === 'string') return error;
  if (error instanceof Error) return error.message;
  if (error && typeof error === 'object') {
    const message = (error as { message?: unknown }).message;
    if (typeof message === 'string') return message;
  }
  return undefined;
}
