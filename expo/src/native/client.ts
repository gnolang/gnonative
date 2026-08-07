// GnoNativeClient: the connect/gRPC-free client. It builds camelCase protojson request objects,
// sends them through GoBridge.invokeMethod/createStream, and parses the base64(protojson) responses.
// Method-for-method parity with the gRPC-path GnoNativeApi (src/api/GnoNativeApi.ts), so migration is
// an import swap plus adapting Uint8Array/bigint arguments to base64/string.
import { GoBridge } from '../GoBridge';
import { base64ToString } from './encoding';
import { GnoNativeError } from './error';
import { streamAsyncIterable } from './stream';
import type * as pb from './types.gen';
import { BridgeStatus, Config, GnoNativeClientApi } from './types';

export class GnoNativeClient implements GnoNativeClientApi {
  bridgeStatus = BridgeStatus.Stopped;
  config: Config;

  constructor(config: Config) {
    this.config = config;
  }

  async initClient(): Promise<boolean> {
    if (this.bridgeStatus === BridgeStatus.Stopped) {
      this.bridgeStatus = BridgeStatus.Starting;
      await GoBridge.initBridgeWithOptions({ useGrpcServers: false });
      this.bridgeStatus = BridgeStatus.Started;
    }

    await this.setRemote(this.config.remote);
    await this.setChainID(this.config.chain_id);
    console.log('✅ GnoNative (native) client initialized.');
    return true;
  }

  async closeBridge(): Promise<void> {
    await GoBridge.closeBridge();
    this.bridgeStatus = BridgeStatus.Stopped;
  }

  // #invoke sends a unary request and parses the protojson response.
  async #invoke<TRes>(method: string, req: object): Promise<TRes> {
    try {
      const res = await GoBridge.invokeMethod(method, JSON.stringify(req));
      return JSON.parse(base64ToString(res)) as TRes;
    } catch (e) {
      throw new GnoNativeError((e as Error)?.message ?? String(e));
    }
  }

  async setRemote(remote: string): Promise<pb.SetRemoteResponseJson> {
    const req: pb.SetRemoteRequestJson = { remote };
    return this.#invoke('SetRemote', req);
  }

  async getRemote(): Promise<string> {
    const res = await this.#invoke<pb.GetRemoteResponseJson>('GetRemote', {});
    return res.remote ?? '';
  }

  async setChainID(chainId: string): Promise<pb.SetChainIDResponseJson> {
    const req: pb.SetChainIDRequestJson = { chainId };
    return this.#invoke('SetChainID', req);
  }

  async getChainID(): Promise<string> {
    const res = await this.#invoke<pb.GetChainIDResponseJson>('GetChainID', {});
    return res.chainId ?? '';
  }

  async createAccount(
    nameOrBech32: string,
    mnemonic: string,
    password: string,
    bip39Passwd?: string,
    account?: number,
    index?: number,
  ): Promise<pb.KeyInfoJson | undefined> {
    const req: pb.CreateAccountRequestJson = {
      nameOrBech32,
      mnemonic,
      password,
      bip39Passwd,
      account,
      index,
    };
    const res = await this.#invoke<pb.CreateAccountResponseJson>('CreateAccount', req);
    return res.key_info;
  }

  async createLedger(
    name: string,
    algorithm: string,
    hrp: string,
    account?: number,
    index?: number,
  ): Promise<pb.KeyInfoJson | undefined> {
    const req: pb.CreateLedgerRequestJson = { name, algorithm, hrp, account, index };
    const res = await this.#invoke<pb.CreateLedgerResponseJson>('CreateLedger', req);
    return res.key_info;
  }

  async generateRecoveryPhrase(): Promise<string> {
    const res = await this.#invoke<pb.GenerateRecoveryPhraseResponseJson>(
      'GenerateRecoveryPhrase',
      {},
    );
    return res.phrase ?? '';
  }

  async listKeyInfo(): Promise<pb.KeyInfoJson[]> {
    const res = await this.#invoke<pb.ListKeyInfoResponseJson>('ListKeyInfo', {});
    return res.key_info ?? [];
  }

  async hasKeyByName(name: string): Promise<boolean> {
    const req: pb.HasKeyByNameRequestJson = { name };
    const res = await this.#invoke<pb.HasKeyByNameResponseJson>('HasKeyByName', req);
    return res.has ?? false;
  }

  async hasKeyByAddress(address: string): Promise<boolean> {
    const req: pb.HasKeyByAddressRequestJson = { address };
    const res = await this.#invoke<pb.HasKeyByAddressResponseJson>('HasKeyByAddress', req);
    return res.has ?? false;
  }

  async hasKeyByNameOrAddress(nameOrBech32: string): Promise<boolean> {
    const req: pb.HasKeyByNameOrAddressRequestJson = { nameOrBech32 };
    const res = await this.#invoke<pb.HasKeyByNameOrAddressResponseJson>(
      'HasKeyByNameOrAddress',
      req,
    );
    return res.has ?? false;
  }

  async getKeyInfoByName(name: string): Promise<pb.KeyInfoJson | undefined> {
    const req: pb.GetKeyInfoByNameRequestJson = { name };
    const res = await this.#invoke<pb.GetKeyInfoByNameResponseJson>('GetKeyInfoByName', req);
    return res.key_info;
  }

  async getKeyInfoByAddress(address: string): Promise<pb.KeyInfoJson | undefined> {
    const req: pb.GetKeyInfoByAddressRequestJson = { address };
    const res = await this.#invoke<pb.GetKeyInfoByAddressResponseJson>('GetKeyInfoByAddress', req);
    return res.key_info;
  }

  async getKeyInfoByNameOrAddress(nameOrBech32: string): Promise<pb.KeyInfoJson | undefined> {
    const req: pb.GetKeyInfoByNameOrAddressRequestJson = { nameOrBech32 };
    const res = await this.#invoke<pb.GetKeyInfoByNameOrAddressResponseJson>(
      'GetKeyInfoByNameOrAddress',
      req,
    );
    return res.key_info;
  }

  async activateAccount(nameOrBech32: string): Promise<pb.ActivateAccountResponseJson> {
    const req: pb.ActivateAccountRequestJson = { nameOrBech32 };
    return this.#invoke('ActivateAccount', req);
  }

  async setPassword(password: string, address: string): Promise<pb.SetPasswordResponseJson> {
    const req: pb.SetPasswordRequestJson = { password, address };
    return this.#invoke('SetPassword', req);
  }

  async renameKey(oldName: string, newName: string): Promise<pb.RenameKeyResponseJson> {
    const req: pb.RenameKeyRequestJson = { oldName, newName };
    return this.#invoke('RenameKey', req);
  }

  async rotatePassword(
    newPassword: string,
    addresses: string[],
  ): Promise<pb.RotatePasswordResponseJson> {
    const req: pb.RotatePasswordRequestJson = { newPassword, addresses };
    return this.#invoke('RotatePassword', req);
  }

  async getActivatedAccount(address: string): Promise<pb.GetActivatedAccountResponseJson> {
    const req: pb.GetActivatedAccountRequestJson = { address };
    return this.#invoke('GetActivatedAccount', req);
  }

  async queryAccount(address: string): Promise<pb.QueryAccountResponseJson> {
    const req: pb.QueryAccountRequestJson = { address };
    return this.#invoke('QueryAccount', req);
  }

  async querySessionAccount(
    masterAddress: string,
    sessionAddress: string,
  ): Promise<pb.QuerySessionAccountResponseJson> {
    const req: pb.QuerySessionAccountRequestJson = { masterAddress, sessionAddress };
    return this.#invoke('QuerySessionAccount', req);
  }

  async deleteAccount(
    nameOrBech32: string,
    password: string | undefined,
    skipPassword: boolean,
  ): Promise<pb.DeleteAccountResponseJson> {
    const req: pb.DeleteAccountRequestJson = { nameOrBech32, password, skipPassword };
    return this.#invoke('DeleteAccount', req);
  }

  async query(path: string, data: string): Promise<pb.QueryResponseJson> {
    const req: pb.QueryRequestJson = { path, data };
    return this.#invoke('Query', req);
  }

  async render(packagePath: string, args: string): Promise<string> {
    const req: pb.RenderRequestJson = { packagePath, args };
    const res = await this.#invoke<pb.RenderResponseJson>('Render', req);
    return res.result ?? '';
  }

  async qEval(packagePath: string, expression: string): Promise<string> {
    const req: pb.QEvalRequestJson = { packagePath, expression };
    const res = await this.#invoke<pb.QEvalResponseJson>('QEval', req);
    return res.result ?? '';
  }

  async call(
    packagePath: string,
    fnc: string,
    args: string[],
    gasFee: string,
    gasWanted: string,
    signerAddress: string,
    send?: pb.CoinJson[],
    maxDeposit?: pb.CoinJson[],
    memo?: string,
  ): Promise<AsyncIterable<pb.CallResponseJson>> {
    const req: pb.CallRequestJson = {
      gasFee,
      gasWanted,
      memo,
      signerAddress,
      Msgs: [{ packagePath, fnc, args, send, maxDeposit }],
    };
    return streamAsyncIterable<pb.CallResponseJson>('Call', JSON.stringify(req));
  }

  async send(
    toAddress: string,
    amount: pb.CoinJson[],
    gasFee: string,
    gasWanted: string,
    signerAddress: string,
    memo?: string,
  ): Promise<AsyncIterable<pb.SendResponseJson>> {
    const req: pb.SendRequestJson = {
      gasFee,
      gasWanted,
      memo,
      signerAddress,
      Msgs: [{ toAddress, amount }],
    };
    return streamAsyncIterable<pb.SendResponseJson>('Send', JSON.stringify(req));
  }

  async run(
    pkg: string,
    gasFee: string,
    gasWanted: string,
    signerAddress: string,
    send?: pb.CoinJson[],
    maxDeposit?: pb.CoinJson[],
    memo?: string,
  ): Promise<AsyncIterable<pb.RunResponseJson>> {
    const req: pb.RunRequestJson = {
      gasFee,
      gasWanted,
      memo,
      signerAddress,
      Msgs: [{ package: pkg, send, maxDeposit }],
    };
    return streamAsyncIterable<pb.RunResponseJson>('Run', JSON.stringify(req));
  }

  async addressToBech32(address: string): Promise<string> {
    const req: pb.AddressToBech32RequestJson = { address };
    const res = await this.#invoke<pb.AddressToBech32ResponseJson>('AddressToBech32', req);
    return res.bech32Address ?? '';
  }

  async addressFromMnemonic(mnemonic: string): Promise<string> {
    const req: pb.AddressFromMnemonicRequestJson = { mnemonic };
    const res = await this.#invoke<pb.AddressFromMnemonicResponseJson>('AddressFromMnemonic', req);
    return res.address ?? '';
  }

  async addressFromBech32(bech32Address: string): Promise<string> {
    const req: pb.AddressFromBech32RequestJson = { bech32Address };
    const res = await this.#invoke<pb.AddressFromBech32ResponseJson>('AddressFromBech32', req);
    return res.address ?? '';
  }

  async pubKeyBytesFromBech32(bech32PubKey: string): Promise<string> {
    const req: pb.PubKeyBytesFromBech32RequestJson = { bech32PubKey };
    const res = await this.#invoke<pb.PubKeyBytesFromBech32ResponseJson>(
      'PubKeyBytesFromBech32',
      req,
    );
    return res.pubKeyBytes ?? '';
  }

  async validateMnemonicWord(word: string): Promise<boolean> {
    const req: pb.ValidateMnemonicWordRequestJson = { word };
    const res = await this.#invoke<pb.ValidateMnemonicWordResponseJson>(
      'ValidateMnemonicWord',
      req,
    );
    return res.valid ?? false;
  }

  async validateMnemonicPhrase(phrase: string): Promise<boolean> {
    const req: pb.ValidateMnemonicPhraseRequestJson = { phrase };
    const res = await this.#invoke<pb.ValidateMnemonicPhraseResponseJson>(
      'ValidateMnemonicPhrase',
      req,
    );
    return res.valid ?? false;
  }

  async signTx(
    txJson: string,
    address: string,
    accountNumber: string = '0',
    sequenceNumber: string = '0',
  ): Promise<pb.SignTxResponseJson> {
    const req: pb.SignTxRequestJson = { txJson, address, accountNumber, sequenceNumber };
    return this.#invoke('SignTx', req);
  }

  async estimateGas(
    txJson: string,
    address: string,
    securityMargin: number,
    updateTx: boolean,
  ): Promise<pb.EstimateGasResponseJson> {
    const req: pb.EstimateGasRequestJson = { txJson, address, securityMargin, updateTx };
    return this.#invoke('EstimateGas', req);
  }

  async estimateTxFees(
    txJson: string,
    address: string,
    gasSecurityMargin: number,
    gasPriceSecurityMargin: number,
    updateTx: boolean,
  ): Promise<pb.EstimateTxFeesResponseJson> {
    const req: pb.EstimateTxFeesRequestJson = {
      txJson,
      address,
      gasSecurityMargin,
      gasPriceSecurityMargin,
      updateTx,
    };
    return this.#invoke('EstimateTxFees', req);
  }

  async makeCallTx(
    packagePath: string,
    fnc: string,
    args: string[],
    gasFee: string,
    gasWanted: string,
    callerAddress: string,
    send?: pb.CoinJson[],
    maxDeposit?: pb.CoinJson[],
    memo?: string,
  ): Promise<pb.MakeTxResponseJson> {
    const req: pb.MakeCallTxRequestJson = {
      gasFee,
      gasWanted,
      memo,
      callerAddress,
      Msgs: [{ packagePath, fnc, args, send, maxDeposit }],
    };
    return this.#invoke('MakeCallTx', req);
  }

  async makeSendTx(
    toAddress: string,
    amount: pb.CoinJson[],
    gasFee: string,
    gasWanted: string,
    callerAddress: string,
    memo?: string,
  ): Promise<pb.MakeTxResponseJson> {
    const req: pb.MakeSendTxRequestJson = {
      gasFee,
      gasWanted,
      memo,
      callerAddress,
      Msgs: [{ toAddress, amount }],
    };
    return this.#invoke('MakeSendTx', req);
  }

  async makeRunTx(
    pkg: string,
    gasFee: string,
    gasWanted: string,
    callerAddress: string,
    send?: pb.CoinJson[],
    maxDeposit?: pb.CoinJson[],
    memo?: string,
  ): Promise<pb.MakeTxResponseJson> {
    const req: pb.MakeRunTxRequestJson = {
      gasFee,
      gasWanted,
      memo,
      callerAddress,
      Msgs: [{ package: pkg, send, maxDeposit }],
    };
    return this.#invoke('MakeRunTx', req);
  }

  async createSession(
    creatorAddress: string,
    sessionKey: string,
    expiresAt: string,
    allowPaths: string[],
    spendLimit: pb.CoinJson[],
    spendPeriod: string,
    gasFee: string,
    gasWanted: string,
    memo?: string,
  ): Promise<AsyncIterable<pb.CreateSessionResponseJson>> {
    const req: pb.CreateSessionRequestJson = {
      gasFee,
      gasWanted,
      memo,
      creatorAddress,
      Msgs: [{ sessionKey, expiresAt, allowPaths, spendLimit, spendPeriod }],
    };
    return streamAsyncIterable<pb.CreateSessionResponseJson>('CreateSession', JSON.stringify(req));
  }

  async makeCreateSessionTx(
    creatorAddress: string,
    sessionKey: string,
    expiresAt: string,
    allowPaths: string[],
    spendLimit: pb.CoinJson[],
    spendPeriod: string,
    gasFee: string,
    gasWanted: string,
    memo?: string,
  ): Promise<pb.MakeTxResponseJson> {
    const req: pb.CreateSessionRequestJson = {
      gasFee,
      gasWanted,
      memo,
      creatorAddress,
      Msgs: [{ sessionKey, expiresAt, allowPaths, spendLimit, spendPeriod }],
    };
    return this.#invoke('MakeCreateSessionTx', req);
  }

  async revokeSession(
    creatorAddress: string,
    sessionKey: string,
    gasFee: string,
    gasWanted: string,
    memo?: string,
  ): Promise<AsyncIterable<pb.RevokeSessionResponseJson>> {
    const req: pb.RevokeSessionRequestJson = {
      gasFee,
      gasWanted,
      memo,
      creatorAddress,
      Msgs: [{ sessionKey }],
    };
    return streamAsyncIterable<pb.RevokeSessionResponseJson>('RevokeSession', JSON.stringify(req));
  }

  async makeRevokeSessionTx(
    creatorAddress: string,
    sessionKey: string,
    gasFee: string,
    gasWanted: string,
    memo?: string,
  ): Promise<pb.MakeTxResponseJson> {
    const req: pb.RevokeSessionRequestJson = {
      gasFee,
      gasWanted,
      memo,
      creatorAddress,
      Msgs: [{ sessionKey }],
    };
    return this.#invoke('MakeRevokeSessionTx', req);
  }

  async revokeAllSessions(
    creatorAddress: string,
    gasFee: string,
    gasWanted: string,
    memo?: string,
  ): Promise<AsyncIterable<pb.RevokeAllSessionsResponseJson>> {
    const req: pb.RevokeAllSessionsRequestJson = { gasFee, gasWanted, memo, creatorAddress };
    return streamAsyncIterable<pb.RevokeAllSessionsResponseJson>(
      'RevokeAllSessions',
      JSON.stringify(req),
    );
  }

  async makeRevokeAllSessionsTx(
    creatorAddress: string,
    gasFee: string,
    gasWanted: string,
    memo?: string,
  ): Promise<pb.MakeTxResponseJson> {
    const req: pb.RevokeAllSessionsRequestJson = { gasFee, gasWanted, memo, creatorAddress };
    return this.#invoke('MakeRevokeAllSessionsTx', req);
  }

  async broadcastTxCommit(
    signedTxJson: string,
  ): Promise<AsyncIterable<pb.BroadcastTxCommitResponseJson>> {
    const req: pb.BroadcastTxCommitRequestJson = { tx_json: signedTxJson };
    return streamAsyncIterable<pb.BroadcastTxCommitResponseJson>(
      'BroadcastTxCommit',
      JSON.stringify(req),
    );
  }

  // debug
  async hello(name: string): Promise<string> {
    const req: pb.HelloRequestJson = { Name: name };
    const res = await this.#invoke<pb.HelloResponseJson>('Hello', req);
    return res.Greeting ?? '';
  }

  // debug
  async helloStream(name: string): Promise<AsyncIterable<pb.HelloStreamResponseJson>> {
    const req: pb.HelloStreamRequestJson = { Name: name };
    return streamAsyncIterable<pb.HelloStreamResponseJson>('HelloStream', JSON.stringify(req));
  }
}
