import React, { useState, useEffect } from "react";
import { ScrollView, Text, View } from "react-native";
import { consoleStyleSheet } from "./styles";

export const ConsoleView = ({ text }: { text: string }) => {
  const [consoleText, setConsoleText] = useState([""]);

  useEffect(() => {
    setConsoleText((prev) => [...prev, text]);
  }, [text]);

  return (
    <ScrollView contentContainerStyle={consoleStyleSheet.container}>
      <Text style={consoleStyleSheet.text}>{consoleText.join("\n")}</Text>
    </ScrollView>
  );
};
