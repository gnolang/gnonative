import {
  CallResponse,
  DeleteAccountResponse,
  GetActivatedAccountResponse,
  HelloStreamResponse,
  QueryAccountResponse,
  QueryResponse,
  ActivateAccountResponse,
  SendResponse,
  RunResponse,
  SetChainIDResponse,
  SetPasswordResponse,
  SetRemoteResponse,
  RotatePasswordResponse,
  KeyInfo,
  SignTxResponse,
  EstimateGasResponse,
  MakeTxResponse,
  BroadcastTxCommitResponse,
  Coin,
} from './vendor/gnonativetypes_pb';

export enum BridgeStatus {
  Stopped,
  Starting,
  Started,
}

export interface Config {
  remote: string;
  chain_id: string;
}

export interface GnoKeyApi {
  setRemote: (remote: string) => Promise<SetRemoteResponse>;
  getRemote: () => Promise<string>;
  setChainID: (chainId: string) => Promise<SetChainIDResponse>;
  getChainID: () => Promise<string>;
  createAccount: (
    nameOrBech32: string,
    mnemonic: string,
    password: string,
    bip39Passwd?: string,
    account?: number,
    index?: number,
  ) => Promise<KeyInfo | undefined>;
  createLedger: (
    name: string,
    algorithm: string,
    hrp: string,
    account?: number,
    index?: number,
  ) => Promise<KeyInfo | undefined>;
  generateRecoveryPhrase: () => Promise<string>;
  listKeyInfo: () => Promise<KeyInfo[]>;
  hasKeyByName: (name: string) => Promise<boolean>;
  hasKeyByAddress: (address: Uint8Array) => Promise<boolean>;
  hasKeyByNameOrAddress: (nameOrBech32: string) => Promise<boolean>;
  getKeyInfoByName: (name: string) => Promise<KeyInfo | undefined>;
  getKeyInfoByAddress: (address: Uint8Array) => Promise<KeyInfo | undefined>;
  getKeyInfoByNameOrAddress: (nameOrBech32: string) => Promise<KeyInfo | undefined>;
  activateAccount: (nameOrBech32: string) => Promise<ActivateAccountResponse>;
  setPassword: (password: string, address: Uint8Array) => Promise<SetPasswordResponse>;
  rotatePassword: (password: string, addresses: Uint8Array[]) => Promise<RotatePasswordResponse>;
  getActivatedAccount: () => Promise<GetActivatedAccountResponse>;
  queryAccount: (address: Uint8Array) => Promise<QueryAccountResponse>;
  deleteAccount: (
    nameOrBech32: string,
    password: string | undefined,
    skipPassword: boolean,
  ) => Promise<DeleteAccountResponse>;
  query: (path: string, data: Uint8Array) => Promise<QueryResponse>;
  render: (packagePath: string, args: string) => Promise<string>;
  qEval: (packagePath: string, expression: string) => Promise<string>;
  call: (
    packagePath: string,
    fnc: string,
    args: string[],
    gasFee: string,
    gasWanted: bigint,
    callerAddress: Uint8Array,
    send?: Coin[],
    maxDeposit?: Coin[],
    memo?: string,
  ) => Promise<AsyncIterable<CallResponse>>;
  send: (
    toAddress: Uint8Array,
    amount: Coin[],
    gasFee: string,
    gasWanted: bigint,
    callerAddress: Uint8Array,
    memo?: string,
  ) => Promise<AsyncIterable<SendResponse>>;
  run: (
    pkg: string,
    gasFee: string,
    gasWanted: bigint,
    callerAddress: Uint8Array,
    send?: Coin[],
    maxDeposit?: Coin[],
    memo?: string,
  ) => Promise<AsyncIterable<RunResponse>>;
  addressToBech32: (address: Uint8Array) => Promise<string>;
  addressFromMnemonic: (mnemonic: string) => Promise<Uint8Array>;
  addressFromBech32: (bech32Address: string) => Promise<Uint8Array>;
  validateMnemonicWord: (word: string) => Promise<boolean>;
  validateMnemonicPhrase: (phrase: string) => Promise<boolean>;
  signTx(
    txJson: string,
    address: Uint8Array,
    accountNumber?: bigint,
    sequenceNumber?: bigint,
  ): Promise<SignTxResponse>;
  estimateGas(
    txJson: string,
    address: Uint8Array,
    securityMargin: number,
    updateTx: boolean,
    accountNumber?: bigint,
    sequenceNumber?: bigint,
  ): Promise<EstimateGasResponse>;
  makeCallTx(
    packagePath: string,
    fnc: string,
    args: string[],
    gasFee: string,
    gasWanted: bigint,
    callerAddress: Uint8Array,
    send?: Coin[],
    maxDeposit?: Coin[],
    memo?: string,
  ): Promise<MakeTxResponse>;
  makeSendTx(
    toAddress: Uint8Array,
    amount: Coin[],
    gasFee: string,
    gasWanted: bigint,
    callerAddress: Uint8Array,
    memo?: string,
  ): Promise<MakeTxResponse>;
  makeRunTx: (
    pkg: string,
    gasFee: string,
    gasWanted: bigint,
    callerAddress: Uint8Array,
    send?: Coin[],
    maxDeposit?: Coin[],
    memo?: string,
  ) => Promise<MakeTxResponse>;
  broadcastTxCommit(signedTxJson: string): Promise<AsyncIterable<BroadcastTxCommitResponse>>;
  // debug
  hello: (name: string) => Promise<string>;
  helloStream: (name: string) => Promise<AsyncIterable<HelloStreamResponse>>;
}
