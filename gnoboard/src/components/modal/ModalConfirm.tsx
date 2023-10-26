import ModalHeader from '@gno/components/modal/ModalHeader';
import Row from '@gno/components/row';
import { colors } from '@gno/styles';
import { StyleSheet, View } from 'react-native';
import { Modal } from 'react-native';
import Text from '../texts';
import Ruller from '../row/Ruller';
import Button from '../buttons';

export type Props = {
  title: string;
  message: string;
  visible: boolean;
  onClose: () => void;
  onConfirm: () => void;
};

const ModalConfirm = ({ visible, onClose, onConfirm, title, message }: Props) => {
  return (
    <Modal visible={visible} transparent={true} animationType='slide'>
      <View style={styles.container}>
        <View style={styles.modalView}>
          <ModalHeader title={title} onClose={onClose} />
          <View style={styles.message}>
            <Text.BodyMedium>{message}</Text.BodyMedium>
          </View>
          <Row>
            <Button title='Confirm' onPress={onConfirm} variant='primary-red' />
          </Row>
          <Ruller />
          <Row>
            <Button title='Close' onPress={onClose} variant='secondary' />
          </Row>
        </View>
      </View>
    </Modal>
  );
};

export default ModalConfirm;

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: colors.modal.backgroundOpaque,
    justifyContent: 'flex-end',
  },
  modalView: {
    backgroundColor: colors.modal.background,
    borderRadius: 16,
    padding: 32,
    paddingTop: 4,
    shadowColor: '#000',
    shadowOffset: {
      width: 0,
      height: 2,
    },
    shadowOpacity: 0.25,
    shadowRadius: 4,
    elevation: 5,
  },
  message: {
    marginBottom: 24,
  },
});
