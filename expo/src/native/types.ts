// Public types for the @gnolang/gnonative/native entry point. Mirrors src/api/types.ts (GnoKeyApi)
// but Json-typed: bytes are base64 strings, int64 values accept/return strings, and responses are the
// plain-JSON `*Json` shapes from types.gen.ts. Migrating from the gRPC path is an import swap plus
// adapting Uint8Array/bigint to string.
import type {
  ActivateAccountResponseJson,
  BroadcastTxCommitResponseJson,
  CallResponseJson,
  CoinJson,
  CreateSessionResponseJson,
  DeleteAccountResponseJson,
  EstimateGasResponseJson,
  EstimateTxFeesResponseJson,
  GetActivatedAccountResponseJson,
  HelloStreamResponseJson,
  KeyInfoJson,
  MakeTxResponseJson,
  QueryAccountResponseJson,
  QueryResponseJson,
  QuerySessionAccountResponseJson,
  RenameKeyResponseJson,
  RevokeAllSessionsResponseJson,
  RevokeSessionResponseJson,
  RotatePasswordResponseJson,
  RunResponseJson,
  SendResponseJson,
  SetChainIDResponseJson,
  SetPasswordResponseJson,
  SetRemoteResponseJson,
  SignTxResponseJson,
} from './types.gen';

export enum BridgeStatus {
  Stopped,
  Starting,
  Started,
}

export interface Config {
  remote: string;
  chain_id: string;
}

// GnoNativeClientApi is the connect-free client surface. Addresses/keys are base64 strings, int64
// amounts (gasWanted, expiresAt, spendPeriod, ...) are strings, and responses are `*Json` shapes.
export interface GnoNativeClientApi {
  setRemote: (remote: string) => Promise<SetRemoteResponseJson>;
  getRemote: () => Promise<string>;
  setChainID: (chainId: string) => Promise<SetChainIDResponseJson>;
  getChainID: () => Promise<string>;
  createAccount: (
    nameOrBech32: string,
    mnemonic: string,
    password: string,
    bip39Passwd?: string,
    account?: number,
    index?: number,
  ) => Promise<KeyInfoJson | undefined>;
  createLedger: (
    name: string,
    algorithm: string,
    hrp: string,
    account?: number,
    index?: number,
  ) => Promise<KeyInfoJson | undefined>;
  generateRecoveryPhrase: () => Promise<string>;
  listKeyInfo: () => Promise<KeyInfoJson[]>;
  hasKeyByName: (name: string) => Promise<boolean>;
  hasKeyByAddress: (address: string) => Promise<boolean>;
  hasKeyByNameOrAddress: (nameOrBech32: string) => Promise<boolean>;
  getKeyInfoByName: (name: string) => Promise<KeyInfoJson | undefined>;
  getKeyInfoByAddress: (address: string) => Promise<KeyInfoJson | undefined>;
  getKeyInfoByNameOrAddress: (nameOrBech32: string) => Promise<KeyInfoJson | undefined>;
  activateAccount: (nameOrBech32: string) => Promise<ActivateAccountResponseJson>;
  setPassword: (password: string, address: string) => Promise<SetPasswordResponseJson>;
  renameKey: (oldName: string, newName: string) => Promise<RenameKeyResponseJson>;
  rotatePassword: (password: string, addresses: string[]) => Promise<RotatePasswordResponseJson>;
  getActivatedAccount: (address: string) => Promise<GetActivatedAccountResponseJson>;
  queryAccount: (address: string) => Promise<QueryAccountResponseJson>;
  querySessionAccount: (
    masterAddress: string,
    sessionAddress: string,
  ) => Promise<QuerySessionAccountResponseJson>;
  deleteAccount: (
    nameOrBech32: string,
    password: string | undefined,
    skipPassword: boolean,
  ) => Promise<DeleteAccountResponseJson>;
  query: (path: string, data: string) => Promise<QueryResponseJson>;
  render: (packagePath: string, args: string) => Promise<string>;
  qEval: (packagePath: string, expression: string) => Promise<string>;
  call: (
    packagePath: string,
    fnc: string,
    args: string[],
    gasFee: string,
    gasWanted: string,
    signerAddress: string,
    send?: CoinJson[],
    maxDeposit?: CoinJson[],
    memo?: string,
  ) => Promise<AsyncIterable<CallResponseJson>>;
  send: (
    toAddress: string,
    amount: CoinJson[],
    gasFee: string,
    gasWanted: string,
    signerAddress: string,
    memo?: string,
  ) => Promise<AsyncIterable<SendResponseJson>>;
  run: (
    pkg: string,
    gasFee: string,
    gasWanted: string,
    signerAddress: string,
    send?: CoinJson[],
    maxDeposit?: CoinJson[],
    memo?: string,
  ) => Promise<AsyncIterable<RunResponseJson>>;
  addressToBech32: (address: string) => Promise<string>;
  addressFromMnemonic: (mnemonic: string) => Promise<string>;
  addressFromBech32: (bech32Address: string) => Promise<string>;
  pubKeyBytesFromBech32: (bech32PubKey: string) => Promise<string>;
  validateMnemonicWord: (word: string) => Promise<boolean>;
  validateMnemonicPhrase: (phrase: string) => Promise<boolean>;
  signTx: (
    txJson: string,
    address: string,
    accountNumber?: string,
    sequenceNumber?: string,
  ) => Promise<SignTxResponseJson>;
  estimateGas: (
    txJson: string,
    address: string,
    securityMargin: number,
    updateTx: boolean,
  ) => Promise<EstimateGasResponseJson>;
  estimateTxFees: (
    txJson: string,
    address: string,
    gasSecurityMargin: number,
    gasPriceSecurityMargin: number,
    updateTx: boolean,
  ) => Promise<EstimateTxFeesResponseJson>;
  makeCallTx: (
    packagePath: string,
    fnc: string,
    args: string[],
    gasFee: string,
    gasWanted: string,
    callerAddress: string,
    send?: CoinJson[],
    maxDeposit?: CoinJson[],
    memo?: string,
  ) => Promise<MakeTxResponseJson>;
  makeSendTx: (
    toAddress: string,
    amount: CoinJson[],
    gasFee: string,
    gasWanted: string,
    callerAddress: string,
    memo?: string,
  ) => Promise<MakeTxResponseJson>;
  makeRunTx: (
    pkg: string,
    gasFee: string,
    gasWanted: string,
    callerAddress: string,
    send?: CoinJson[],
    maxDeposit?: CoinJson[],
    memo?: string,
  ) => Promise<MakeTxResponseJson>;
  createSession: (
    creatorAddress: string,
    sessionKey: string,
    expiresAt: string,
    allowPaths: string[],
    spendLimit: CoinJson[],
    spendPeriod: string,
    gasFee: string,
    gasWanted: string,
    memo?: string,
  ) => Promise<AsyncIterable<CreateSessionResponseJson>>;
  makeCreateSessionTx: (
    creatorAddress: string,
    sessionKey: string,
    expiresAt: string,
    allowPaths: string[],
    spendLimit: CoinJson[],
    spendPeriod: string,
    gasFee: string,
    gasWanted: string,
    memo?: string,
  ) => Promise<MakeTxResponseJson>;
  revokeSession: (
    creatorAddress: string,
    sessionKey: string,
    gasFee: string,
    gasWanted: string,
    memo?: string,
  ) => Promise<AsyncIterable<RevokeSessionResponseJson>>;
  makeRevokeSessionTx: (
    creatorAddress: string,
    sessionKey: string,
    gasFee: string,
    gasWanted: string,
    memo?: string,
  ) => Promise<MakeTxResponseJson>;
  revokeAllSessions: (
    creatorAddress: string,
    gasFee: string,
    gasWanted: string,
    memo?: string,
  ) => Promise<AsyncIterable<RevokeAllSessionsResponseJson>>;
  makeRevokeAllSessionsTx: (
    creatorAddress: string,
    gasFee: string,
    gasWanted: string,
    memo?: string,
  ) => Promise<MakeTxResponseJson>;
  broadcastTxCommit: (signedTxJson: string) => Promise<AsyncIterable<BroadcastTxCommitResponseJson>>;
  // debug
  hello: (name: string) => Promise<string>;
  helloStream: (name: string) => Promise<AsyncIterable<HelloStreamResponseJson>>;
}
