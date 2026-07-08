// Uses the buf/connect-free entry point: no polyfills, no @bufbuild/@connectrpc.
import { GnoNativeProvider, useGnoNativeContext } from '@gnolang/gnonative/native';
import React, { useEffect, useState } from 'react';
import { StyleSheet, Text, View } from 'react-native';

const config = {
  remote: 'https://gno.berty.io',
  chain_id: 'dev',
};

export default function App() {
  return (
    <GnoNativeProvider config={config}>
      <InnerApp />
    </GnoNativeProvider>
  );
}

const InnerApp = () => {
  const { gnonative } = useGnoNativeContext();
  const [greeting, setGreeting] = useState('');

  useEffect(() => {
    (async () => {
      try {
        const accounts = await gnonative.listKeyInfo();
        console.log(accounts);

        const remote = await gnonative.getRemote();
        const chainId = await gnonative.getChainID();
        console.log(`Remote ${remote} ChainId ${chainId}`);

        const phrase = await gnonative.generateRecoveryPhrase();
        // On the /native path, addresses are base64 strings (not Uint8Array).
        const address = await gnonative.addressFromMnemonic(phrase);
        const addressStr = await gnonative.addressToBech32(address);

        console.log('Phrase:', phrase);
        console.log('Address:', addressStr);

        setGreeting(await gnonative.hello('Gno'));

        // Streaming works with no polyfills. The field is `Greeting` (proto json_name).
        for await (const res of await gnonative.helloStream('Gno')) {
          console.log(res.Greeting);
        }
      } catch (error) {
        console.log(error);
      }
    })();
  }, []);

  return (
    <View style={styles.container}>
      <Text>Gnonative Native App</Text>
      <Text>{greeting}</Text>
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
