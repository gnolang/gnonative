import Button from "@gno/components/button";
import { ConsoleView } from "@gno/components/consoleview";
import TextInput from "@gno/components/textinput";
import { GoBridge } from "@gno/native_modules";
import { screenStyleSheet as styles } from "@gno/styles";
import { useState } from "react";
import { Linking, ScrollView, StyleSheet, Text, View } from "react-native";

function HomeScreen() {
  const [postContent, setPostContent] = useState("");
  const [appConsole, setAppConsole] = useState<string>("");
  const [loading, setLoading] = useState<string | undefined>(undefined);

  const onPostPress = async () => {
    setLoading("Replying to a post...");
    setAppConsole("replying to a post...");
    GoBridge.createReply(postContent, "")
      .then((data) => {
        setAppConsole(data);
        setPostContent("");
      })
      .catch((err) => {
        setAppConsole(err);
      })
      .finally(() => setLoading(undefined));
  };

  const onCreateAccountPress = async () => {
    setLoading("Creating account...");
    GoBridge.hello("create account")
      .then((data) => {
        setAppConsole(data);
      })
      .catch((err) => {
        setAppConsole(err);
      })
      .finally(() => setLoading(undefined));
  };

  const loadInBrowser = () => {
    Linking.openURL(
      "http://testnet.gno.berty.io/r/demo/boards:gnomobile/1"
    ).catch((err) => console.error("Couldn't load page", err));
  };

  const customStyle = StyleSheet.create({
    loadingContainer: {
      flex: 1,
      justifyContent: "center",
      alignItems: "center",
    },
  });

  return (
    <ScrollView contentContainerStyle={styles.scrollViewContent}>
      <Button
        title="Create Account"
        loading={loading}
        onPress={onCreateAccountPress}
      />
      <Text>Content:</Text>
      <View style={customStyles.sendGroupLikeWhatsapp}>
        <TextInput
          style={customStyles.inputMsg}
          value={postContent}
          onChangeText={setPostContent}
        />
        <Button title="Send" loading={loading} onPress={onPostPress} />
      </View>
      <ConsoleView text={appConsole} />
      <Button
        title="Open http://testnet.gno.berty.io/r/demo/boards:gnomobile/1"
        loading={loading}
        onPress={loadInBrowser}
      />
    </ScrollView>
  );
}

const customStyles = StyleSheet.create({
  sendGroupLikeWhatsapp: {
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "space-between",
  },
  inputMsg: { width: "80%" },
});

export default HomeScreen;
