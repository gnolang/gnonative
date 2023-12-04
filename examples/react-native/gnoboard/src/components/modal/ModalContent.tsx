import { colors } from '@gno/styles';
import { ViewProps } from 'react-native';
import styled from 'styled-components/native';

const ModalContainer = styled.View`
  flex: 1;
  background-color: ${colors.modal.backgroundOpaque};
  justify-content: flex-end;
`;

const ModalView = styled.View`
  background-color: ${colors.modal.background};
  border-radius: 16px;
  padding: 32px;
  padding-top: 4px;
  shadow-color: #000;
  shadow-opacity: 0.25;
  elevation: 5;
  shadom-offset: 0px 2px;
  shadow-radius: 4px;
`;

const ModalContent: React.FC<ViewProps> = (props: ViewProps) => (
  <ModalContainer>
    <ModalView>{props.children}</ModalView>
  </ModalContainer>
);

export default ModalContent;
