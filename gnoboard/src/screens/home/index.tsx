import Button from "@gno/components/button";
import TextInput from "@gno/components/textinput";
import { GoBridge } from "@gno/native_modules";
import { screenStyleSheet as styles } from "@gno/styles";
import { useState } from "react";
import { Linking, ScrollView, Text, View } from "react-native";

function HomeScreen() {
  const [postContent, setPostContent] = useState("");

  const onPostPress = async () => {
    const data = await GoBridge.hello(postContent);
    console.log(data);
  };

  const loadInBrowser = () => {
    Linking.openURL(
      "http://testnet.gno.berty.io/r/demo/boards:gnomobile/1"
    ).catch((err) => console.error("Couldn't load page", err));
  };

  return (
    <ScrollView contentContainerStyle={styles.scrollViewContent}>
      <Text>Type your message:</Text>
      <TextInput value={postContent} onChangeText={setPostContent} />
      <Button title="Post" onPress={onPostPress} />
      <Button title="Check Render" onPress={loadInBrowser} />
    </ScrollView>
  );
}

export default HomeScreen;
