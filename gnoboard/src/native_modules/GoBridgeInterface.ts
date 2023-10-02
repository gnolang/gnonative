export interface GoBridgeInterface {
  initBridge(): Promise<void>;
  closeBridge(): Promise<void>;
  setPassword(password: string): Promise<void>;
  generateRecoveryPhrase(): Promise<string>;
  listKeyInfo(): Promise<Array<Object>>;
  createAccount(
    nameOrBech32: string,
    mnemonic: string,
    bip39Passw: string,
    password: string,
    account: Number,
    index: Number,
  ): Promise<Object>;
  selectAccount(nameOrBech32: string): Promise<Object>;
  call(
    packagePath: string,
    fnc: string,
    args: Array<string>,
    gasFee: string,
    gasWanted: Number,
  ): Promise<string>;
  exportJsonConfig(): Promise<string>;
}
