import { View } from 'react-native';
import React from 'react';
import ModalHeader, { Props } from './ModalHeader';

export default {
  title: 'Modal',
  component: ModalHeader,
  argTypes: {
    iconType: {
      options: ['arrowLeft', 'close'],
      control: { type: 'radio' },
    },
  },
  decorators: [
    (Story: React.FC) => (
      <View style={{ width: '100%', height: '100%' }}>
        <Story />
      </View>
    ),
  ],
  args: {
    title: 'Title goes here',
    subtitle: 'Subtile goes here',
  },
};

export const Header = ({ title = 'Confirm Override', subtitle = 'Override something', ...rest }: Props) => {
  return <ModalHeader {...rest} onClose={() => {}} title={title} subtitle={subtitle} />;
};
