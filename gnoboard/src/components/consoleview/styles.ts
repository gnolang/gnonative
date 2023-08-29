import { StyleSheet } from "react-native";

export const consoleStyleSheet = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#000000",
    borderColor: "#000000",
    borderWidth: 1,
    borderRadius: 5,
    padding: 5,
    margin: 5,
    maxHeight: 400,
  },
  contentContainer: {
    minHeight: 150,
  },
  text: {
    color: "#ffffff",
    fontSize: 12,
    fontFamily: "monospace",
  },
});
