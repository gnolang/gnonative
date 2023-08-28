import React from "react";
import {
  TextInput as RNTextInput,
  StyleSheet,
  TextInputProps,
} from "react-native";

const styles = StyleSheet.create({
  input: {
    height: 40,
    marginVertical: 8,
    borderWidth: 1,
    padding: 10,
    borderRadius: 5,
    width: "100%", // Full width of the container
  },
});

const TextInput = (props: TextInputProps) => {
  return <RNTextInput {...props} style={[styles.input, props.style]} />;
};

export default TextInput;
