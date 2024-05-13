import { EventEmitter, NativeModulesProxy, Subscription } from 'expo-modules-core';

// Import the native module. On web, it will be resolved to Gnonative.web.ts
// and on native platforms to Gnonative.ts
import { ChangeEventPayload, GnonativeViewProps } from './Gnonative.types';
import GnonativeModule from './GnonativeModule';
import GnonativeView from './GnonativeView';
import 'fast-text-encoding';

(Symbol as any).asyncIterator = Symbol.asyncIterator || Symbol.for('Symbol.asyncIterator');

const emitter = new EventEmitter(GnonativeModule ?? NativeModulesProxy.Gnonative);

export function addChangeListener(listener: (event: ChangeEventPayload) => void): Subscription {
  return emitter.addListener<ChangeEventPayload>('onChange', listener);
}

export { ChangeEventPayload, GnonativeView, GnonativeViewProps };
export { useGno } from './hooks/use-gno';
export * from './provider/gnokey-provider';
export * from '@buf/gnolang_gnonative.bufbuild_es/gnonativetypes_pb';
