import { EventEmitter, NativeModulesProxy, Subscription } from 'expo-modules-core';

// Import the native module. On web, it will be resolved to Gnonative.web.ts
// and on native platforms to Gnonative.ts
import GnonativeModule from './GnonativeModule';
import GnonativeView from './GnonativeView';
import { ChangeEventPayload, GnonativeViewProps } from './Gnonative.types';

export function initBridge(): Promise<void> {
  return GnonativeModule.initBridge();
}

export function closeBridge(): Promise<void> {
  return GnonativeModule.closeBridge();
}

export function getTcpPort(): Promise<number> {
  return GnonativeModule.getTcpPort();
}

export function invokeGrpcMethod(method: string, jsonMessage: string): Promise<string> {
  return GnonativeModule.invokeGrpcMethod(method, jsonMessage);
}

export function createStreamClient(method: string, jsonMessage: string): Promise<string> {
  return GnonativeModule.createStreamClient(method, jsonMessage);
}

export function streamClientReceive(): Promise<string> {
  return GnonativeModule.streamClientReceive();
}

export function closeStreamClient(id: string): Promise<void> {
  return GnonativeModule.streamClientReceive(id);
}

const emitter = new EventEmitter(GnonativeModule ?? NativeModulesProxy.Gnonative);

export function addChangeListener(listener: (event: ChangeEventPayload) => void): Subscription {
  return emitter.addListener<ChangeEventPayload>('onChange', listener);
}

export { ChangeEventPayload, GnonativeView, GnonativeViewProps };
