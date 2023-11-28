import { SetRemoteRequest, SetRemoteResponse } from '@gno/api/gnomobiletypes_pb';
import { GetRemoteRequest } from '@gno/api/gnomobiletypes_pb';
import { SetChainIDRequest, SetChainIDResponse } from '@gno/api/gnomobiletypes_pb';
import { GetChainIDRequest } from '@gno/api/gnomobiletypes_pb';
import { SetPasswordRequest, SetPasswordResponse } from '@gno/api/gnomobiletypes_pb';
import { SelectAccountRequest } from '@gno/api/gnomobiletypes_pb';
import { SelectAccountResponse } from '@gno/api/gnomobiletypes_pb';
import { CreateAccountRequest } from '@gno/api/gnomobiletypes_pb';
import { GenerateRecoveryPhraseRequest } from '@gno/api/gnomobiletypes_pb';
import { ListKeyInfoRequest } from '@gno/api/gnomobiletypes_pb';
import { HasKeyByNameRequest } from '@gno/api/gnomobiletypes_pb';
import { HasKeyByAddressRequest } from '@gno/api/gnomobiletypes_pb';
import { HasKeyByNameOrAddressRequest } from '@gno/api/gnomobiletypes_pb';
import { GetKeyInfoByNameRequest } from '@gno/api/gnomobiletypes_pb';
import { GetKeyInfoByAddressRequest } from '@gno/api/gnomobiletypes_pb';
import { GetKeyInfoByNameOrAddressRequest } from '@gno/api/gnomobiletypes_pb';
import { GetActiveAccountRequest } from '@gno/api/gnomobiletypes_pb';
import { GetActiveAccountResponse } from '@gno/api/gnomobiletypes_pb';
import { QueryAccountRequest } from '@gno/api/gnomobiletypes_pb';
import { QueryAccountResponse } from '@gno/api/gnomobiletypes_pb';
import { DeleteAccountRequest, DeleteAccountResponse } from '@gno/api/gnomobiletypes_pb';
import { QueryRequest } from '@gno/api/gnomobiletypes_pb';
import { QueryResponse } from '@gno/api/gnomobiletypes_pb';
import { RenderRequest } from '@gno/api/gnomobiletypes_pb';
import { QEvalRequest } from '@gno/api/gnomobiletypes_pb';
import { CallRequest } from '@gno/api/gnomobiletypes_pb';
import { CallResponse } from '@gno/api/gnomobiletypes_pb';
import { AddressToBech32Request } from '@gno/api/gnomobiletypes_pb';
import { AddressFromBech32Request } from '@gno/api/gnomobiletypes_pb';
import * as Grpc from '@gno/grpc/client';
import { GnoAccount } from '@gno/native_modules/types';
import { GoBridge } from '@gno/native_modules';
import { PromiseClient } from '@connectrpc/connect';
import { GnomobileService } from '@gno/api/rpc_connect';

interface GnoResponse {
  setRemote: (remote: string) => Promise<SetRemoteResponse>;
  getRemote: () => Promise<string>;
  setChainID: (chainId: string) => Promise<SetChainIDResponse>;
  getChainID: () => Promise<string>;
  createAccount: (nameOrBech32: string, mnemonic: string, password: string) => Promise<GnoAccount | undefined>;
  generateRecoveryPhrase: () => Promise<string>;
  listKeyInfo: () => Promise<GnoAccount[]>;
  hasKeyByName: (name: string) => Promise<boolean>;
  hasKeyByAddress: (address: Uint8Array) => Promise<boolean>;
  hasKeyByNameOrAddress: (nameOrBech32: string) => Promise<boolean>;
  getKeyInfoByName: (name: string) => Promise<GnoAccount | undefined>;
  getKeyInfoByAddress: (address: Uint8Array) => Promise<GnoAccount | undefined>;
  getKeyInfoByNameOrAddress: (nameOrBech32: string) => Promise<GnoAccount | undefined>;
  selectAccount: (nameOrBech32: string) => Promise<SelectAccountResponse>;
  setPassword: (password: string) => Promise<SetPasswordResponse>;
  getActiveAccount: () => Promise<GetActiveAccountResponse>;
  queryAccount: (address: Uint8Array) => Promise<QueryAccountResponse>;
  deleteAccount: (nameOrBech32: string, password: string | undefined, skipPassword: boolean) => Promise<DeleteAccountResponse>;
  query: (path: string, data: Uint8Array) => Promise<QueryResponse>;
  render: (packagePath: string, args: string) => Promise<string>;
  qEval: (packagePath: string, expression: string) => Promise<string>;
  call: (packagePath: string, fnc: string, args: Array<string>, gasFee: string, gasWanted: number, send?: string, memo?: string) => Promise<CallResponse>;
  addressToBech32: (address: Uint8Array) => Promise<string>;
  addressFromBech32: (bech32Address: string) => Promise<Uint8Array>;
}

let clientInstance: PromiseClient<typeof GnomobileService> | undefined = undefined;
let bridgeInstance: boolean = false;

export const useGno = (): GnoResponse => {
  const getClient = async () => {
    if (!bridgeInstance) {
      console.log('Initializing bridge...');
      await GoBridge.initBridge();
      bridgeInstance = true;
    }

    if (clientInstance) return clientInstance;

    console.log('Creating GRPC client instance...');

    const port = await GoBridge.getTcpPort();
    clientInstance = Grpc.createClient(port);
    return clientInstance;
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
    const response = await client.setChainID(new SetChainIDRequest({ chainId }));
    return response;
  };

  const getChainID = async () => {
    const client = await getClient();
    const response = await client.getChainID(new GetChainIDRequest());
    return response.chainId;
  };

  const createAccount = async (nameOrBech32: string, mnemonic: string, password: string) => {
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
    const response = await client.generateRecoveryPhrase(new GenerateRecoveryPhraseRequest());
    return response.phrase;
  };

  const hasKeyByName = async (name: string) => {
    const client = await getClient();
    const response = await client.hasKeyByName(new HasKeyByNameRequest({ name }));
    return response.has;
  };

  const hasKeyByAddress = async (address: Uint8Array) => {
    const client = await getClient();
    const response = await client.hasKeyByAddress(new HasKeyByAddressRequest({ address }));
    return response.has;
  };

  const hasKeyByNameOrAddress = async (nameOrBech32: string) => {
    const client = await getClient();
    const response = await client.hasKeyByNameOrAddress(new HasKeyByNameOrAddressRequest({ nameOrBech32 }));
    return response.has;
  };

  const getKeyInfoByName = async (name: string) => {
    const client = await getClient();
    const response = await client.getKeyInfoByName(new GetKeyInfoByNameRequest({ name }));
    return response.key;
  };

  const getKeyInfoByAddress = async (address: Uint8Array) => {
    const client = await getClient();
    const response = await client.getKeyInfoByAddress(new GetKeyInfoByAddressRequest({ address }));
    return response.key;
  };

  const getKeyInfoByNameOrAddress = async (nameOrBech32: string) => {
    const client = await getClient();
    const response = await client.getKeyInfoByNameOrAddress(new GetKeyInfoByNameOrAddressRequest({ nameOrBech32 }));
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
    const response = await client.setPassword(new SetPasswordRequest({ password }));
    return response;
  };

  const getActiveAccount = async () => {
    const client = await getClient();
    const response = await client.getActiveAccount(new GetActiveAccountRequest());
    return response;
  };

  const queryAccount = async (address: Uint8Array) => {
    const client = await getClient();
    const reponse = await client.queryAccount(new QueryAccountRequest({ address }));
    return reponse;
  };

  const deleteAccount = async (nameOrBech32: string, password: string | undefined, skipPassword: boolean) => {
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

  const call = async (packagePath: string, fnc: string, args: Array<string>, gasFee: string, gasWanted: number, send?: string, memo?: string) => {
    const client = await getClient();
    const reponse = await client.call(
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
    const response = await client.addressToBech32(new AddressToBech32Request({ address }));
    return response.bech32Address;
  };

  const addressFromBech32 = async (bech32Address: string) => {
    const client = await getClient();
    const response = await client.addressFromBech32(new AddressFromBech32Request({ bech32Address }));
    return response.address;
  };

  return {
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
