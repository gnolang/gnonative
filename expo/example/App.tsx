import { GnokeyProvider, useGnokeyContext } from '@gnolang/gnonative';
import React, { useEffect, useState } from 'react';
import { StyleSheet, Text, View } from 'react-native';

const config = {
  remote: 'https://gno.berty.io',
  chain_id: 'dev',
};

export default function App() {
  return (
    <GnokeyProvider config={config}>
      <InnerApp />
    </GnokeyProvider>
  );
}

const InnerApp = () => {
  const gno = useGnokeyContext();
  const [greeting, setGreeting] = useState('');

  useEffect(() => {
    (async () => {
      try {
        const accounts = await gno.listKeyInfo();
        console.log(accounts);

        setGreeting(await gno.hello('Gno'));

        for await (const res of await gno.helloStream('Gno')) {
          console.log(res.greeting);
        }
      } catch (error) {
        console.log(error);
      }
    })();
  }, []);

  return (
    <View style={styles.container}>
      <Text>Gnonative App</Text>
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
