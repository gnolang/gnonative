import React, { useState, useEffect } from "react";
import { Button, ScrollView, Text, View } from "react-native";
import { consoleStyleSheet } from "./styles";

export const ConsoleView = ({ text }: { text: string }) => {
  const [consoleText, setConsoleText] = useState([""]);

  useEffect(() => {
    setConsoleText((prev) => [...prev, text]);
  }, [text]);

  return (
    <View style={consoleStyleSheet.container}>
      <ScrollView
        contentContainerStyle={consoleStyleSheet.contentContainer}
        showsVerticalScrollIndicator={true}
      >
        <Text style={consoleStyleSheet.text}>{consoleText.join("\n")}</Text>
      </ScrollView>
      <Button title="clear" onPress={() => setConsoleText([""])} />
    </View>
  );
};
