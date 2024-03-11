import { View } from 'react-native';
import Alert, { Props } from '.';
import React from 'react';

export default {
  title: 'Alert',
  component: Alert,
  decorators: [
    (Story: React.FC) => (
      <View style={{ width: '100%', height: '100%' }}>
        <Story />
      </View>
    ),
  ],
  args: {
    message: 'Here goes an error.',
    severity: 'error',
  },
};

export const Basic = ({ message = 'Here goes an error.', severity = 'error' }: Props) => {
  return <Alert message={message} severity={severity} />;
};
