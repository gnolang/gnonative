import { View } from 'react-native';
import React from 'react';
import ReenterPassword, { Props } from './ReenterPassword';

export default {
  title: 'ReenterPassword',
  component: ReenterPassword,
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

export const Basic = (props: Props) => {
  return <ReenterPassword {...props} />;
};
