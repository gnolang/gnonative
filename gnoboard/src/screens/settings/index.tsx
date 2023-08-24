import TextInput from "@gno/components/textinput";
import { GoBridge } from "@gno/native_modules";
import { GnoConfig } from "@gno/native_modules/types";
import { screenStyleSheet as styles } from "@gno/styles";
import React, { useEffect } from "react";
import { Text, ActivityIndicator, ScrollView } from "react-native";

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

  // view with scrool view inside to avoid keyboard covering the input
  return (
    <ScrollView contentContainerStyle={styles.scrollViewContent}>
      <Text>Home:</Text>
      <TextInput value={config.TxCfg.RootCfg.Home} />
      <Text>Remote:</Text>
      <TextInput value={config.TxCfg.RootCfg.Remote} />
      <Text>Quiet:</Text>
      <TextInput value={String(config.TxCfg.RootCfg.Quiet)} />
      <Text>Insecure Password Stdin:</Text>
      <TextInput value={String(config.TxCfg.RootCfg.InsecurePasswordStdin)} />
      <Text>Config:</Text>
      <TextInput value={config.TxCfg.RootCfg.Config} />

      <Text>Broadcast:</Text>
      <TextInput value={String(config.TxCfg.Broadcast)} />
      <Text>ChainID:</Text>
      <TextInput value={config.TxCfg.ChainID} />
      <Text>Gas Fee:</Text>
      <TextInput value={config.TxCfg.GasFee} />
      <Text>Gas Wanted:</Text>
      <TextInput value={String(config.TxCfg.GasWanted)} />
      <Text>Memo:</Text>
      <TextInput value={String(config.TxCfg.Memo)} />
      <Text>Password:</Text>
      <TextInput value={config.Password} />
    </ScrollView>
  );
}

export default SettingsScreen;
