import { createContext, useContext, useState, useCallback } from 'react';

import { KeyInfoJson } from '@gnolang/gnonative';

interface GnoboardProviderProps {
  children: React.ReactNode;
}
interface GnoboardContextType {
  account: KeyInfoJson | undefined;
  setAccount: (keyInfo : KeyInfoJson | undefined) => void;
}

const GnoboardContext = createContext<GnoboardContextType | null>(null);

const GnoboardProvider: React.FC<GnoboardProviderProps> = ({ children }) => {
  const [account, setAccount] = useState<KeyInfoJson | undefined>(undefined)

  const value = {
    account,
    setAccount
  };

  return <GnoboardContext.Provider value={value}>{children}</GnoboardContext.Provider>;
};

function useGnoboardContext() {
  const context = useContext(GnoboardContext) as GnoboardContextType;

  if (context === undefined) {
    throw new Error('useGnoboardContext must be used within a GnoboardProvider');
  }
  return context;
}

export { GnoboardProvider, useGnoboardContext };
