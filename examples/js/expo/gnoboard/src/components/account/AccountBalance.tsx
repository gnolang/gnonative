import Text from '../texts';
import styled from 'styled-components/native';
import { colors } from '@gno/styles';
import { BaseAccount } from '@gnolang/gnonative';
import Row from '../row';

export type Props = {
  accountInfo: BaseAccount | undefined;
  unknownAddress: boolean;
};

const AccountBalance = (props: Props) => {
  if (props.unknownAddress) {
    return (
      <CenterView>
        <Row>
          <Text.Body>Your account is not known on the blockchain.</Text.Body>
        </Row>
      </CenterView>
    );
  }

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
    <Row>
      <Text.Body style={{ paddingRight: 4 }}>{props.denom}</Text.Body>
      <Text.Body>{props.amount.toString()}</Text.Body>
    </Row>
  );
};

const CenterView = styled.View`
  background-color: ${colors.grayscale['200']};
  padding: 8px 8px 16px 8px;
  border-radius: 8px;
  border: 1px solid ${colors.grayscale['700']};
  justify-content: center;
  align-items: center;
`;

export default AccountBalance;
