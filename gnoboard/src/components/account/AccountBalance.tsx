import Text from '../texts';
import styled from 'styled-components/native';
import { colors } from '@gno/styles';
import { BaseAccount } from '@gno/api/gnomobiletypes_pb';
import { View } from 'react-native';

export type Props = {
  accountInfo: BaseAccount | undefined;
};

const AccountBalance = (props: Props) => {
  if (!props.accountInfo) {
    return null;
  }

  const coins = props.accountInfo.coins;

  return (
    <CenterView>
      {coins.map((coin, index) => (
        <Coin key={index} amount={coin.amount} denom={coin.denom} />
      ))}
    </CenterView>
  );
};

const Coin = (props: { amount: bigint; denom: string }) => {
  return (
    <View>
      <Text.Body style={{ paddingRight: 4 }}>{props.denom}</Text.Body>
      <Text.Body>{props.amount.toString()}</Text.Body>
    </View>
  );
};

const CenterView = styled.View`
  background-color: ${colors.grayscale['200']};
  padding: 8px;
  border-radius: 8px;
  border: 1px solid ${colors.grayscale['700']};
  justify-content: center;
  align-items: center;
`;

export default AccountBalance;
