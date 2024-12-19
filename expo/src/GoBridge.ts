import GnonativeModule from './GnonativeModule';

export interface GoBridgeInterface {
  initBridge(): Promise<void>;
  closeBridge(): Promise<void>;
  getTcpPort(): Promise<number>;
  invokeGrpcMethod(method: string, jsonMessage: string): Promise<string>;
  createStreamClient(method: string, jsonMessage: string): Promise<string>;
  streamClientReceive(id: string): Promise<string>;
  closeStreamClient(id: string): Promise<void>;
}

class GoBridge implements GoBridgeInterface {
  initBridge(): Promise<void> {
    return GnonativeModule.initBridge();
  }

  closeBridge(): Promise<void> {
    return GnonativeModule.closeBridge();
  }

  getTcpPort(): Promise<number> {
    return GnonativeModule.getTcpPort();
  }

  invokeGrpcMethod(method: string, jsonMessage: string): Promise<string> {
    return GnonativeModule.invokeGrpcMethod(method, jsonMessage);
  }

  createStreamClient(method: string, jsonMessage: string): Promise<string> {
    return GnonativeModule.createStreamClient(method, jsonMessage);
  }

  streamClientReceive(id: string): Promise<string> {
    return GnonativeModule.streamClientReceive(id);
  }

  closeStreamClient(id: string): Promise<void> {
    return GnonativeModule.closeStreamClient(id);
  }
}

const goBridge: GoBridgeInterface = new GoBridge();
export { goBridge as GoBridge };
