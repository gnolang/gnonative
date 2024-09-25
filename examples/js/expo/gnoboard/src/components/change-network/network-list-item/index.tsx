import Icons from '@gno/components/icons';
import Text from '@gno/components/texts';
import { NetworkMetainfo } from '@gno/types';
import { colors } from '@gno/styles';
import styled from 'styled-components/native';

export interface Props {
  currentChainId: string | undefined;
  networkMetainfo: NetworkMetainfo;
  onPress: (item: NetworkMetainfo) => void;
}

const NetworkListItem: React.FC<Props> = ({ networkMetainfo, currentChainId, onPress }: Props) => (
  <Row style={{ margin: 4 }} onPress={() => onPress(networkMetainfo)}>
    <LeftItens>
      <Text.BodyMedium style={{ color: colors.white }}>{networkMetainfo.chainName}</Text.BodyMedium>
      <Text.Caption1 style={{ color: colors.white }}>{networkMetainfo.gnoAddress}</Text.Caption1>
    </LeftItens>

    <RightItens>{currentChainId === networkMetainfo.chainId && <InUse />}</RightItens>
  </Row>
);

const InUse = () => (
  <>
    <Icons.CheckMark color={colors.white} />
    <Text.Caption1 style={{ paddingLeft: 8, color: colors.white }}>in use</Text.Caption1>
  </>
);

const Row = styled.TouchableOpacity`
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  background-color: ${colors.button.primary};
  height: auto;
  padding: 9px 16px;
  border-radius: 18px;
  transition: 0.2s;
`;

const LeftItens = styled.View`
  display: flex;
  flex-direction: column;
  align-items: flex-start;
`;

const RightItens = styled.View`
  display: flex;
  flex-direction: row;
  align-items: center;
`;

export default NetworkListItem;
