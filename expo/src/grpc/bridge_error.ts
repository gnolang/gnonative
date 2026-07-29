import { Code, ConnectError } from '@connectrpc/connect';

import { ErrCode, ErrDetailsSchema } from '../api/vendor/rpc_pb';

/**
 * Marker for a bridge rejection carrying a JSON envelope rather than plain text,
 * so one is told from an ordinary message without parsing every rejection as
 * JSON. Must match BridgeErrorPrefix in framework/service/bridge_error.go.
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
   * English text to show for this failure, with no framing to strip. Empty when
   * there is none to give. Usually the library's default wording for the code;
   * for ErrChainRejected, the reason the chain itself gave.
   */
  message: string;
};

/**
 * Rebuilds a ConnectError, with its ErrDetails, from a bridge rejection.
 *
 * The native bridge can only carry a string, so the Go side re-encodes the codes
 * into the rejection message; this puts them back where a caller expects them,
 * reachable with findDetails(ErrDetailsSchema) rather than by recognising a
 * shape in prose. Anything that is not one of these envelopes is returned
 * untouched.
 */
export function bridgeErrorToConnectError(error: unknown): unknown {
  const message = messageOf(error);
  if (!message?.startsWith(BRIDGE_ERROR_PREFIX)) return error;

  let envelope: BridgeErrorEnvelope;
  try {
    envelope = JSON.parse(message.slice(BRIDGE_ERROR_PREFIX.length)) as BridgeErrorEnvelope;
  } catch {
    // A malformed envelope must not lose the failure it describes.
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
 * bridge client gets it from bridgeErrorToConnectError. Neither looks at the
 * message text.
 *
 * Show `message`, or branch on `code` and use your own — an app knows which of
 * its screens fixes the problem, the library does not. Except for
 * ErrChainRejected, whose reason comes from the chain and cannot be bettered.
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
