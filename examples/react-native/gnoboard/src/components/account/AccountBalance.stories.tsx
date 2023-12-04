import { View } from 'react-native';
import React from 'react';
import AccountBalance, { Props } from './AccountBalance';

export default {
  title: 'AccountBalance',
  component: AccountBalance,
  decorators: [
    (Story: React.FC) => (
      <View style={{ width: '100%', height: '100%', padding: 24 }}>
        <Story />
      </View>
    ),
  ],
  args: {
    accountInfo: {
      coins: [
        { denom: 'GNOT', amount: '0.0002324' },
        { denom: 'GNO', amount: '0.0002324' },
        { denom: 'BTC', amount: '0.0302324112' },
      ],
    },
  },
};

export const Basic = (props: Props) => {
  return <AccountBalance {...props} />;
};
