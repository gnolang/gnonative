import React, { createContext, useEffect, useRef } from 'react';
import { useGno, GnoResponse } from '@gno/hooks/use-gno';
import { AppState } from 'react-native';

interface GnomobileContextProps {
  gno: GnoResponse;
}

export const GnomobileContext = createContext<GnomobileContextProps | null>(null);

export const GnomobileProvider: React.FC<React.PropsWithChildren<unknown>> = ({ children }) => {
  const gno = useGno();

  const appState = useRef(AppState.currentState);

  useEffect(() => {
    const subscription = AppState.addEventListener('change', (nextAppState) => {
      if (appState.current.match(/inactive|background/) && nextAppState === 'active') {
        console.log('initBridge()');
        gno.initBridge();
      }

      if (appState.current.match(/active/) && nextAppState === 'background') {
        console.log('closeBridge()');
        gno.closeBridge();
      }

      appState.current = nextAppState;
      console.log('AppState', appState.current);
    });

    return () => {
      subscription.remove();
    };
  }, []);

  return <GnomobileContext.Provider value={{ gno }}>{children}</GnomobileContext.Provider>;
};
