import { Spacer } from '@gno/components/row';
import { Modal as NativeModal } from 'react-native';
import Text from '../texts';
import Ruller from '../row/Ruller';
import Button from '../buttons';
import ModalHeader from './ModalHeader';
import ModalContent from './ModalContent';

export type Props = {
  title: string;
  message: string;
  visible: boolean;
  onClose: () => void;
  onConfirm: () => void;
};

const ModalConfirm = ({ visible, onClose, onConfirm, title, message }: Props) => {
  return (
    <NativeModal visible={visible} transparent={true} animationType='slide'>
      <ModalContent>
        <ModalHeader title={title} onClose={onClose} />
        <Text.BodyMedium>{message}</Text.BodyMedium>
        <Spacer />
        <Button title='Confirm' onPress={onConfirm} variant='primary-red' />
        <Ruller />
        <Button title='Close' onPress={onClose} variant='secondary' />
      </ModalContent>
    </NativeModal>
  );
};

export default ModalConfirm;
