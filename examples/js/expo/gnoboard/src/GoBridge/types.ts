import { KeyInfo } from '@buf/gnolang_gnonative.bufbuild_es/gnonativetypes_pb';

export type GnoConfig = {
  Remote: string;
  ChainID: string;
  KeyName: string;
  Password: string;
  GasFee: string;
  GasWanted: bigint;
  Mnemonic: string;
};

export type NetworkMetainfo = {
  chainId: string;
  chainName: string;
  gnoAddress: string;
};

export type GnoAccount = KeyInfo;
