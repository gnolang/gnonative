export type GnoConfig = {
  Remote: string;
  ChainID: string;
  KeyName: string;
  Password: string;
  GasFee: string;
  GasWanted: number;
  Mnemonic: string;
};

export type GnoAccount = {
  address_b64: string;
  name: string;
  pubKey_b64: string;
  type: number;
};
