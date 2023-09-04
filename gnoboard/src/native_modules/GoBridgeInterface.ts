export interface GoBridgeInterface {
  initBridge(): Promise<void>;
  closeBridge(): Promise<void>;
  clientExec(command: string): Promise<string>;
  exportJsonConfig(): Promise<string>;
}
