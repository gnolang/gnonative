import { GoBridge } from "@gno/native_modules";
import { GnoConfig } from "@gno/native_modules/types";
import React, { useEffect } from "react";
import {
  Text,
  View,
  ActivityIndicator,
  TextInput,
  StyleSheet,
  ScrollView,
} from "react-native";

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
      <TextInput style={styles.input} value={config.TxCfg.RootCfg.Home} />
      <Text>Remote:</Text>
      <TextInput style={styles.input} value={config.TxCfg.RootCfg.Remote} />
      <Text>Quiet:</Text>
      <TextInput
        style={styles.input}
        value={String(config.TxCfg.RootCfg.Quiet)}
      />
      <Text>Insecure Password Stdin:</Text>
      <TextInput
        style={styles.input}
        value={String(config.TxCfg.RootCfg.InsecurePasswordStdin)}
      />
      <Text>Config:</Text>
      <TextInput style={styles.input} value={config.TxCfg.RootCfg.Config} />

      <Text>Broadcast:</Text>
      <TextInput style={styles.input} value={String(config.TxCfg.Broadcast)} />
      <Text>ChainID:</Text>
      <TextInput style={styles.input} value={config.TxCfg.ChainID} />
      <Text>Gas Fee:</Text>
      <TextInput style={styles.input} value={config.TxCfg.GasFee} />
      <Text>Gas Wanted:</Text>
      <TextInput style={styles.input} value={String(config.TxCfg.GasWanted)} />
      <Text>Memo:</Text>
      <TextInput style={styles.input} value={String(config.TxCfg.Memo)} />
      <Text>Password:</Text>
      <TextInput style={styles.input} value={config.Password} />
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  scrollViewContent: {
    flexGrow: 1,
    justifyContent: "center",
    padding: 20,
  },
  container: {
    flex: 1,
    paddingTop: 340,
    paddingBottom: 300,
    justifyContent: "center",
    paddingHorizontal: 20,
  },
  input: {
    height: 40,
    borderColor: "gray",
    borderWidth: 1,
    padding: 10,
    borderRadius: 5,
    marginBottom: 10,
    width: "100%", // Full width of the container
  },
});

export default SettingsScreen;
