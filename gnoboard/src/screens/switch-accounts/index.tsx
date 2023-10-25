import Layout from '@gno/components/pages';
import Text from '@gno/components/texts';
import { useGno } from '@gno/hooks/use-gno';
import { RouterWelcomeStackProp } from '@gno/router/custom-router';
import { useNavigation } from '@react-navigation/native';
import { useEffect, useState } from 'react';
import Loading from '../loading';
import SideMenuAccountList from '@gno/components/common/side-menu-account-list/side-menu-account-list';
import { GnoAccount } from '@gno/native_modules/types';
import { RoutePath } from '@gno/router/path';

const SwitchAccounts = () => {
  const gno = useGno();
  const navigation = useNavigation<RouterWelcomeStackProp>();
  const [loading, setLoading] = useState<string | undefined>(undefined);
  const [accounts, setAccounts] = useState<GnoAccount[]>([]);

  useEffect(() => {
    const unsubscribe = navigation.addListener('focus', async () => {
      try {
        setLoading('Loading accounts...');
        const response = await gno.listKeyInfo();

        setAccounts(response);
        setLoading(undefined);
      } catch (error: unknown | Error) {
        setLoading(error?.toString());
        console.log(error);
      }
    });
    return unsubscribe;
  }, [navigation]);

  const onChangeAccountHandler = async (value: GnoAccount) => {
    try {
      setLoading('Changing account...');
      const response = await gno.selectAccount(value.name);
      // TODO: If !response.hasPassword then have the user enter the password.
      setLoading(undefined);
      navigation.navigate(RoutePath.Home);
    } catch (error: unknown | Error) {
      setLoading(error?.toString());
      console.log(error);
    }
  };

  if (loading) return <Loading message={loading} />;

  return (
    <Layout.Container>
      <Layout.Header />
      <Layout.Body>
        <Text.Title>Switch Accounts</Text.Title>
        {accounts.length === 0 ? <Text.Body>No accounts found</Text.Body> : null}
        <SideMenuAccountList accounts={accounts} changeAccount={onChangeAccountHandler} />
      </Layout.Body>
    </Layout.Container>
  );
};

export default SwitchAccounts;
