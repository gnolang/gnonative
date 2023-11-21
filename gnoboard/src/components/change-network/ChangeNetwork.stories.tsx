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
  return (
    <>
      <NetworkList networkMetainfos={chains} currentNetworkId='test3' />
    </>
  );
};

export const ListItem = ({ networkMetainfo }: Props) => {
  return (
    <>
      <Spacer />
      <NetworkListItem networkMetainfo={networkMetainfo} currentNetworkId={undefined} />
      <Spacer />
      <NetworkListItem networkMetainfo={networkMetainfo} currentNetworkId={'test3'} />
    </>
  );
};
