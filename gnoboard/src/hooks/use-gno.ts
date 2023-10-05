import { GnoAccount } from '@gno/native_modules/types';
import * as GoBridge from '@gno/utils/bridge';

interface GnoResponse {
  createAccount: (nameOrBech32: string, mnemonic: string, password: string) => Promise<GnoAccount>;
  generateRecoveryPhrase: () => Promise<string>;
  initBridge: () => Promise<boolean>;
  listKeyInfo: () => Promise<GnoAccount[]>;
  selectAccount: (nameOrBech32: string) => Promise<GnoAccount>;
  setPassword: (password: string) => Promise<void>;
}

export const useGno = (): GnoResponse => {
  const createAccount = async (nameOrBech32: string, mnemonic: string, password: string) => {
    console.log('nameOrBech32', nameOrBech32);
    console.log('mnemonic', mnemonic);

    return await GoBridge.createAccount(nameOrBech32, mnemonic, '', password, 0, 0);
  };

  const generateRecoveryPhrase = async () => {
    return await GoBridge.generateRecoveryPhrase();
  };

  const initBridge = async () => {
    return await GoBridge.initBridge();
  };

  const listKeyInfo = async () => {
    return await GoBridge.listKeyInfo();
  };

  const selectAccount = async (nameOrBech32: string) => {
    return await GoBridge.selectAccount(nameOrBech32);
  };

  const setPassword = async (password: string) => {
    return await GoBridge.setPassword(password);
  };

  return {
    setPassword,
    createAccount,
    generateRecoveryPhrase,
    initBridge,
    listKeyInfo,
    selectAccount,
  };
};
