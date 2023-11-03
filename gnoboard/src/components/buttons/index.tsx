import React from 'react';
import { ActivityIndicator, TouchableOpacityProps } from 'react-native';
import { StyleSheet } from 'react-native';
import { colors } from '@gno/styles/colors';
import styled from 'styled-components/native';
import Text from '../texts';

type ButtonVariant = 'primary' | 'secondary' | 'tertiary' | 'white' | 'primary2' | 'primary-red' | 'secondary-red';

export type Props = {
  title: string;
  onPress: () => void;
  loading?: boolean;
  variant: ButtonVariant;
} & TouchableOpacityProps;

const Button: React.FC<Props> = ({ title, onPress, loading = false, variant, ...rest }) => {
  return (
    <TouchableOpacityButton variant={variant} onPress={onPress} {...rest}>
      {loading ? <ActivityIndicator size='small' /> : <Text.Body style={styles.buttonText}>{title}</Text.Body>}
    </TouchableOpacityButton>
  );
};

const TouchableOpacityButton = styled.TouchableOpacity<{ variant: ButtonVariant }>`
  background-color: ${(props) => getStyle(props.variant)};
  width: 100%;
  height: 48px;
  justify-content: center;
  border-radius: 28px;
`;

const getStyle = (variant: ButtonVariant) => {
  switch (variant) {
    case 'primary':
      return colors.primary;
    case 'secondary':
      return colors.button.secondary;
    case 'tertiary':
      return colors.tertiary;
    case 'white':
      return colors.white;
    case 'primary2':
      return 'transparent';
    case 'primary-red':
      return colors.red[500];
    case 'secondary-red':
      return colors.red[300];
    default:
      return colors.blue;
  }
};

const styles = StyleSheet.create({
  buttonText: {
    color: colors.white,
    fontWeight: 'bold',
    textAlign: 'center',
  },
});

export default Button;
