import {
  CallResponse,
  DeleteAccountResponse,
  GetActiveAccountResponse,
  HelloStreamResponse,
  QueryAccountResponse,
  QueryResponse,
  SelectAccountResponse,
  SendResponse,
  SetChainIDResponse,
  SetPasswordResponse,
  SetRemoteResponse,
  KeyInfo,
} from '@buf/gnolang_gnonative.bufbuild_es/gnonativetypes_pb';

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
  ) => Promise<KeyInfo | undefined>;
  generateRecoveryPhrase: () => Promise<string>;
  listKeyInfo: () => Promise<KeyInfo[]>;
  hasKeyByName: (name: string) => Promise<boolean>;
  hasKeyByAddress: (address: Uint8Array) => Promise<boolean>;
  hasKeyByNameOrAddress: (nameOrBech32: string) => Promise<boolean>;
  getKeyInfoByName: (name: string) => Promise<KeyInfo | undefined>;
  getKeyInfoByAddress: (address: Uint8Array) => Promise<KeyInfo | undefined>;
  getKeyInfoByNameOrAddress: (nameOrBech32: string) => Promise<KeyInfo | undefined>;
  selectAccount: (nameOrBech32: string) => Promise<SelectAccountResponse>;
  setPassword: (password: string) => Promise<SetPasswordResponse>;
  getActiveAccount: () => Promise<GetActiveAccountResponse>;
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
    gasWanted: number,
    send?: string,
    memo?: string,
  ) => Promise<AsyncIterable<CallResponse>>;
  send: (
    toAddress: Uint8Array,
    send: string,
    gasFee: string,
    gasWanted: number,
    memo?: string,
  ) => Promise<AsyncIterable<SendResponse>>;
  addressToBech32: (address: Uint8Array) => Promise<string>;
  addressFromBech32: (bech32Address: string) => Promise<Uint8Array>;
  // debug
  hello: (name: string) => Promise<string>;
  helloStream: (name: string) => Promise<AsyncIterable<HelloStreamResponse>>;
}
