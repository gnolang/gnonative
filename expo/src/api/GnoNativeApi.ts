import { Client } from '@connectrpc/connect';

import { GnoKeyApi, BridgeStatus, Config } from './types';
import {
  BroadcastTxCommitResponse,
  CallResponse,
  DeleteAccountResponse,
  GetActivatedAccountResponse,
  HelloStreamResponse,
  KeyInfo,
  MakeTxResponse,
  Coin,
  QueryAccountResponse,
  QueryResponse,
  ActivateAccountResponse,
  SendResponse,
  RunResponse,
  SetChainIDResponse,
  SetPasswordResponse,
  SetRemoteResponse,
  SignTxResponse,
  RotatePasswordResponse,
  EstimateGasResponse,
} from './vendor/gnonativetypes_pb';
import { GnoNativeService } from './vendor/rpc_pb';
import { GoBridge, GoBridgeInterface } from '../GoBridge';
import * as Grpc from '../grpc/client';

export class GnoNativeApi implements GnoKeyApi, GoBridgeInterface {
  bridgeStatus = BridgeStatus.Stopped;
  config: Config;
  clientInstance: Client<typeof GnoNativeService> | undefined;

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
      await this.clientInstance?.setRemote({ remote: this.config.remote });
      await this.clientInstance?.setChainID({ chainId: this.config.chain_id });
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
    const response = client.setRemote({ remote });
    return response;
  }

  async getRemote(): Promise<string> {
    const client = this.#getClient();
    const response = await client.getRemote({});
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
    const response = client.estimateGas({
      txJson,
      address,
      securityMargin,
      updateTx,
      accountNumber,
      sequenceNumber,
    });
    return response;
  }

  async setChainID(chainId: string): Promise<SetChainIDResponse> {
    const client = this.#getClient();
    const response = client.setChainID({ chainId });
    return response;
  }

  async getChainID() {
    const client = this.#getClient();
    const response = await client.getChainID({});
    return response.chainId;
  }

  async createAccount(
    nameOrBech32: string,
    mnemonic: string,
    password: string,
  ): Promise<KeyInfo | undefined> {
    const client = this.#getClient();
    const reponse = await client.createAccount({
      nameOrBech32,
      mnemonic,
      password,
    });
    return reponse.key;
  }

  async createLedger(
    name: string,
    algorithm: string,
    hrp: string,
  ): Promise<KeyInfo | undefined> {
    const client = this.#getClient();
    const reponse = await client.createLedger({
      name,
      algorithm,
      hrp,
    });
    return reponse.key;
  }

  async generateRecoveryPhrase() {
    const client = this.#getClient();
    const response = await client.generateRecoveryPhrase({});
    return response.phrase;
  }

  async hasKeyByName(name: string) {
    const client = this.#getClient();
    const response = await client.hasKeyByName({ name });
    return response.has;
  }

  async hasKeyByAddress(address: Uint8Array) {
    const client = this.#getClient();
    const response = await client.hasKeyByAddress({ address });
    return response.has;
  }

  async hasKeyByNameOrAddress(nameOrBech32: string) {
    const client = this.#getClient();
    const response = await client.hasKeyByNameOrAddress({ nameOrBech32 });
    return response.has;
  }

  async getKeyInfoByName(name: string): Promise<KeyInfo | undefined> {
    const client = this.#getClient();
    const response = await client.getKeyInfoByName({ name });
    return response.key;
  }

  async getKeyInfoByAddress(address: Uint8Array): Promise<KeyInfo | undefined> {
    const client = this.#getClient();
    const response = await client.getKeyInfoByAddress({ address });
    return response.key;
  }

  async getKeyInfoByNameOrAddress(nameOrBech32: string): Promise<KeyInfo | undefined> {
    const client = this.#getClient();
    const response = await client.getKeyInfoByNameOrAddress({ nameOrBech32 });
    return response.key;
  }

  async listKeyInfo(): Promise<KeyInfo[]> {
    const client = this.#getClient();
    const response = await client.listKeyInfo({});
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
    maxDeposit?: Coin[],
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
          maxDeposit,
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

  async makeRunTx(
    pkg: string,
    gasFee: string,
    gasWanted: bigint,
    callerAddress: Uint8Array,
    send?: Coin[],
    maxDeposit?: Coin[],
    memo?: string,
  ): Promise<MakeTxResponse> {
    const client = this.#getClient();
    const reponse = client.makeRunTx({
      gasFee,
      gasWanted,
      memo,
      callerAddress,
      msgs: [
        {
          package: pkg,
          send,
          maxDeposit,
        },
      ],
    });
    return reponse;
  }

  async activateAccount(nameOrBech32: string): Promise<ActivateAccountResponse> {
    const client = this.#getClient();
    const response = client.activateAccount({ nameOrBech32 });
    return response;
  }

  async getClient(): Promise<Client<typeof GnoNativeService>> {
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
    const response = client.setPassword({ password, address });
    return response;
  }

  async rotatePassword(
    newPassword: string,
    addresses: Uint8Array[],
  ): Promise<RotatePasswordResponse> {
    const client = this.#getClient();
    const response = client.rotatePassword({ newPassword, addresses });
    return response;
  }

  async getActivatedAccount(): Promise<GetActivatedAccountResponse> {
    const client = this.#getClient();
    const response = client.getActivatedAccount({});
    return response;
  }

  async queryAccount(address: Uint8Array): Promise<QueryAccountResponse> {
    const client = this.#getClient();
    const reponse = client.queryAccount({ address });
    return reponse;
  }

  async deleteAccount(
    nameOrBech32: string,
    password: string | undefined,
    skipPassword: boolean,
  ): Promise<DeleteAccountResponse> {
    const client = this.#getClient();
    const response = client.deleteAccount({ nameOrBech32, password, skipPassword });
    return response;
  }

  async query(path: string, data: Uint8Array): Promise<QueryResponse> {
    const client = this.#getClient();
    const reponse = client.query({ path, data });
    return reponse;
  }

  async render(packagePath: string, args: string) {
    const client = this.#getClient();
    const reponse = await client.render({ packagePath, args });
    return reponse.result;
  }

  async qEval(packagePath: string, expression: string) {
    const client = this.#getClient();
    const reponse = await client.qEval({ packagePath, expression });
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
    maxDeposit?: Coin[],
    memo?: string,
  ): Promise<AsyncIterable<CallResponse>> {
    const client = this.#getClient();
    const reponse = client.call({
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
          maxDeposit,
        },
      ],
    });
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
    const reponse = client.send({
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

  async run(
    pkg: string,
    gasFee: string,
    gasWanted: bigint,
    callerAddress: Uint8Array,
    send?: Coin[],
    maxDeposit?: Coin[],
    memo?: string,
  ): Promise<AsyncIterable<RunResponse>> {
    const client = this.#getClient();
    const reponse = client.run({
      gasFee,
      gasWanted,
      callerAddress,
      memo,
      msgs: [
        {
          package: pkg,
          send,
          maxDeposit,
        },
      ],
    });
    return reponse;
  }

  async addressToBech32(address: Uint8Array) {
    const client = this.#getClient();
    const response = await client.addressToBech32({ address });
    return response.bech32Address;
  }

  async addressFromMnemonic(mnemonic: string) {
    const client = this.#getClient();
    const response = await client.addressFromMnemonic({ mnemonic });
    return response.address;
  }

  async addressFromBech32(bech32Address: string) {
    const client = this.#getClient();
    const response = await client.addressFromBech32({ bech32Address });
    return response.address;
  }

  async validateMnemonicWord(word: string) {
    const client = this.#getClient();
    const response = await client.validateMnemonicWord({ word });
    return response.valid;
  }

  async validateMnemonicPhrase(phrase: string) {
    const client = this.#getClient();
    const response = await client.validateMnemonicPhrase({ phrase });
    return response.valid;
  }

  async broadcastTxCommit(signedTxJson: string): Promise<AsyncIterable<BroadcastTxCommitResponse>> {
    const client = this.#getClient();
    const response = client.broadcastTxCommit({ signedTxJson });
    return response;
  }

  // // debug
  async hello(name: string) {
    const client = this.#getClient();
    const response = await client.hello({ name });
    return response.greeting;
  }

  // // debug
  async helloStream(name: string): Promise<AsyncIterable<HelloStreamResponse>> {
    const client = this.#getClient();
    return client.helloStream({ name });
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
