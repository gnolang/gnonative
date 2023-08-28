export type GnoConfig = {
  TxCfg: {
    RootCfg: {
      Home: string;
      Remote: string;
      Quiet: boolean;
      InsecurePasswordStdin: boolean;
      Config: string;
    };
    GasWanted: number;
    GasFee: string;
    Memo: string;
    Broadcast: boolean;
    ChainID: string;
  };
  KeyName: string;
  Password: string;
};
