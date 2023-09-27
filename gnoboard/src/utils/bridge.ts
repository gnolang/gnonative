import { GoBridge } from "@gno/native_modules";

export const initBridge = async (): Promise<boolean> => {
  try {
    console.log("bridge methods: ", Object.keys(GoBridge));

    await GoBridge.initBridge();

    return true;
  } catch (err: any) {
    if (err?.message?.indexOf("already instantiated") !== -1) {
      console.log("bridge already started: ", err);
      return true;
    } else {
      console.error("unable to init bridge: ", err);
    }

    return false;
  }
};

export const closeBridge = async (): Promise<boolean> => {
  try {
    await GoBridge.closeBridge();
    return true;
  } catch (err: any) {
    console.error("unable to close bridge: ", err);
    return false;
  }
};

export const setPassword = async (password: string): Promise<boolean> => {
  try {
    await GoBridge.setPassword(password);
    return true;
  } catch (err: any) {
    console.error("unable to close bridge: ", err);
    return false;
  }
};

export const listKeyInfo = async (): Promise<string[]> => {
  try {
    const keys = await GoBridge.listKeyInfo();
    return keys;
  } catch (err: any) {
    console.error("unable to list keys: ", err);
    return [];
  }
};

export const createAccount = async (
  nameOrBech32: string,
  mnemonic: string,
  bip39Passw: string,
  password: string,
  account: Number,
  index: Number,
): Promise<string> => {
  try {
    const addr = await GoBridge.createAccount(
      nameOrBech32,
      mnemonic,
      bip39Passw,
      password,
      account,
      index,
    );
    return addr;
  } catch (err: any) {
    console.error("unable to close bridge: ", err);
    return "";
  }
};

export const selectAccount = async (nameOrBech32: string): Promise<string> => {
  try {
    const addr = await GoBridge.selectAccount(nameOrBech32);
    return addr;
  } catch (err: any) {
    console.error("unable to close bridge: ", err);
    return "";
  }
};
