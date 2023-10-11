import { View } from 'react-native';
import React from 'react';
import Input, { Props } from '.';

export default {
  title: 'Input',
  component: Input,
  decorators: [
    (Story: React.FC) => (
      <View style={{ width: '100%', height: '100%', padding: 24 }}>
        <Story />
      </View>
    ),
  ],
  args: {
    title: 'some text',
  },
};

export const Basic = (props: Props) => {
  return <Input {...props} />;
};

export const Error = ({ error = 'error', ...rest }: Props) => {
  return <Input error={error} {...rest} />;
};
