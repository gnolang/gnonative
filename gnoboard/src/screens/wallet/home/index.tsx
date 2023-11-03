import React, { useEffect } from 'react';
import Layout from '@gno/components/pages';
import Button from '@gno/components/buttons';
import { useNavigation } from '@react-navigation/native';
import { RoutePath } from '@gno/router/path';
import { RouterWelcomeStackProp } from '@gno/router/custom-router';
import Text from '@gno/components/texts';
import styled from 'styled-components/native';
import CurrentAccount from '@gno/components/account/CurrentAccoutn';
import { useGno } from '@gno/hooks/use-gno';
import Loading from '@gno/screens/loading';
import { GnoAccount } from '@gno/native_modules/types';
import { QueryAccountResponse } from '@gno/api/gnomobiletypes_pb';
import { AccountBalance } from '@gno/components/account';
import { Spacer } from '@gno/components/row';

export const Home: React.FC = () => {
  const navigation = useNavigation<RouterWelcomeStackProp>();
  const gno = useGno();

  const [loading, setLoading] = React.useState<string | undefined>(undefined);
  const [account, setAccount] = React.useState<GnoAccount | undefined>(undefined);
  const [balance, setBalance] = React.useState<QueryAccountResponse | undefined>(undefined);
  const [unknownAddress, setUnknownAddress] = React.useState<boolean>(false);

  useEffect(() => {
    const unsubscribe = navigation.addListener('focus', async () => {
      setUnknownAddress(false);
      setAccount(undefined);
      setBalance(undefined);

      try {
        const response = await gno.getActiveAccount();
        setAccount(response.key);
        if (response.key) {
          const balance = await gno.queryAccount(response.key.address);
          setBalance(balance);
        }
      } catch (error: unknown | Error) {
        if (error?.rawMessage === 'ErrUnknownAddress(#110)') {
          setUnknownAddress(true);
        }
      } finally {
        setLoading(undefined);
      }
    });
    return unsubscribe;
  }, [navigation]);

  if (loading) {
    return <Loading message={loading} />;
  }

  return (
    <Layout.Container>
      <Layout.Body>
        <Text.Title>Gnomobile</Text.Title>
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
