import { createContext, useContext, useEffect, useState } from 'react';

import { GnoNativeApi } from '../api';
import { Config } from '../api/types';

export interface GnoNativeContextProps {
  gnonative: GnoNativeApi;
}

interface GnoNativeProviderProps {
  config: Config;
  children: React.ReactNode;
}

const GnoNativeContext = createContext<GnoNativeContextProps | null>(null);

const GnoNativeProvider: React.FC<GnoNativeProviderProps> = ({ children, config }) => {
  const [initialized, setInitialized] = useState(false);
  const [api] = useState<GnoNativeApi>(new GnoNativeApi(config));

  useEffect(() => {
    (async () => {
      await init(config);
      setInitialized(true);
    })();
  }, []);

  async function init(config): Promise<boolean> {
    console.log(
      '🍄 Initializing GnoNative Context on remote: %s chain_id: %s',
      config.remote,
      config.chain_id,
    );

    try {
      await api.initClient();
    } catch (error) {
      console.error(error);
      return false;
    }

    return true;
  }

  const value = {
    gnonative: api,
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
