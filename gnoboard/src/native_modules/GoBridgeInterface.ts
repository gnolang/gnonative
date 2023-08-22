export interface GoBridgeInterface {
  initBridge(): Promise<void>;
  closeBridge(): Promise<void>;
  getAccountAndTxCfg(): Promise<string>;
  hello(name: string): Promise<string>;
}
