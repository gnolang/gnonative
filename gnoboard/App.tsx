import React, { useState, useEffect } from 'react';
import { StatusBar } from 'expo-status-bar';
import { StyleSheet, Text, View } from 'react-native';

import { GoBridge } from '@gno/native_modules/GoBridge';

export default function App() {
  const [greeting, setGreeting] = React.useState<string>('');

  useEffect(() => {
    const getGreeting = async () => setGreeting(await GoBridge.hello('Gno.land'));

    getGreeting();
    }, []);


  return (
    <View style={styles.container}>
      <Text>{greeting}</Text>
      <StatusBar style="auto" />
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
