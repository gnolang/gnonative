import {Code, ConnectError} from '@connectrpc/connect';

/**
 * Marker for a bridge rejection carrying a JSON envelope rather than plain text.
 * Must match BridgeErrorPrefix in framework/service/bridge_error.go.
 */
const BRIDGE_ERROR_PREFIX = 'gnonative-error:';

type BridgeErrorEnvelope = {
  detail?: {code?: number; message?: string};
  error?: string;
  connectCode?: number;
};

/**
 * Unwraps a bridge rejection into a ConnectError, or returns it untouched.
 *
 * This example pins a generated API from before ErrDetails carried a code, so it
 * takes only the message out of the envelope rather than rebuilding the typed
 * detail. Use @gnolang/gnonative's errDetailOf if you want the code itself.
 */
export function bridgeErrorToConnectError(error: unknown): unknown {
  const message = messageOf(error);
  if (!message?.startsWith(BRIDGE_ERROR_PREFIX)) return error;

  let envelope: BridgeErrorEnvelope;
  try {
    envelope = JSON.parse(
      message.slice(BRIDGE_ERROR_PREFIX.length),
    ) as BridgeErrorEnvelope;
  } catch {
    // A malformed envelope must not lose the failure it describes.
    return error;
  }

  const code = envelope.connectCode;
  return new ConnectError(
    envelope.error ?? envelope.detail?.message ?? message,
    typeof code === 'number' && code in Code ? code : Code.Unknown,
    undefined,
    undefined,
    error,
  );
}

function messageOf(error: unknown): string | undefined {
  if (typeof error === 'string') return error;
  if (error instanceof Error) return error.message;
  if (error && typeof error === 'object') {
    const message = (error as {message?: unknown}).message;
    if (typeof message === 'string') return message;
  }
  return undefined;
}
