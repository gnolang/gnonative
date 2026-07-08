import { GnoNativeProvider, useGnoNativeContext } from "@gnolang/gnonative";
import React from "react";
import { StyleSheet, TextInput, View } from "react-native";
import { StatusBar } from "expo-status-bar";

const config = {
  remote: "https://gno.berty.io",
  chain_id: "dev",
};

export default function App() {
  return (
    <GnoNativeProvider config={config}>
      <InnerApp />
    </GnoNativeProvider>
  );
}

function InnerApp() {
  const { gnonative } = useGnoNativeContext();
  const [board, setBoard] = React.useState("");

  React.useEffect(() => {
    gnonative
      .render("gno.land/r/demo/boards", "testboard/1")
      .then((res) => setBoard(res))
      .catch((err) => setBoard(String(err)));
  }, []);

  return (
    <View style={styles.container}>
      <TextInput multiline={true} numberOfLines={40} value={board} />
      <StatusBar style="auto" />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 20,
    backgroundColor: "#fff",
    alignItems: "center",
    justifyContent: "center",
  },
});
