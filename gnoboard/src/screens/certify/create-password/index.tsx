import React, { useState } from 'react';
import Layout from '@gno/components/pages';
import Text from '@gno/components/texts';
import TextInput from '@gno/components/textinput';
import Button from '@gno/components/buttons';
import { useGno } from '@gno/hooks/use-gno';
import { RouterWelcomeStack, RouterWelcomeStackProp } from '@gno/router/custom-router';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { useNavigation } from '@react-navigation/native';
import { RoutePath } from '@gno/router/path';
import SeedBox from '@gno/components/seedbox';
import Alert from '@gno/components/alert';
import { Spacer } from '@gno/components/row';
import { ModalConfirm } from '@gno/components/modal';

const text = {
  title: 'Create\na Password',
  desc: 'This will be used to unlock your wallet.',
};

export type Props = NativeStackScreenProps<RouterWelcomeStack, 'CreatePassword'>;

const CreatePassword: React.FC<Props> = ({ route }) => {
  const phrase = route.params.phrase;
  const [name, setName] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [error, setError] = useState<string | undefined>(undefined);
  const [showModal, setShowModal] = useState(false);

  const gno = useGno();
  const navigation = useNavigation<RouterWelcomeStackProp>();

  const createAccount = async (override: boolean = false) => {
    if (!name || !password || !confirmPassword) return;

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

      const response = await gno.createAccount(name, phrase, password);
      console.log('createAccount response: ' + response);
      await gno.selectAccount(name);
      await gno.setPassword(password);
      navigation.navigate(RoutePath.Home);
    } catch (error) {
      setError(JSON.stringify(error));
      console.log(JSON.stringify(error));
    }
  };

  return (
    <>
      <Layout.Container>
        <Layout.Header />
        <Layout.Body>
          <Text.Title>{text.title}</Text.Title>
          <Text.Body>{text.desc}</Text.Body>
          <Alert severity='error' message={error} />
          <TextInput placeholder='Account Name' value={name} onChangeText={setName} autoCapitalize='none' />
          <TextInput placeholder='Password' value={password} onChangeText={setPassword} secureTextEntry={true} error={error} />
          <TextInput
            placeholder='Confirm Password'
            value={confirmPassword}
            onChangeText={setConfirmPassword}
            secureTextEntry={true}
            error={error}
          />
          <SeedBox placeholder='Your seed phrase' value={phrase} editable={false} />
          <Spacer />
          <Button title='Save' onPress={() => createAccount()} variant='primary' />
        </Layout.Body>
      </Layout.Container>
      <ModalConfirm
        title={`Overriding ${name}'s Account`}
        message='This account name is already in use. Are you sure you want to override the existing account?'
        visible={showModal}
        onClose={() => setShowModal(false)}
        onConfirm={() => createAccount(true)}
      />
    </>
  );
};

export default CreatePassword;
