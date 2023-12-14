import React from "react";
import { StyleSheet, TextInput, View } from "react-native";

// order matters here
import "react-native-polyfill-globals/auto";

import { StatusBar } from "expo-status-bar";
import { useGno } from "@gno/hooks/use-gno";

// Polyfill async.Iterator. For some reason, the Babel presets and plugins are not doing the trick.
// Code from here: https://www.typescriptlang.org/docs/handbook/release-notes/typescript-2-3.html#caveats
(Symbol as any).asyncIterator =
  Symbol.asyncIterator || Symbol.for("Symbol.asyncIterator");

export default function App() {
  const gno = useGno();
  const [board, setBoard] = React.useState("");

  React.useEffect(() => {
    gno
      .render("gno.land/r/demo/boards", "gnonative/1")
      .then((res) => setBoard(res))
      .catch((err) => setBoard(err));
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
