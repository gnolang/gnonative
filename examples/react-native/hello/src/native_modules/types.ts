import { KeyInfo } from '@gno/api/gnonativetypes_pb';

export type GnoConfig = {
  Remote: string;
  ChainID: string;
  KeyName: string;
  Password: string;
  GasFee: string;
  GasWanted: number;
  Mnemonic: string;
};

export type NetworkMetainfo = {
  chainId: string;
  chainName: string;
  gnoAddress: string;
};

export type GnoAccount = KeyInfo;
