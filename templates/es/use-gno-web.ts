import {
  KeyInfo,
  SetRemoteRequest,
  SetRemoteResponse,
} from "./gnonativetypes_pb";
import { GetRemoteRequest } from "./gnonativetypes_pb";
import { SetChainIDRequest, SetChainIDResponse } from "./gnonativetypes_pb";
import { GetChainIDRequest } from "./gnonativetypes_pb";
import { SetPasswordRequest, SetPasswordResponse } from "./gnonativetypes_pb";
import { SelectAccountRequest } from "./gnonativetypes_pb";
import { SelectAccountResponse } from "./gnonativetypes_pb";
import { CreateAccountRequest } from "./gnonativetypes_pb";
import { GenerateRecoveryPhraseRequest } from "./gnonativetypes_pb";
import { ListKeyInfoRequest } from "./gnonativetypes_pb";
import { HasKeyByNameRequest } from "./gnonativetypes_pb";
import { HasKeyByAddressRequest } from "./gnonativetypes_pb";
import { HasKeyByNameOrAddressRequest } from "./gnonativetypes_pb";
import { GetKeyInfoByNameRequest } from "./gnonativetypes_pb";
import { GetKeyInfoByAddressRequest } from "./gnonativetypes_pb";
import { GetKeyInfoByNameOrAddressRequest } from "./gnonativetypes_pb";
import { GetActiveAccountRequest } from "./gnonativetypes_pb";
import { GetActiveAccountResponse } from "./gnonativetypes_pb";
import { QueryAccountRequest } from "./gnonativetypes_pb";
import { QueryAccountResponse } from "./gnonativetypes_pb";
import {
  DeleteAccountRequest,
  DeleteAccountResponse,
} from "./gnonativetypes_pb";
import { QueryRequest } from "./gnonativetypes_pb";
import { QueryResponse } from "./gnonativetypes_pb";
import { RenderRequest } from "./gnonativetypes_pb";
import { QEvalRequest } from "./gnonativetypes_pb";
import { CallRequest } from "./gnonativetypes_pb";
import { CallResponse } from "./gnonativetypes_pb";
import { AddressToBech32Request } from "./gnonativetypes_pb";
import { AddressFromBech32Request } from "./gnonativetypes_pb";
import { PromiseClient } from "@connectrpc/connect";
import { GnoNativeService } from "./rpc_connect";
import { createPromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";

export type GnoAccount = KeyInfo;

export interface GnoResponse {
  setRemote: (remote: string) => Promise<SetRemoteResponse>;
  getRemote: () => Promise<string>;
  setChainID: (chainId: string) => Promise<SetChainIDResponse>;
  getChainID: () => Promise<string>;
  createAccount: (
    nameOrBech32: string,
    mnemonic: string,
    password: string,
  ) => Promise<GnoAccount | undefined>;
  generateRecoveryPhrase: () => Promise<string>;
  listKeyInfo: () => Promise<GnoAccount[]>;
  hasKeyByName: (name: string) => Promise<boolean>;
  hasKeyByAddress: (address: Uint8Array) => Promise<boolean>;
  hasKeyByNameOrAddress: (nameOrBech32: string) => Promise<boolean>;
  getKeyInfoByName: (name: string) => Promise<GnoAccount | undefined>;
  getKeyInfoByAddress: (address: Uint8Array) => Promise<GnoAccount | undefined>;
  getKeyInfoByNameOrAddress: (
    nameOrBech32: string,
  ) => Promise<GnoAccount | undefined>;
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
    args: Array<string>,
    gasFee: string,
    gasWanted: number,
    send?: string,
    memo?: string,
  ) => Promise<AsyncIterable<CallResponse>>;
  addressToBech32: (address: Uint8Array) => Promise<string>;
  addressFromBech32: (bech32Address: string) => Promise<Uint8Array>;
  closeBridge: () => Promise<void>;
  initBridge: () => Promise<void>;
}

let clientInstance: PromiseClient<typeof GnoNativeService> | undefined =
  undefined;
let bridgeInstance = false;

export const useGno = (): GnoResponse => {
  const getClient = async () => {
    if (!bridgeInstance) {
      await initBridge();
    }

    if (clientInstance) return clientInstance;

    console.log("Creating GRPC client instance...");

    const port = 26658;
    const transport = createConnectTransport({
      baseUrl: "http://127.0.0.1:" + port,
    });

    clientInstance = createPromiseClient(GnoNativeService, transport);

    console.log("Creating GRPC client instance... done.");

    // Set the initial configuration where it's different from the default.
    await clientInstance.setRemote(
      new SetRemoteRequest({ remote: "testnet.gno.berty.io:26657" }),
    );

    return clientInstance;
  };

  const closeBridge = async () => {
    if (bridgeInstance) {
      console.log("Closing bridge...");
      //   await GoBridge.closeBridge();
      console.log("Bridge closed.");
      bridgeInstance = false;
    }
  };

  const initBridge = async () => {
    if (!bridgeInstance) {
      console.log("Initializing bridge...");
      //   await GoBridge.initBridge();
      console.log("Bridge initialized.");
      bridgeInstance = true;
    }
  };

  const setRemote = async (remote: string) => {
    const client = await getClient();
    const response = await client.setRemote(new SetRemoteRequest({ remote }));
    return response;
  };

  const getRemote = async () => {
    const client = await getClient();
    const response = await client.getRemote(new GetRemoteRequest());
    return response.remote;
  };

  const setChainID = async (chainId: string) => {
    const client = await getClient();
    const response = await client.setChainID(
      new SetChainIDRequest({ chainId }),
    );
    return response;
  };

  const getChainID = async () => {
    const client = await getClient();
    const response = await client.getChainID(new GetChainIDRequest());
    return response.chainId;
  };

  const createAccount = async (
    nameOrBech32: string,
    mnemonic: string,
    password: string,
  ) => {
    const client = await getClient();
    const reponse = await client.createAccount(
      new CreateAccountRequest({
        nameOrBech32,
        mnemonic,
        password,
      }),
    );
    return reponse.key;
  };

  const generateRecoveryPhrase = async () => {
    const client = await getClient();
    const response = await client.generateRecoveryPhrase(
      new GenerateRecoveryPhraseRequest(),
    );
    return response.phrase;
  };

  const hasKeyByName = async (name: string) => {
    const client = await getClient();
    const response = await client.hasKeyByName(
      new HasKeyByNameRequest({ name }),
    );
    return response.has;
  };

  const hasKeyByAddress = async (address: Uint8Array) => {
    const client = await getClient();
    const response = await client.hasKeyByAddress(
      new HasKeyByAddressRequest({ address }),
    );
    return response.has;
  };

  const hasKeyByNameOrAddress = async (nameOrBech32: string) => {
    const client = await getClient();
    const response = await client.hasKeyByNameOrAddress(
      new HasKeyByNameOrAddressRequest({ nameOrBech32 }),
    );
    return response.has;
  };

  const getKeyInfoByName = async (name: string) => {
    const client = await getClient();
    const response = await client.getKeyInfoByName(
      new GetKeyInfoByNameRequest({ name }),
    );
    return response.key;
  };

  const getKeyInfoByAddress = async (address: Uint8Array) => {
    const client = await getClient();
    const response = await client.getKeyInfoByAddress(
      new GetKeyInfoByAddressRequest({ address }),
    );
    return response.key;
  };

  const getKeyInfoByNameOrAddress = async (nameOrBech32: string) => {
    const client = await getClient();
    const response = await client.getKeyInfoByNameOrAddress(
      new GetKeyInfoByNameOrAddressRequest({ nameOrBech32 }),
    );
    return response.key;
  };

  const listKeyInfo = async () => {
    const client = await getClient();
    const response = await client.listKeyInfo(new ListKeyInfoRequest());
    return response.keys;
  };

  const selectAccount = async (nameOrBech32: string) => {
    const client = await getClient();
    const response = await client.selectAccount(
      new SelectAccountRequest({
        nameOrBech32,
      }),
    );
    return response;
  };

  const setPassword = async (password: string) => {
    const client = await getClient();
    const response = await client.setPassword(
      new SetPasswordRequest({ password }),
    );
    return response;
  };

  const getActiveAccount = async () => {
    const client = await getClient();
    const response = await client.getActiveAccount(
      new GetActiveAccountRequest(),
    );
    return response;
  };

  const queryAccount = async (address: Uint8Array) => {
    const client = await getClient();
    const reponse = await client.queryAccount(
      new QueryAccountRequest({ address }),
    );
    return reponse;
  };

  const deleteAccount = async (
    nameOrBech32: string,
    password: string | undefined,
    skipPassword: boolean,
  ) => {
    const client = await getClient();
    const response = await client.deleteAccount(
      new DeleteAccountRequest({
        nameOrBech32,
        password,
        skipPassword,
      }),
    );
    return response;
  };

  const query = async (path: string, data: Uint8Array) => {
    const client = await getClient();
    const reponse = await client.query(
      new QueryRequest({
        path,
        data,
      }),
    );
    return reponse;
  };

  const render = async (packagePath: string, args: string) => {
    const client = await getClient();
    const reponse = await client.render(
      new RenderRequest({
        packagePath,
        args,
      }),
    );
    return reponse.result;
  };

  const qEval = async (packagePath: string, expression: string) => {
    const client = await getClient();
    const reponse = await client.qEval(
      new QEvalRequest({
        packagePath,
        expression,
      }),
    );
    return reponse.result;
  };

  const call = async (
    packagePath: string,
    fnc: string,
    args: Array<string>,
    gasFee: string,
    gasWanted: number,
    send?: string,
    memo?: string,
  ) => {
    const client = await getClient();
    const reponse = client.call(
      new CallRequest({
        packagePath,
        fnc,
        args,
        gasFee,
        gasWanted: BigInt(gasWanted),
        send,
        memo,
      }),
    );
    return reponse;
  };

  const addressToBech32 = async (address: Uint8Array) => {
    const client = await getClient();
    const response = await client.addressToBech32(
      new AddressToBech32Request({ address }),
    );
    return response.bech32Address;
  };

  const addressFromBech32 = async (bech32Address: string) => {
    const client = await getClient();
    const response = await client.addressFromBech32(
      new AddressFromBech32Request({ bech32Address }),
    );
    return response.address;
  };

  return {
    initBridge,
    closeBridge,
    setRemote,
    getRemote,
    setChainID,
    getChainID,
    createAccount,
    generateRecoveryPhrase,
    listKeyInfo,
    hasKeyByName,
    hasKeyByAddress,
    hasKeyByNameOrAddress,
    getKeyInfoByName,
    getKeyInfoByAddress,
    getKeyInfoByNameOrAddress,
    selectAccount,
    setPassword,
    getActiveAccount,
    queryAccount,
    deleteAccount,
    query,
    render,
    qEval,
    call,
    addressToBech32,
    addressFromBech32,
  };
};
