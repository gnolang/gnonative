import { View } from 'react-native';
import React from 'react';
import Button, { Props } from '.';

export default {
  title: 'Button',
  component: Button,
  decorators: [
    (Story: React.FC) => (
      <View style={{ width: '100%', height: '100%', padding:24 }}>
        <Story />
      </View>
    ),
  ],
  args: {
    title: 'Button label',
  },
};

export const Basic = (props: Props) => {
  return <Button {...props} />;
};
