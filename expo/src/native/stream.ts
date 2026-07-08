// AsyncIterable over the streaming bridge (createStream -> streamReceive -> closeStream).
// Termination: the native module rejects streamReceive
// with the literal message "EOF" (from Go's io.EOF) when the stream ends cleanly.
import { GoBridge } from '../GoBridge';
import { base64ToString } from './encoding';
import { GnoNativeError } from './error';

function isEOF(e: unknown): boolean {
  return typeof (e as { message?: unknown })?.message === 'string' && (e as Error).message === 'EOF';
}

// streamAsyncIterable starts a server-streaming method and yields each parsed JSON response.
// The stream is created lazily on first iteration and always closed when iteration ends.
export function streamAsyncIterable<T>(method: string, requestJson: string): AsyncIterable<T> {
  return {
    async *[Symbol.asyncIterator](): AsyncIterator<T> {
      const streamId = await GoBridge.createStream(method, requestJson);
      try {
        for (;;) {
          let res: string;
          try {
            res = await GoBridge.streamReceive(streamId);
          } catch (e) {
            if (isEOF(e)) {
              break;
            }
            throw new GnoNativeError((e as Error)?.message ?? String(e));
          }
          yield JSON.parse(base64ToString(res)) as T;
        }
      } finally {
        try {
          await GoBridge.closeStream(streamId);
        } catch {
          // Ignore close errors: the stream may already be finished/unregistered.
        }
      }
    },
  };
}
