export interface GoBridgeInterface {
  initBridge(): Promise<void>;
  closeBridge(): Promise<void>;
  call(packagePath: string, fnc: string, args: Array<string>, gasFee: string, gasWanted: Number, password: string): Promise<string>;
  exportJsonConfig(): Promise<string>;
}
