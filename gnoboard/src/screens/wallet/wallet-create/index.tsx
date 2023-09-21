import React, { useEffect } from 'react';
import Layout from '@gno/components/pages';
import Button from '@gno/components/buttons';
import { useNavigation } from '@react-navigation/native';
import { RoutePath } from '@gno/router/path';
import { RouterWelcomeStackProp } from '@gno/router/custom-router';
import Text from '@gno/components/texts';
import { useGno } from '@gno/hooks/use-gno';

export const WalletCreate: React.FC = () => {
  const gno = useGno();
  const navigation = useNavigation<RouterWelcomeStackProp>();

  // useEffect(() => {
  //   const unsubscribe = navigation.addListener('focus', async () => {
  //     try {
  //       const info = await gno.listKeyInfo();
  //       console.log('accounts ok keybase:', info);
  //     } catch (error) {
  //       console.log(error);
  //     }
  //   });
  //   return unsubscribe;
  // }, [navigation]);

  return (
    <Layout.Container>
      <Layout.Body>
        <Text.Title>Gnomobile</Text.Title>
        <Button title='Create New Wallet' onPress={() => navigation.navigate(RoutePath.GenerateSeedPhrase)} />
        <Button title='Import Wallet' onPress={() => navigation.navigate(RoutePath.ImportPrivateKey)} />
        <Button title='Switch Accounts' onPress={() => navigation.navigate(RoutePath.SwitchAccounts)} />
        <Button title='Developer Mode' onPress={() => navigation.navigate(RoutePath.DevMode)} />
      </Layout.Body>
    </Layout.Container>
  );
};

export default WalletCreate;
