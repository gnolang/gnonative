import { createContext, useContext, useState, useCallback } from 'react';

import { KeyInfo } from '@gnolang/gnonative';

interface GnoboardProviderProps {
  children: React.ReactNode;
}
interface GnoboardContextType {
  account: KeyInfo | undefined;
  setAccount: (keyInfo : KeyInfo | undefined) => void;
}

const GnoboardContext = createContext<GnoboardContextType | null>(null);

const GnoboardProvider: React.FC<GnoboardProviderProps> = ({ children }) => {
  const [account, setAccount] = useState<KeyInfo | undefined>(undefined)

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
