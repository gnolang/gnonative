// @gnolang/gnonative/native — the buf/connect-free entry point.
//
// This barrel must not import from ../grpc/*, ../api/vendor/*, or ../index so that Metro never pulls
// the buf/connect runtime into an app that only uses /native. It uses only GoBridge (which is
// buf-free) and the generated plain-JSON types.
import 'fast-text-encoding';

// Ensure Symbol.asyncIterator exists for `for await` over streams (RN/Hermes older runtimes).
(Symbol as { asyncIterator?: symbol }).asyncIterator =
  Symbol.asyncIterator || Symbol.for('Symbol.asyncIterator');

export { GnoNativeClient } from './client';
export { GnoNativeProvider, useGnoNativeContext } from './provider';
export type { GnoNativeContextProps } from './provider';
export { GnoNativeError } from './error';
export { bytesToBase64, base64ToBytes, base64ToString } from './encoding';
export { BridgeStatus } from './types';
export type { Config, GnoNativeClientApi } from './types';
export * from './types.gen';
