import { SetPasswordRequest, SetPasswordResponse } from '@gno/api/gnomobiletypes_pb';
import { SelectAccountRequest } from '@gno/api/gnomobiletypes_pb';
import { CreateAccountRequest } from '@gno/api/gnomobiletypes_pb';
import { ListKeyInfoRequest } from '@gno/api/gnomobiletypes_pb';
import { DeleteAccountRequest, DeleteAccountResponse } from '@gno/api/gnomobiletypes_pb';
import { CallRequest } from '@gno/api/gnomobiletypes_pb';
import { CallResponse } from '@gno/api/gnomobiletypes_pb';
import * as Grpc from '@gno/grpc/client';
import { GnoAccount } from '@gno/native_modules/types';
import { GoBridge } from '@gno/native_modules';
import { PromiseClient } from '@connectrpc/connect';
import { GnomobileService } from '@gno/api/rpc_connect';
import { GenerateRecoveryPhraseRequest } from '@gno/api/gnomobiletypes_pb';

interface GnoResponse {
  createAccount: (nameOrBech32: string, mnemonic: string, password: string) => Promise<GnoAccount | undefined>;
  generateRecoveryPhrase: () => Promise<string>;
  listKeyInfo: () => Promise<GnoAccount[]>;
  selectAccount: (nameOrBech32: string) => Promise<GnoAccount | undefined>;
  setPassword: (password: string) => Promise<SetPasswordResponse>;
  deleteAccount: (nameOrBech32: string, password: string, skipPassword: boolean) => Promise<DeleteAccountResponse>;
  call: (packagePath: string, fnc: string, args: Array<string>, gasFee: string, gasWanted: number) => Promise<CallResponse>;
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
    clientInstance = await Grpc.createClient(port);
    return clientInstance;
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
    return response.key;
  };

  const setPassword = async (password: string) => {
    const client = await getClient();
    const response = await client.setPassword(new SetPasswordRequest({ password }));
    return response;
  };

  const deleteAccount = async (nameOrBech32: string, password: string, skipPassword: boolean) => {
    const client = await getClient();
    const response = await client.deleteAccount(
      new DeleteAccountRequest({
        nameOrBech32,
        password,
        skipPassword, }));
    return response;
  };

  const call = async (packagePath: string, fnc: string, args: Array<string>, gasFee: string, gasWanted: number) => {
    const client = await getClient();
    const reponse = await client.call(
      new CallRequest({
        packagePath,
        fnc,
        args,
        gasFee,
        gasWanted: BigInt(gasWanted),
      }),
    );
    return reponse;
  };

  return {
    createAccount,
    generateRecoveryPhrase,
    listKeyInfo,
    selectAccount,
    setPassword,
    deleteAccount,
    call,
  };
};
