import TextInput from '@gno/components/textinput';
import { GoBridge } from '@gno/native_modules';
import { GnoConfig } from '@gno/native_modules/types';
import { screenStyleSheet as styles } from '@gno/styles';
import React, { useEffect } from 'react';
import { Text, ActivityIndicator, ScrollView } from 'react-native';

function SettingsScreen() {
  const [config, setConfig] = React.useState<GnoConfig | null>(null);

  useEffect(() => {
    const getGreeting = async () => {
      const data = await GoBridge.exportJsonConfig();
      const json = JSON.parse(data) as GnoConfig;
      console.log(json);
      setConfig(json);
    };

    getGreeting();
  }, []);

  if (!config) {
    return <ActivityIndicator />;
  }

  // view with scroll view inside to avoid keyboard covering the input
  return (
    <ScrollView contentContainerStyle={styles.scrollViewContent}>
      <Text>Remote:</Text>
      <TextInput value={config.Remote} />
      <Text>ChainID:</Text>
      <TextInput value={config.ChainID} />
      <Text>KeyName:</Text>
      <TextInput value={String(config.KeyName)} />
      <Text>Password:</Text>
      <TextInput value={config.Password} />
      <Text>Gas Fee:</Text>
      <TextInput value={config.GasFee} />
      <Text>Gas Wanted:</Text>
      <TextInput value={String(config.GasWanted)} />
    </ScrollView>
  );
}

export default SettingsScreen;
