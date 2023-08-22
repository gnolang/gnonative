export interface GoBridgeInterface {
  initBridge(): Promise<void>;
  closeBridge(): Promise<void>;
  hello(name: string): Promise<string>;
  exportJsonConfig(): Promise<string>;
}
