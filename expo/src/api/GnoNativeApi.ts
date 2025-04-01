import { PromiseClient } from '@connectrpc/connect';

import { GnoKeyApi, BridgeStatus, Config } from './types';
import {
  AddressFromBech32Request,
  AddressToBech32Request,
  AddressFromMnemonicRequest,
  BroadcastTxCommitResponse,
  CallRequest,
  CallResponse,
  CreateAccountRequest,
  DeleteAccountRequest,
  DeleteAccountResponse,
  GenerateRecoveryPhraseRequest,
  GetActivatedAccountRequest,
  GetActivatedAccountResponse,
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
  KeyInfo,
  ListKeyInfoRequest,
  MakeTxResponse,
  MsgCall,
  Coin,
  MsgSend,
  QEvalRequest,
  QueryAccountRequest,
  QueryAccountResponse,
  QueryRequest,
  QueryResponse,
  RenderRequest,
  ActivateAccountRequest,
  ActivateAccountResponse,
  SendRequest,
  SendResponse,
  SetChainIDRequest,
  SetChainIDResponse,
  SetPasswordRequest,
  SetPasswordResponse,
  SetRemoteRequest,
  SetRemoteResponse,
  SignTxResponse,
  RotatePasswordRequest,
  RotatePasswordResponse,
  EstimateGasRequest,
  EstimateGasResponse,
} from './vendor/gnonativetypes_pb';
import { GnoNativeService } from './vendor/rpc_connect';
import { GoBridge, GoBridgeInterface } from '../GoBridge';
import * as Grpc from '../grpc/client';

export class GnoNativeApi implements GnoKeyApi, GoBridgeInterface {
  bridgeStatus = BridgeStatus.Stopped;
  config: Config;
  clientInstance: PromiseClient<typeof GnoNativeService> | undefined;

  constructor(config: Config) {
    this.config = config;
  }

  async initClient() {
    if (this.bridgeStatus === BridgeStatus.Stopped) {
      console.log('Bridge stopped. Initializing...');
      await this.#initBridge();
    }

    if (this.clientInstance) {
      console.error('GoBridge already initialized.');
      return true;
    }

    const port = await GoBridge.getTcpPort();
    console.log('GoBridge GRPC client instance port: %s', port);
    const client = Grpc.createClient(port);
    this.clientInstance = client;
    console.log('GoBridge GRPC client instance. Done.');

    try {
      await this.clientInstance?.setRemote(new SetRemoteRequest({ remote: this.config.remote }));
      await this.clientInstance?.setChainID(
        new SetChainIDRequest({ chainId: this.config.chain_id }),
      );
      console.log('✅ GnoNative bridge initialized.');
      return true;
    } catch (error) {
      console.error(error);
      return false;
    }
  }

  async #initBridge() {
    if (this.bridgeStatus === BridgeStatus.Stopped) {
      console.log('Initializing bridge...');
      this.bridgeStatus = BridgeStatus.Starting;
      await GoBridge.initBridge();
      console.log('Bridge initialized.');
      this.bridgeStatus = BridgeStatus.Started;
    }
  }

  #getClient() {
    if (!this.clientInstance) {
      throw new Error('GoBridge client instance not initialized.');
    }

    return this.clientInstance;
  }

  async setRemote(remote: string): Promise<SetRemoteResponse> {
    const client = this.#getClient();
    const response = client.setRemote(new SetRemoteRequest({ remote }));
    return response;
  }

  async getRemote(): Promise<string> {
    const client = this.#getClient();
    const response = await client.getRemote(new GetRemoteRequest());
    return response.remote;
  }

  async signTx(
    txJson: string,
    address: Uint8Array,
    accountNumber: bigint = BigInt(0),
    sequenceNumber: bigint = BigInt(0),
  ): Promise<SignTxResponse> {
    const client = this.#getClient();
    const response = client.signTx({ txJson, address, accountNumber, sequenceNumber });
    return response;
  }

  async estimateGas(
    txJson: string,
    address: Uint8Array,
    securityMargin: number,
    updateTx: boolean,
    accountNumber: bigint = BigInt(0),
    sequenceNumber: bigint = BigInt(0),
  ): Promise<EstimateGasResponse> {
    const client = this.#getClient();
    const response = client.estimateGas(
      new EstimateGasRequest({
        txJson,
        address,
        securityMargin,
        updateTx,
        accountNumber,
        sequenceNumber,
      }),
    );
    return response;
  }

  async setChainID(chainId: string): Promise<SetChainIDResponse> {
    const client = this.#getClient();
    const response = client.setChainID(new SetChainIDRequest({ chainId }));
    return response;
  }

  async getChainID() {
    const client = this.#getClient();
    const response = await client.getChainID(new GetChainIDRequest());
    return response.chainId;
  }

  async createAccount(
    nameOrBech32: string,
    mnemonic: string,
    password: string,
  ): Promise<KeyInfo | undefined> {
    const client = this.#getClient();
    const reponse = await client.createAccount(
      new CreateAccountRequest({
        nameOrBech32,
        mnemonic,
        password,
      }),
    );
    return reponse.key;
  }

  async generateRecoveryPhrase() {
    const client = this.#getClient();
    const response = await client.generateRecoveryPhrase(new GenerateRecoveryPhraseRequest());
    return response.phrase;
  }

  async hasKeyByName(name: string) {
    const client = this.#getClient();
    const response = await client.hasKeyByName(new HasKeyByNameRequest({ name }));
    return response.has;
  }

  async hasKeyByAddress(address: Uint8Array) {
    const client = this.#getClient();
    const response = await client.hasKeyByAddress(new HasKeyByAddressRequest({ address }));
    return response.has;
  }

  async hasKeyByNameOrAddress(nameOrBech32: string) {
    const client = this.#getClient();
    const response = await client.hasKeyByNameOrAddress(
      new HasKeyByNameOrAddressRequest({ nameOrBech32 }),
    );
    return response.has;
  }

  async getKeyInfoByName(name: string): Promise<KeyInfo | undefined> {
    const client = this.#getClient();
    const response = await client.getKeyInfoByName(new GetKeyInfoByNameRequest({ name }));
    return response.key;
  }

  async getKeyInfoByAddress(address: Uint8Array): Promise<KeyInfo | undefined> {
    const client = this.#getClient();
    const response = await client.getKeyInfoByAddress(new GetKeyInfoByAddressRequest({ address }));
    return response.key;
  }

  async getKeyInfoByNameOrAddress(nameOrBech32: string): Promise<KeyInfo | undefined> {
    const client = this.#getClient();
    const response = await client.getKeyInfoByNameOrAddress(
      new GetKeyInfoByNameOrAddressRequest({ nameOrBech32 }),
    );
    return response.key;
  }

  async listKeyInfo(): Promise<KeyInfo[]> {
    const client = this.#getClient();
    const response = await client.listKeyInfo(new ListKeyInfoRequest());
    return response.keys;
  }

  async makeCallTx(
    packagePath: string,
    fnc: string,
    args: string[],
    gasFee: string,
    gasWanted: bigint,
    callerAddress: Uint8Array,
    send?: Coin[],
    memo?: string,
  ): Promise<MakeTxResponse> {
    const client = this.#getClient();
    const reponse = client.makeCallTx({
      gasFee,
      gasWanted,
      memo,
      callerAddress,
      msgs: [
        {
          packagePath,
          fnc,
          args,
          send,
        },
      ],
    });
    return reponse;
  }

  async makeSendTx(
    toAddress: Uint8Array,
    amount: Coin[],
    gasFee: string,
    gasWanted: bigint,
    callerAddress: Uint8Array,
    memo?: string,
  ): Promise<MakeTxResponse> {
    const client = this.#getClient();
    const reponse = client.makeSendTx({
      gasFee,
      gasWanted,
      memo,
      callerAddress,
      msgs: [
        {
          toAddress,
          amount,
        },
      ],
    });
    return reponse;
  }

  async activateAccount(nameOrBech32: string): Promise<ActivateAccountResponse> {
    const client = this.#getClient();
    const response = client.activateAccount(
      new ActivateAccountRequest({
        nameOrBech32,
      }),
    );
    return response;
  }

  async getClient(): Promise<PromiseClient<typeof GnoNativeService>> {
    if (!this.clientInstance) {
      throw new Error('GoBridge client instance not initialized.');
    }
    return this.clientInstance;
  }

  isInitialized(): boolean {
    return this.clientInstance !== undefined;
  }

  async setPassword(password: string, address: Uint8Array): Promise<SetPasswordResponse> {
    const client = this.#getClient();
    const response = client.setPassword(new SetPasswordRequest({ password, address }));
    return response;
  }

  async rotatePassword(
    newPassword: string,
    addresses: Uint8Array[],
  ): Promise<RotatePasswordResponse> {
    const client = this.#getClient();
    const response = client.rotatePassword(new RotatePasswordRequest({ newPassword, addresses }));
    return response;
  }

  async getActivatedAccount(): Promise<GetActivatedAccountResponse> {
    const client = this.#getClient();
    const response = client.getActivatedAccount(new GetActivatedAccountRequest());
    return response;
  }

  async queryAccount(address: Uint8Array): Promise<QueryAccountResponse> {
    const client = this.#getClient();
    const reponse = client.queryAccount(new QueryAccountRequest({ address }));
    return reponse;
  }

  async deleteAccount(
    nameOrBech32: string,
    password: string | undefined,
    skipPassword: boolean,
  ): Promise<DeleteAccountResponse> {
    const client = this.#getClient();
    const response = client.deleteAccount(
      new DeleteAccountRequest({
        nameOrBech32,
        password,
        skipPassword,
      }),
    );
    return response;
  }

  async query(path: string, data: Uint8Array): Promise<QueryResponse> {
    const client = this.#getClient();
    const reponse = client.query(
      new QueryRequest({
        path,
        data,
      }),
    );
    return reponse;
  }

  async render(packagePath: string, args: string) {
    const client = this.#getClient();
    const reponse = await client.render(
      new RenderRequest({
        packagePath,
        args,
      }),
    );
    return reponse.result;
  }

  async qEval(packagePath: string, expression: string) {
    const client = this.#getClient();
    const reponse = await client.qEval(
      new QEvalRequest({
        packagePath,
        expression,
      }),
    );
    return reponse.result;
  }

  async call(
    packagePath: string,
    fnc: string,
    args: string[],
    gasFee: string,
    gasWanted: bigint,
    callerAddress: Uint8Array,
    send?: Coin[],
    memo?: string,
  ): Promise<AsyncIterable<CallResponse>> {
    const client = this.#getClient();
    const reponse = client.call(
      new CallRequest({
        gasFee,
        gasWanted,
        memo,
        callerAddress,
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
  }

  async send(
    toAddress: Uint8Array,
    amount: Coin[],
    gasFee: string,
    gasWanted: bigint,
    callerAddress: Uint8Array,
    memo?: string,
  ): Promise<AsyncIterable<SendResponse>> {
    const client = this.#getClient();
    const reponse = client.send(
      new SendRequest({
        gasFee,
        gasWanted,
        memo,
        callerAddress,
        msgs: [
          new MsgSend({
            toAddress,
            amount,
          }),
        ],
      }),
    );
    return reponse;
  }

  async addressToBech32(address: Uint8Array) {
    const client = this.#getClient();
    const response = await client.addressToBech32(new AddressToBech32Request({ address }));
    return response.bech32Address;
  }

  async addressFromMnemonic(mnemonic: string) {
    const client = this.#getClient();
    const response = await client.addressFromMnemonic(new AddressFromMnemonicRequest({ mnemonic }));
    return response.address;
  }

  async addressFromBech32(bech32Address: string) {
    const client = this.#getClient();
    const response = await client.addressFromBech32(
      new AddressFromBech32Request({ bech32Address }),
    );
    return response.address;
  }

  async broadcastTxCommit(signedTxJson: string): Promise<AsyncIterable<BroadcastTxCommitResponse>> {
    const client = this.#getClient();
    const response = client.broadcastTxCommit({ signedTxJson });
    return response;
  }

  // // debug
  async hello(name: string) {
    const client = this.#getClient();
    const response = await client.hello(new HelloRequest({ name }));
    return response.greeting;
  }

  // // debug
  async helloStream(name: string): Promise<AsyncIterable<HelloStreamResponse>> {
    const client = this.#getClient();
    return client.helloStream(new HelloRequest({ name }));
  }

  // Go Bridge Interface:
  initBridge(): Promise<void> {
    return GoBridge.initBridge();
  }
  closeBridge(): Promise<void> {
    return GoBridge.closeBridge();
  }
  getTcpPort(): Promise<number> {
    return GoBridge.getTcpPort();
  }
  invokeGrpcMethod(method: string, jsonMessage: string): Promise<string> {
    return GoBridge.invokeGrpcMethod(method, jsonMessage);
  }
  createStreamClient(method: string, jsonMessage: string): Promise<string> {
    return GoBridge.createStreamClient(method, jsonMessage);
  }
  streamClientReceive(id: string): Promise<string> {
    return GoBridge.streamClientReceive(id);
  }
  closeStreamClient(id: string): Promise<void> {
    return GoBridge.closeStreamClient(id);
  }
}
