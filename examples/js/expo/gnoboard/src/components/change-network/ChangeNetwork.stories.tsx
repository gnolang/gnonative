import { View } from 'react-native';
import NetworkListItem, { Props } from './network-list-item';
import { Spacer } from '../row';
import NetworkList from './network-list';
import chains from '@gno/resources/chains/chains.json';

export default {
  title: 'ChangNetwork',
  component: NetworkListItem,
  decorators: [
    (Story: React.FC) => (
      <View style={{ width: '100%', height: '100%' }}>
        <Story />
      </View>
    ),
  ],
  args: {
    networkMetainfo: {
      networkName: 'test3',
      gnoAddress: 'https://test3.gno.land',
    },
  },
};

export const Basic = () => {
  // This callback isn't used by the storybook
  const onNetworkChange = async () => {}
  return (
    <>
      <NetworkList networkMetainfos={chains} currentChainId='test3' onNetworkChange={onNetworkChange} />
    </>
  );
};

export const ListItem = ({ networkMetainfo }: Props) => {
  // This callback isn't used by the storybook
  const onPress = async () => {}
  return (
    <>
      <Spacer />
      <NetworkListItem networkMetainfo={networkMetainfo} currentChainId={undefined} onPress={onPress} />
      <Spacer />
      <NetworkListItem networkMetainfo={networkMetainfo} currentChainId={'test3'} onPress={onPress} />
    </>
  );
};
