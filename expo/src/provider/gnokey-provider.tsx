import {
  ListKeyInfoRequest,
  SelectAccountRequest,
  SelectAccountResponse,
  SetChainIDRequest,
  SetRemoteRequest,
} from '@buf/gnolang_gnonative.bufbuild_es/gnonativetypes_pb';
import { GnoNativeService } from '@buf/gnolang_gnonative.connectrpc_es/rpc_connect';
import { PromiseClient } from '@connectrpc/connect';
import { createContext, useContext, useEffect, useState } from 'react';

import { GoBridge } from '../GoBridge';
import * as Grpc from '../grpc/client';
import { GnoAccount } from '../hooks/types';

export interface GnokeyContextProps {
  initGnokey: (config: ConfigProps) => Promise<boolean>;
  selectAccount: (nameOrBech32: string) => Promise<SelectAccountResponse>;
  listKeyInfo: () => Promise<GnoAccount[]>;
}

interface ConfigProps {
  remote: string;
  chain_id: string;
}

interface GnokeyProviderProps {
  config: ConfigProps;
  children: React.ReactNode;
}

enum BridgeStatus {
  Stopped,
  Starting,
  Started,
}

const GnokeyContext = createContext<GnokeyContextProps | null>(null);

let clientInstance: PromiseClient<typeof GnoNativeService> | undefined = undefined;
let bridgeStatus: BridgeStatus = BridgeStatus.Stopped;

const GnokeyProvider: React.FC<GnokeyProviderProps> = ({ children, config }) => {
  const [initialized, setInitialized] = useState(false);

  useEffect(() => {
    (async () => {
      await initGnokey(config);
      setInitialized(true);
    })();
  }, []);

  async function initGnokey(config): Promise<boolean> {
    console.log(
      '🍄 Initializing Gnokey on remote: %s chain_id: %s',
      config.remote,
      config.chain_id,
    );

    if (bridgeStatus === BridgeStatus.Stopped) {
      await initBridge();
    }

    if (clientInstance) {
      console.error('GoBridge already initialized.');
      return true;
    }

    const port = await GoBridge.getTcpPort();
    console.log('GoBridge GRPC client instance port: %s', port);
    clientInstance = Grpc.createClient(port);
    console.log('GoBridge GRPC client instance. Done.');

    try {
      await clientInstance.setRemote(new SetRemoteRequest({ remote: 'gno.land:26657' }));
      await clientInstance.setChainID(new SetChainIDRequest({ chainId: 'portal-loop' }));

      console.log('✅ Gnokey bridge initialized.');
    } catch (error) {
      console.error(error);
      return false;
    }

    return true;
  }

  const initBridge = async () => {
    if (bridgeStatus === BridgeStatus.Stopped) {
      console.log('Initializing bridge...');
      bridgeStatus = BridgeStatus.Starting;
      await GoBridge.initBridge();
      console.log('Bridge initialized.');
      bridgeStatus = BridgeStatus.Started;
    }
  };

  const getClient = async () => {
    if (!clientInstance) {
      throw new Error('GoBridge client instance not initialized.');
    }

    return clientInstance;
  };

  const selectAccount = async (nameOrBech32: string) => {
    const client = await getClient();
    const response = await client.selectAccount(
      new SelectAccountRequest({
        nameOrBech32,
      }),
    );
    return response;
  };

  const listKeyInfo = async () => {
    const client = await getClient();
    const response = await client.listKeyInfo(new ListKeyInfoRequest());
    return response.keys;
  };

  const value = {
    selectAccount,
    initGnokey,
    listKeyInfo,
  };

  if (!initialized) {
    return null;
  }

  return <GnokeyContext.Provider value={value}>{children}</GnokeyContext.Provider>;
};

function useGnokeyContext() {
  const context = useContext(GnokeyContext) as GnokeyContextProps;

  if (context === undefined) {
    throw new Error('useGnokeyContext must be used within a GnokeyProvider');
  }
  return context;
}

export { useGnokeyContext, GnokeyProvider };
