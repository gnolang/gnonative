import React, { useEffect } from 'react';
import Layout from '@gno/components/pages';
import Button from '@gno/components/buttons';
import { useNavigation } from '@react-navigation/native';
import { RoutePath } from '@gno/router/path';
import { RouterWelcomeStackProp } from '@gno/router/custom-router';
import Text from '@gno/components/texts';
import styled from 'styled-components/native';
import CurrentAccount from '@gno/components/account/CurrentAccoutn';
import Loading from '@gno/screens/loading';
import { AccountBalance } from '@gno/components/account';
import { Spacer } from '@gno/components/row';
import { ConnectError } from '@connectrpc/connect';
import { ErrCode, GRPCError, useGnoNativeContext, KeyInfo, QueryAccountResponse } from '@gnolang/gnonative';
import { useGnoboardContext } from '@gno/provider/gnoboard-provider';

export const Home: React.FC = () => {
  const navigation = useNavigation<RouterWelcomeStackProp>();
  const { gnonative } = useGnoNativeContext();

  const { account } = useGnoboardContext();
  const [loading, setLoading] = React.useState<string | undefined>(undefined);
  const [balance, setBalance] = React.useState<QueryAccountResponse | undefined>(undefined);
  const [unknownAddress, setUnknownAddress] = React.useState<boolean>(false);

  useEffect(() => {
    const unsubscribe = navigation.addListener('focus', async () => {
      setUnknownAddress(false);
      setBalance(undefined);

      try {
        if (account) {
          const balance = await gnonative.queryAccount(account.address);
          setBalance(balance);
        }
      } catch (error) {
        if (error instanceof ConnectError) {
          const err = new GRPCError(error);
          if (err.errCode() === ErrCode.ErrNoActiveAccount) {
            setUnknownAddress(true);
          }
        }
      } finally {
        setLoading(undefined);
      }
    });
    return unsubscribe;
  }, [navigation, account]);

  if (loading) {
    return <Loading message={loading} />;
  }

  return (
    <Layout.Container>
      <Layout.Body>
        <Text.Title>Gno Native Kit</Text.Title>
        <CurrentAccount account={account} />
        <AccountBalance accountInfo={balance?.accountInfo} unknownAddress={unknownAddress} />
        <ButtonGroup>
          <Button title='Create New Account' onPress={() => navigation.navigate(RoutePath.GenerateSeedPhrase)} variant='primary' />
          <Spacer />
          <Button title='Import Account' onPress={() => navigation.navigate(RoutePath.ImportPrivateKey)} variant='primary' />
          <Spacer />
          <Button title='Switch Accounts' onPress={() => navigation.navigate(RoutePath.SwitchAccounts)} variant='primary' />
          <Spacer />
          <Button title='Developer Mode' onPress={() => navigation.navigate(RoutePath.DevMode)} variant='primary' />
          <Spacer />
          <Button title='Change Network' onPress={() => navigation.navigate(RoutePath.ChangeNetwork)} variant='primary' />
          <Spacer />
          <Button title='Remove Account' onPress={() => navigation.navigate(RoutePath.RemoveAccount)} variant='secondary-red' />
        </ButtonGroup>
      </Layout.Body>
    </Layout.Container>
  );
};

const ButtonGroup = styled.View`
  margin-top: 32px;
`;

export default Home;
