import React from 'react';
import { TouchableOpacity, View } from 'react-native';
import Icons from '../icons';
import { useNavigation } from '@react-navigation/native';
import styled from 'styled-components/native';

type Props = {
  iconType?: 'close' | 'back';
  onCloseHandler?: () => void;
  title?: string;
  subtitle?: string;
};

const Header: React.FC<Props> = ({ iconType = 'close', onCloseHandler, title = '', subtitle = '' }) => {
  const navigate = useNavigation();

  if (!onCloseHandler) {
    onCloseHandler = () => {
      navigate.goBack();
    };
  }

  return (
    <Wrapper>
      <TouchableOpacity onPress={onCloseHandler}>{iconType === 'close' ? <Icons.Close /> : <Icons.ArrowLeft />}</TouchableOpacity>
      <View>
        <TitleText>{title}</TitleText>
        <SubtitleText>{subtitle}</SubtitleText>
      </View>
    </Wrapper>
  );
};

const SubtitleText = styled.Text`
  font-size: 12px;
  color: #000000;
`;

const TitleText = styled.Text`
  font-size: 16px;
  font-weight: bold;
  color: #000000;
`;

const Wrapper = styled.View`
  width: 100%;
  height: 64px;
  padding-left: 2px;
  padding-top: 10px;
`;

export default Header;
