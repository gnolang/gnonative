export interface GoBridgeInterface {
  initBridge(): Promise<void>;
  closeBridge(): Promise<void>;
  getTcpPort(): Promise<number>;
}
