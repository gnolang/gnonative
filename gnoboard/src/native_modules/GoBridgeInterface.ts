export interface GoBridgeInterface {
  initBridge(): Promise<void>;
  closeBridge(): Promise<void>;
  listKeyInfo(): Promise<Array<string>>;
  createAccount(
    nameOrBech32: string,
    mnemonic: string,
    bip39Passw: string,
    password: string,
    account: Number,
    index: Number,
  ): Promise<string>;
  selectAccount(nameOrBech32: string): Promise<string>;
  call(
    packagePath: string,
    fnc: string,
    args: Array<string>,
    gasFee: string,
    gasWanted: Number,
    password: string,
  ): Promise<string>;
  exportJsonConfig(): Promise<string>;
}
