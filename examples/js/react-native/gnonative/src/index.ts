import { NativeModulesProxy, EventEmitter, Subscription } from 'expo-modules-core';

// Import the native module. On web, it will be resolved to Gnonative.web.ts
// and on native platforms to Gnonative.ts
import GnonativeModule from './GnonativeModule';
import GnonativeView from './GnonativeView';
import { ChangeEventPayload, GnonativeViewProps } from './Gnonative.types';

// Get the native constant value.
export const PI = GnonativeModule.PI;

export function hello(): string {
  return GnonativeModule.hello();
}

export async function setValueAsync(value: string) {
  return await GnonativeModule.setValueAsync(value);
}

const emitter = new EventEmitter(GnonativeModule ?? NativeModulesProxy.Gnonative);

export function addChangeListener(listener: (event: ChangeEventPayload) => void): Subscription {
  return emitter.addListener<ChangeEventPayload>('onChange', listener);
}

export { GnonativeView, GnonativeViewProps, ChangeEventPayload };
