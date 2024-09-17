export interface GoBridgeInterface {
  initBridge(): Promise<void>;
  closeBridge(): Promise<void>;
  getTcpPort(): Promise<number>;
  startGnokeyMobileService(): Promise<void>;
  invokeGrpcMethod(method: string, jsonMessage: string): Promise<string>;
  createStreamClient(method: string, jsonMessage: string): Promise<string>;
  streamClientReceive(id: string): Promise<string>;
  closeStreamClient(id: string): Promise<void>;
}
