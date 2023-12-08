import Text from '../texts';
import { GnoAccount } from '@gno/native_modules/types';
import styled from 'styled-components/native';

export type Props = {
  account: GnoAccount | undefined;
};

const CurrentAccount = ({ account }: Props) => {
  if (!account)
    return (
      <CenterView>
        <Text.HeaderTitle>No Account Selected</Text.HeaderTitle>
      </CenterView>
    );

  return (
    <CenterView>
      <Text.HeaderTitle>{account.name}</Text.HeaderTitle>
    </CenterView>
  );
};

const CenterView = styled.View`
  align-items: center;
  margin-top: 16px;
  margin-bottom: 16px;
`;

export default CurrentAccount;
