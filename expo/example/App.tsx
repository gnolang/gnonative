import * as Gnonative from '@berty/gnonative';
import React, { useEffect, useState } from 'react';
import { StyleSheet, Text, View } from 'react-native';

export default function App() {
  const gno = Gnonative.useGno();
  const [greeting, setGreeting] = useState('');

  useEffect(() => {
    const greeting = async () => {
      try {
        setGreeting(await gno.hello('Gno'));

        for await (const res of await gno.helloStream('Gno')) {
          console.log(res.greeting);
        }
      } catch (error) {
        console.log(error);
      }
    };
    greeting();
  }, []);

  return (
    <View style={styles.container}>
      <Text>Hey {greeting}</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});
