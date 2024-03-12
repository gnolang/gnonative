import React, { ReactNode } from 'react';
import { SafeAreaView } from 'react-native';

interface Props {
  children: ReactNode;
}

const Container: React.FC<Props> = ({ children }) => {
  return <SafeAreaView style={{ flex: 1 }}>{children}</SafeAreaView>;
};

export default Container;
