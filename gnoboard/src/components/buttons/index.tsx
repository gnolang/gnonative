import React from 'react';
import { ActivityIndicator, TouchableOpacity, Text, View, TouchableOpacityProps } from 'react-native';
import { StyleSheet } from 'react-native';
import { colors } from '@gno/styles/colors';

const styles = StyleSheet.create({
  button: {
    backgroundColor: colors.blue,
    borderRadius: 28,
    marginVertical: 8,
    paddingHorizontal: 24,
    padding: 8,
  },
  buttonText: {
    color: colors.white,
    fontSize: 16,
    fontWeight: 'bold',
    textAlign: 'center',
  },
});

export type Props = {
  title: string;
  onPress: () => void;
  loading?: boolean;
} & TouchableOpacityProps;

const Button: React.FC<Props> = ({ title, onPress, loading = false, ...rest }) => {
  return (
    <TouchableOpacity onPress={onPress} {...rest}>
      <View style={styles.button}>{loading ? <ActivityIndicator size='small' /> : <Text style={styles.buttonText}>{title}</Text>}</View>
    </TouchableOpacity>
  );
};

export default Button;
