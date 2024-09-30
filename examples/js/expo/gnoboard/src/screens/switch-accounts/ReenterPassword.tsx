import Alert from '@gno/components/alert';
import Button from '@gno/components/buttons';
import { Modal } from '@gno/components/modal';
import { Spacer } from '@gno/components/row';
import TextInput from '@gno/components/textinput';
import Text from '@gno/components/texts';
import { ErrCode, GRPCError, useGnoNativeContext } from '@gnolang/gnonative';
import { useState } from 'react';
import { Modal as NativeModal } from 'react-native';
import { ConnectError } from '@connectrpc/connect';

export type Props = {
  visible: boolean;
  accountName: string;
  onClose: (sucess: boolean) => void;
};

const ReenterPassword = ({ visible, accountName, onClose }: Props) => {
  const { gnonative } = useGnoNativeContext();
  const [password, setPassword] = useState('');
  const [error, setError] = useState<string | undefined>(undefined);

  const onConfirm = async () => {
    if (!password) return;

    try {
      setError(undefined);
      await gnonative.setPassword(password);
      onClose(true);
    } catch (error) {
      if (error instanceof ConnectError) {
        const err = new GRPCError(error);
        if (err.errCode() === ErrCode.ErrDecryptionFailed) {
          setError('Wrong password, please try again.');
          return;
        }
      }
      setError(JSON.stringify(error));
    }
  };

  return (
    <NativeModal visible={visible} transparent={true} animationType='slide'>
      <Modal.Content>
        <Modal.Header title='Re-enter your password' onClose={() => onClose(false)} />
        <Text.BodyMedium>Please, reenter the password for the selected account.</Text.BodyMedium>
        <Spacer />
        <TextInput placeholder={`Password for ${accountName}'s Account`} error={error} secureTextEntry={true} onChangeText={setPassword} />
        <Spacer />
        <Alert severity='error' message={error} />
        <Spacer />
        <Button title='Confirm' onPress={onConfirm} variant='primary' />
      </Modal.Content>
    </NativeModal>
  );
};

export default ReenterPassword;
