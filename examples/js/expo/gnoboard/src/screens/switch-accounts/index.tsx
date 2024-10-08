import Layout from '@gno/components/pages';
import Text from '@gno/components/texts';
import { KeyInfo, useGnoNativeContext } from '@gnolang/gnonative';
import { RouterWelcomeStackProp } from '@gno/router/custom-router';
import { useNavigation } from '@react-navigation/native';
import { useEffect, useState } from 'react';
import { useGnoboardContext } from '@gno/provider/gnoboard-provider';
import Loading from '../loading';
import SideMenuAccountList from '@gno/components/common/side-menu-account-list/side-menu-account-list';
import { RoutePath } from '@gno/router/path';
import ReenterPassword from './ReenterPassword';

const SwitchAccounts = () => {
  const { gnonative } = useGnoNativeContext();
  const navigation = useNavigation<RouterWelcomeStackProp>();
  const [loading, setLoading] = useState<string | undefined>(undefined);
  const [accounts, setAccounts] = useState<KeyInfo[]>([]);
  const [reenterPassword, setReenterPassword] = useState<string | undefined>(undefined);
  const { setAccount } = useGnoboardContext();

  useEffect(() => {
    const unsubscribe = navigation.addListener('focus', async () => {
      try {
        setLoading('Loading accounts...');
        const response = await gnonative.listKeyInfo();

        setAccounts(response);
        setLoading(undefined);
      } catch (error: unknown | Error) {
        setLoading(error?.toString());
        console.log(error);
      }
    });
    return unsubscribe;
  }, [navigation]);

  const onChangeAccountHandler = async (value: KeyInfo) => {
    try {
      setLoading('Changing account...');
      const response = await gnonative.activateAccount(value.name);
      setAccount(value);
      setLoading(undefined);
      if (!response.hasPassword) {
        setReenterPassword(value.name);
        return;
      }
      navigation.navigate(RoutePath.Home);
    } catch (error: unknown | Error) {
      setLoading(error?.toString());
      console.log(error);
    }
  };

  const onCloseReenterPassword = async (sucess: boolean) => {
    setReenterPassword(undefined);
    if (sucess) {
      navigation.navigate(RoutePath.Home);
    }
  };

  if (loading) return <Loading message={loading} />;

  return (
    <>
      <Layout.Container>
        <Layout.Header />
        <Layout.Body>
          <Text.Title>Switch Accounts</Text.Title>
          {accounts.length === 0 ? <Text.Body>No accounts found</Text.Body> : null}
          <SideMenuAccountList accounts={accounts} changeAccount={onChangeAccountHandler} />
        </Layout.Body>
      </Layout.Container>
      {reenterPassword ? (
        <ReenterPassword visible={Boolean(reenterPassword)} accountName={reenterPassword} onClose={onCloseReenterPassword} />
      ) : null}
    </>
  );
};

export default SwitchAccounts;
