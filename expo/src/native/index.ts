// The @gnolang/gnonative client surface. The root export (../index) re-exports this module, and the
// `./native` subpath is kept as an alias. It uses only GoBridge and the plain-JSON apitypes.
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
export * from './apitypes';
