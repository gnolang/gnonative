import { StyleSheet, Text, View } from 'react-native';

import * as Gnonative from 'gnonative';

export default function App() {
  const gno = Gnonative.useGno();

  return (
    <View style={styles.container}>
      <Text>Hey {gno.hello('Gno')}</Text>
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
