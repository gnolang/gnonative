// React provider for the /native client. Clone of src/provider/gnonative-provider.tsx typed to
// GnoNativeClient (the connect-free client).
import { createContext, useContext, useEffect, useState } from 'react';

import { GnoNativeClient } from './client';
import { Config } from './types';

export interface GnoNativeContextProps {
  gnonative: GnoNativeClient;
}

interface GnoNativeProviderProps {
  config: Config;
  children: React.ReactNode;
}

const GnoNativeContext = createContext<GnoNativeContextProps | null>(null);

const GnoNativeProvider: React.FC<GnoNativeProviderProps> = ({ children, config }) => {
  const [initialized, setInitialized] = useState(false);
  const [client] = useState<GnoNativeClient>(new GnoNativeClient(config));

  useEffect(() => {
    (async () => {
      await init(config);
      setInitialized(true);
    })();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  async function init(config: Config): Promise<boolean> {
    console.log(
      '🍄 Initializing GnoNative (native) Context on remote: %s chain_id: %s',
      config.remote,
      config.chain_id,
    );

    try {
      await client.initClient();
    } catch (error) {
      console.error(error);
      return false;
    }

    return true;
  }

  const value = {
    gnonative: client,
  };

  if (!initialized) {
    return null;
  }

  return <GnoNativeContext.Provider value={value}>{children}</GnoNativeContext.Provider>;
};

function useGnoNativeContext() {
  const context = useContext(GnoNativeContext) as GnoNativeContextProps;

  if (context === undefined) {
    throw new Error('useGnoNativeContext must be used within a GnoNativeProvider');
  }
  return context;
}

export { GnoNativeProvider, useGnoNativeContext };
