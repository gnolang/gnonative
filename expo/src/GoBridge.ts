import GnonativeModule from './GnonativeModule';

// Options for initBridgeWithOptions (the /native path). useGrpcServers defaults to true.
export interface InitBridgeOptions {
  useGrpcServers?: boolean;
}

export interface GoBridgeInterface {
  initBridge(): Promise<void>;
  // initBridgeWithOptions starts the bridge and, when useGrpcServers is false, skips the in-process
  // connect/gRPC servers entirely (the /native path serves JS through invokeMethod/createStream).
  initBridgeWithOptions(options: InitBridgeOptions): Promise<void>;
  closeBridge(): Promise<void>;
  getTcpPort(): Promise<number>;
  invokeGrpcMethod(method: string, jsonMessage: string): Promise<string>;
  createStreamClient(method: string, jsonMessage: string): Promise<string>;
  streamClientReceive(id: string): Promise<string>;
  closeStreamClient(id: string): Promise<void>;
  // Connect/gRPC-free dispatcher API (base64(protojson) payloads).
  invokeMethod(method: string, jsonMessage: string): Promise<string>;
  createStream(method: string, jsonMessage: string): Promise<string>;
  streamReceive(id: string): Promise<string>;
  closeStream(id: string): Promise<void>;
}

class GoBridge implements GoBridgeInterface {
  initBridge(): Promise<void> {
    return GnonativeModule.initBridge();
  }

  initBridgeWithOptions(options: InitBridgeOptions): Promise<void> {
    return GnonativeModule.initBridgeWithOptions(options);
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
