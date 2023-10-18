import Layout from '@gno/components/pages';
import Text from '@gno/components/texts';
import { useGno } from '@gno/hooks/use-gno';
import { RouterWelcomeStack, RouterWelcomeStackProp } from '@gno/router/custom-router';
import { useNavigation } from '@react-navigation/native';
import { useState } from 'react';
import Button from '@gno/components/buttons';
import styled from 'styled-components/native';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import Alert from '@gno/components/alert';

export type Props = NativeStackScreenProps<RouterWelcomeStack, 'RemoveConfirm'>;

const RemoveConfirm = ({ route }: Props) => {
  const accountName = route.params.accountName;

  const gno = useGno();
  const navigation = useNavigation<RouterWelcomeStackProp>();
  const [error, setError] = useState<string | undefined>(undefined);

  const onConfirm = async () => {
    try {
      await gno.deleteAccount(accountName, undefined, true);
      navigation.navigate('WalletCreate');
    } catch (error) {
      setError(error?.toString());
    }
  };

  const onCancel = () => {
    navigation.goBack();
  };

  return (
    <Layout.Container>
      <Layout.Header />
      <Layout.Body>
        <Text.Title>Remove Account {accountName}</Text.Title>
        <Text.Body style={{ color: 'red' }}>
          Only proceed if you wish to remove this account from your wallet. You can always recover it with your seed phrase or your private
          key.
        </Text.Body>
        <Alert severity='error' message={error} />
        <HorizontalView>
          <Button title='Cancel' onPress={onCancel} />
          <Button title='Remove' onPress={onConfirm} />
        </HorizontalView>
      </Layout.Body>
    </Layout.Container>
  );
};

const HorizontalView = styled.View`
  flex-direction: row;
  margin-top: 24px;
  margin-left: 30px;
  margin-right: 30px;
  justify-content: space-between;
`;

export default RemoveConfirm;
