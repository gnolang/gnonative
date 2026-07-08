import GnonativeModule from './GnonativeModule';

export interface GoBridgeInterface {
  initBridge(): Promise<void>;
  closeBridge(): Promise<void>;
  // Dispatcher API: unary calls and server-streaming. Payloads are base64(JSON).
  invokeMethod(method: string, jsonMessage: string): Promise<string>;
  createStream(method: string, jsonMessage: string): Promise<string>;
  streamReceive(id: string): Promise<string>;
  closeStream(id: string): Promise<void>;
}

class GoBridge implements GoBridgeInterface {
  initBridge(): Promise<void> {
    return GnonativeModule.initBridge();
  }

  closeBridge(): Promise<void> {
    return GnonativeModule.closeBridge();
  }

  invokeMethod(method: string, jsonMessage: string): Promise<string> {
    return GnonativeModule.invokeMethod(method, jsonMessage);
  }

  createStream(method: string, jsonMessage: string): Promise<string> {
    return GnonativeModule.createStream(method, jsonMessage);
  }

  streamReceive(id: string): Promise<string> {
    return GnonativeModule.streamReceive(id);
  }

  closeStream(id: string): Promise<void> {
    return GnonativeModule.closeStream(id);
  }
}

const goBridge: GoBridgeInterface = new GoBridge();
export { goBridge as GoBridge };
