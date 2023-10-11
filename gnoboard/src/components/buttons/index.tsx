import React from 'react';
import { ActivityIndicator, TouchableOpacity, Text, View } from 'react-native';
import { StyleSheet } from 'react-native';
import { colors } from '@gno/styles/colors';

const styles = StyleSheet.create({
  button: {
    backgroundColor: colors.blue,
    borderRadius: 4,
    marginVertical: 8,
    padding: 8,
  },
  buttonText: {
    color: colors.white,
    fontSize: 16,
    fontWeight: 'bold',
    textAlign: 'center',
  },
});

export interface Props {
  title: string;
  onPress: () => void;
  loading?: boolean;
}

const Button: React.FC<Props> = ({ title, onPress, loading = false }) => {
  return (
    <TouchableOpacity onPress={onPress}>
      <View style={styles.button}>{loading ? <ActivityIndicator size='small' /> : <Text style={styles.buttonText}>{title}</Text>}</View>
    </TouchableOpacity>
  );
};

export default Button;
