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
import ModalConfirm from '../../../components/modal/ModalConfirm';
import { Spacer } from '@gno/components/row';

const walletContent = {
  title: 'Import with Seed Phrase',
  desc: 'Import an existing wallet with a 12 or 24-word seed phrase.',
  terms: 'This phrase will only be stored on this device. Adena can’t recover it for you.',
};

const EnterSeedPhrase = () => {
  const [recoveryPhrase, setRecoveryPhrase] = useState('');
  const [name, setName] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [error, setError] = useState<string | undefined>(undefined);
  const [showModal, setShowModal] = useState(false);
  const gno = useGno();
  const navigation = useNavigation<RouterWelcomeStackProp>();

  const recoverAccount = async (override: boolean = false) => {
    if (!recoveryPhrase || !name || !password || !confirmPassword) return;

    if (password !== confirmPassword) {
      setError('Passwords do not match.');
      console.log('password and confirmPassword are not the same');
      return;
    }

    try {
      if (!override) {
        const hasKeyByName = await gno.hasKeyByName(name);
        if (hasKeyByName) {
          setShowModal(true);
          return;
        }
      }

      const response = await gno.createAccount(name, recoveryPhrase, password);
      await gno.selectAccount(name);
      await gno.setPassword(password);
      console.log('createAccount response: ' + response);
      navigation.navigate(RoutePath.Home);
    } catch (error) {
      setError(JSON.stringify(error));
      console.log('create account error: ', JSON.stringify(error));
    }
  };

  return (
    <>
      <Layout.Container>
        <Layout.Header />
        <Layout.Body>
          <Text.Title>{walletContent.title}</Text.Title>
          <Text.Body>{walletContent.desc}</Text.Body>
          <Alert severity='error' message={error} />
          <SeedBox placeholder='Enter your seed phrase' value={recoveryPhrase} onChangeText={(value) => setRecoveryPhrase(value.trim())} />
          <TextInput placeholder='Account Name' value={name} onChangeText={setName} autoCapitalize='none' />
          <TextInput placeholder='Password' value={password} onChangeText={setPassword} secureTextEntry={true} error={error} />
          <TextInput
            placeholder='Confirm Password'
            value={confirmPassword}
            onChangeText={setConfirmPassword}
            secureTextEntry={true}
            error={error}
          />
          <Spacer />
          <Button title='Import' onPress={() => recoverAccount()} variant='primary' />
        </Layout.Body>
      </Layout.Container>
      <ModalConfirm
        title={`Overriding ${name}'s Account`}
        message='This account name is already in use. Are you sure you want to override the existing account?'
        visible={showModal}
        onClose={() => setShowModal(false)}
        onConfirm={() => recoverAccount(true)}
      />
    </>
  );
};

export default EnterSeedPhrase;
