import { View } from 'react-native';
import React from 'react';
import CurrentAccount, { Props } from './CurrentAccoutn';

export default {
  title: 'CurrentAccount',
  component: CurrentAccount,
  decorators: [
    (Story: React.FC) => (
      <View style={{ width: '100%', height: '100%', padding: 24, backgroundColor:'red' }}>
        <Story />
      </View>
    ),
  ],
  args: {
    account: {
      name: 'Account 1',
    }
  },
};

export const Basic = (props: Props) => {
  return <CurrentAccount {...props} />;
};
