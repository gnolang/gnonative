import { GnoAccount } from "./types";

export interface GoBridgeInterface {
  initBridge(): Promise<void>;
  closeBridge(): Promise<void>;
  setPassword(password: string): Promise<void>;
  generateRecoveryPhrase(): Promise<string>;
  listKeyInfo(): Promise<Array<GnoAccount>>;
  createAccount(
    nameOrBech32: string,
    mnemonic: string,
    bip39Passw: string,
    password: string,
    account: number,
    index: number,
  ): Promise<GnoAccount>;
  selectAccount(nameOrBech32: string): Promise<GnoAccount>;
  getActiveAccount(): Promise<GnoAccount>;
  call(
    packagePath: string,
    fnc: string,
    args: Array<string>,
    gasFee: string,
    gasWanted: number,
  ): Promise<string>;
  exportJsonConfig(): Promise<string>;
  getTcpPort(): Promise<number>;
}
