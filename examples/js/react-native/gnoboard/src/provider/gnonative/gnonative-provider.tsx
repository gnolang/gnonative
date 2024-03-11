import React, { createContext, useEffect, useRef } from 'react';
import { GnoResponse, useGno } from '@gno/hooks/use-gno';
import { AppState } from 'react-native';

interface GnoNativeContextProps {
  gno: GnoResponse;
}

export const GnoNativeContext = createContext<GnoNativeContextProps | null>(null);

export const GnoNativeProvider: React.FC<React.PropsWithChildren<unknown>> = ({ children }) => {
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

  return <GnoNativeContext.Provider value={{ gno }}>{children}</GnoNativeContext.Provider>;
};
