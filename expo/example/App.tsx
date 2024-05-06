import { GnokeyProvider } from '@gnolang/gnonative';
import { useGnokeyContext } from '@gnolang/gnonative/provider/gnokey-provider';
import React, { useEffect } from 'react';
import { StyleSheet, Text, View } from 'react-native';

const config = {
  remote: 'https://gno.berty.io',
  chain_id: 'dev',
};

export default function App() {
  return (
    <GnokeyProvider config={config}>
      <View style={styles.container}>
        <InnerApp />
      </View>
    </GnokeyProvider>
  );
}

const InnerApp = () => {
  const gno = useGnokeyContext();

  useEffect(() => {
    (async () => {
      try {
        const accounts = await gno.listKeyInfo();
        console.log(accounts);
      } catch (error) {
        console.log(error);
      }
    })();
  }, []);

  return (
    <View>
      <Text>Inner App</Text>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});
