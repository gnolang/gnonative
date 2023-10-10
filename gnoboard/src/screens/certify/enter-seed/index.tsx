import React, { useState } from 'react';
import Layout from '@gno/components/pages';
import Text from '@gno/components/texts';
import Button from '@gno/components/buttons';
import { useGno } from '@gno/hooks/use-gno';
import { useNavigation } from '@react-navigation/native';
import { RouterWelcomeStackProp } from '@gno/router/custom-router';
import { RoutePath } from '@gno/router/path';
import SeedBox from '@gno/components/seedbox';
import TextInput from '@gno/components/textinput';
import Alert from '@gno/components/alert';

const walletContent = {
  title: 'Import with Seed Phrase',
  desc: 'Import an existing wallet with a 12 or 24-word seed phrase.',
  terms: 'This phrase will only be stored on this device. Adena can’t recover it for you.',
};

const EnterSeedPharse = () => {
  const [recoveryPhrase, setRecoveryPhrase] = useState('');
  const [name, setName] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [error, setError] = useState<string | undefined>(undefined);

  const gno = useGno();
  const navigation = useNavigation<RouterWelcomeStackProp>();

  const nextHandler = async () => {
    if (!recoveryPhrase || !name || !password || !confirmPassword) return;

    if (password !== confirmPassword) {
      setError('Passwords do not match.');
      console.log('password and confirmPassword are not the same');
      return;
    }

    try {
      const response = await gno.createAccount(name, recoveryPhrase, password);
      await gno.selectAccount(name);
      await gno.setPassword(password);
      console.log('createAccount response: ' + response);
      navigation.navigate(RoutePath.WalletCreate);
    } catch (error) {
      console.log('createAccount error: ', error);
    }
  };

  return (
    <Layout.Container>
      <Layout.Header />
      <Layout.Body>
        <Text.Title>{walletContent.title}</Text.Title>
        <Text.Body>{walletContent.desc}</Text.Body>
        <Alert severity='error' message={error} />
        <SeedBox placeholder='Enter your seed phrase' value={recoveryPhrase} onChangeText={(value) => setRecoveryPhrase(value.trim())} />
        <TextInput placeholder='Account Name' value={name} onChangeText={setName} />
        <TextInput placeholder='Password' value={password} onChangeText={setPassword} secureTextEntry={true} error={error} />
        <TextInput placeholder='Confirm Password' value={confirmPassword} onChangeText={setConfirmPassword} secureTextEntry={true} error={error} />
        <Button title='Import' onPress={nextHandler} />
      </Layout.Body>
    </Layout.Container>
  );
};

export default EnterSeedPharse;
