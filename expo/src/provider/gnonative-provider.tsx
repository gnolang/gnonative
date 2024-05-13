import {
  AddressFromBech32Request,
  AddressToBech32Request,
  CallRequest,
  CallResponse,
  CreateAccountRequest,
  DeleteAccountRequest,
  DeleteAccountResponse,
  GenerateRecoveryPhraseRequest,
  GetActiveAccountRequest,
  GetActiveAccountResponse,
  GetChainIDRequest,
  GetKeyInfoByAddressRequest,
  GetKeyInfoByNameOrAddressRequest,
  GetKeyInfoByNameRequest,
  GetRemoteRequest,
  HasKeyByAddressRequest,
  HasKeyByNameOrAddressRequest,
  HasKeyByNameRequest,
  HelloRequest,
  HelloStreamResponse,
  ListKeyInfoRequest,
  MsgCall,
  MsgSend,
  QEvalRequest,
  QueryAccountRequest,
  QueryAccountResponse,
  QueryRequest,
  QueryResponse,
  RenderRequest,
  SelectAccountRequest,
  SelectAccountResponse,
  SendRequest,
  SendResponse,
  SetChainIDRequest,
  SetChainIDResponse,
  SetPasswordRequest,
  SetPasswordResponse,
  SetRemoteRequest,
  SetRemoteResponse,
} from '@buf/gnolang_gnonative.bufbuild_es/gnonativetypes_pb';
import { GnoNativeService } from '@buf/gnolang_gnonative.connectrpc_es/rpc_connect';
import { PromiseClient } from '@connectrpc/connect';
import { createContext, useContext, useEffect, useState } from 'react';

import { GoBridge } from '../GoBridge';
import * as Grpc from '../grpc/client';
import { GnoAccount } from '../hooks/types';

export interface GnoNativeContextProps {
  init: (config: ConfigProps) => Promise<boolean>;

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
  getKeyInfoByNameOrAddress: (nameOrBech32: string) => Promise<GnoAccount | undefined>;
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
  // closeBridge: () => Promise<void>;
  // initBridge: () => Promise<void>;
  // debug
  hello: (name: string) => Promise<string>;
  helloStream: (name: string) => Promise<AsyncIterable<HelloStreamResponse>>;
}

interface ConfigProps {
  remote: string;
  chain_id: string;
}

interface GnoNativeProviderProps {
  config: ConfigProps;
  children: React.ReactNode;
}

enum BridgeStatus {
  Stopped,
  Starting,
  Started,
}

const GnoNativeContext = createContext<GnoNativeContextProps | null>(null);

const GnoNativeProvider: React.FC<GnoNativeProviderProps> = ({ children, config }) => {
  const [initialized, setInitialized] = useState(false);
  const [clientInstance, setClientInstance] = useState<
    PromiseClient<typeof GnoNativeService> | undefined
  >(undefined);
  const [bridgeStatus, setBridgeStatus] = useState(BridgeStatus.Stopped);

  useEffect(() => {
    (async () => {
      await init(config);
      setInitialized(true);
    })();
  }, []);

  async function init(config): Promise<boolean> {
    console.log(
      '🍄 Initializing GnoNative on remote: %s chain_id: %s',
      config.remote,
      config.chain_id,
    );

    if (bridgeStatus === BridgeStatus.Stopped) {
      console.log('Bridge stopped. Initializing...');
      await initBridge();
    }

    if (clientInstance) {
      console.error('GoBridge already initialized.');
      return true;
    }

    const port = await GoBridge.getTcpPort();
    console.log('GoBridge GRPC client instance port: %s', port);
    const client = await Grpc.createClient(port);
    setClientInstance(client);
    console.log('GoBridge GRPC client instance. Done.');

    try {
      await client.setRemote(new SetRemoteRequest({ remote: config.remote }));
      await client.setChainID(new SetChainIDRequest({ chainId: config.chain_id }));

      console.log('✅ GnoNative bridge initialized.');
    } catch (error) {
      console.error(error);
      return false;
    }

    return true;
  }

  const initBridge = async () => {
    if (bridgeStatus === BridgeStatus.Stopped) {
      console.log('Initializing bridge...');
      setBridgeStatus(BridgeStatus.Starting);
      await GoBridge.initBridge();
      console.log('Bridge initialized.');
      setBridgeStatus(BridgeStatus.Started);
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
    const response = await client.hasKeyByNameOrAddress(
      new HasKeyByNameOrAddressRequest({ nameOrBech32 }),
    );
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

  const getClient = async () => {
    if (!clientInstance) {
      throw new Error('GoBridge client instance not initialized.');
    }

    return clientInstance;
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
    args: string[],
    gasFee: string,
    gasWanted: number,
    send?: string,
    memo?: string,
  ) => {
    const client = await getClient();
    const reponse = client.call(
      new CallRequest({
        gasFee,
        gasWanted: BigInt(gasWanted),
        memo,
        msgs: [
          new MsgCall({
            packagePath,
            fnc,
            args,
            send,
          }),
        ],
      }),
    );
    return reponse;
  };

  const send = async (
    toAddress: Uint8Array,
    send: string,
    gasFee: string,
    gasWanted: number,
    memo?: string,
  ) => {
    const client = await getClient();
    const reponse = client.send(
      new SendRequest({
        gasFee,
        gasWanted: BigInt(gasWanted),
        memo,
        msgs: [
          new MsgSend({
            toAddress,
            send,
          }),
        ],
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
    const response = await client.addressFromBech32(
      new AddressFromBech32Request({ bech32Address }),
    );
    return response.address;
  };

  // debug
  const hello = async (name: string) => {
    const client = await getClient();
    const response = await client.hello(new HelloRequest({ name }));
    return response.greeting;
  };

  // debug
  const helloStream = async (name: string) => {
    const client = await getClient();
    return client.helloStream(new HelloRequest({ name }));
  };

  const value = {
    init,
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
    send,
    addressToBech32,
    addressFromBech32,
    // debug
    hello,
    helloStream,
  };

  if (!initialized) {
    return null;
  }

  return <GnoNativeContext.Provider value={value}>{children}</GnoNativeContext.Provider>;
};

function useGnoNativeContext() {
  const context = useContext(GnoNativeContext) as GnoNativeContextProps;

  if (context === undefined) {
    throw new Error('useGnoNativeContext must be used within a GnoNativeProvider');
  }
  return context;
}

export { useGnoNativeContext, GnoNativeProvider };
