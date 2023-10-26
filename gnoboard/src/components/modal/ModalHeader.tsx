import React from 'react';
import { TouchableOpacity, View } from 'react-native';
import styled from 'styled-components';
import Text from '../texts';
import Icons from '../icons';

export type Props = {
  subtitle?: string;
  title?: string;
  onClose: () => void;
  iconType?: 'close' | 'arrowLeft';
};

function ModalHeader(props: Props) {
  const { title, subtitle, iconType = 'close', onClose } = props;

  const saveAndClose = () => {
    onClose();
  };

  return (
    <Content>
      <TouchableStyled onPress={saveAndClose}>
        {iconType === 'close' ? <Icons.Close color='#667386' /> : <Icons.ArrowLeft />}
      </TouchableStyled>
      <View>
        <Text.HeaderTitleText>{title}</Text.HeaderTitleText>
        <Text.HeaderSubtitleText>{subtitle}</Text.HeaderSubtitleText>
      </View>
    </Content>
  );
}

const Content = styled(View)`
  height: 76px;
  padding-top: 18px;
`;

const TouchableStyled = styled(TouchableOpacity)`
  position: absolute;
  padding-top: 18px;
  padding-left: 12px;
  z-index: 1;
`;

export default ModalHeader;
