import { EventEmitter, NativeModulesProxy, Subscription } from 'expo-modules-core';

// Import the native module. On web, it will be resolved to Gnonative.web.ts
// and on native platforms to Gnonative.ts
import { ChangeEventPayload, GnonativeViewProps } from './Gnonative.types';
import GnonativeModule from './GnonativeModule';
import GnonativeView from './GnonativeView';
import 'fast-text-encoding';

// Symbol.asyncIterator is needed by gRPC streams. Engines that provide it natively
// (e.g. Hermes V1) make it read-only, so assigning to it throws: only define it when missing.
if (!Symbol.asyncIterator) {
  Object.defineProperty(Symbol, 'asyncIterator', { value: Symbol.for('Symbol.asyncIterator') });
}

const emitter = new EventEmitter(GnonativeModule ?? NativeModulesProxy.Gnonative);

export function addChangeListener(listener: (event: ChangeEventPayload) => void): Subscription {
  return emitter.addListener<ChangeEventPayload>('onChange', listener);
}

export { ChangeEventPayload, GnonativeView, GnonativeViewProps };
export * from './provider/gnonative-provider';
export * from './api/vendor/gnonativetypes_pb';
export * from './api/vendor/rpc_pb';
export { GRPCError } from './grpc/error';
export { bridgeErrorToConnectError, errDetailOf } from './grpc/bridge_error';
export type { ErrDetail } from './grpc/bridge_error';

export * from './api/GnoNativeApi';
