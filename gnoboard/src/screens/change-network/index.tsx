import Layout from '@gno/components/pages';
import Text from '@gno/components/texts';
import { RouterWelcomeStackProp } from '@gno/router/custom-router';
import { useNavigation } from '@react-navigation/native';
import { useEffect, useState } from 'react';
import Loading from '../loading';
import NetworkList from '@gno/components/change-network/network-list';
import chains from '@gno/resources/chains/chains.json';
import { useGno } from '@gno/hooks/use-gno';
import { RoutePath } from '@gno/router/path';
import { NetworkMetainfo } from '@gno/native_modules/types';

const ChangeNetwork = () => {
  const gno = useGno();
  const navigation = useNavigation<RouterWelcomeStackProp>();
  const [loading, setLoading] = useState<string | undefined>(undefined);
  const [currentChainId, setCurrentChainId] = useState<string | undefined>(undefined);
  const [currentRemote, setCurrentRemote] = useState<string | undefined>(undefined);

  useEffect(() => {
    const unsubscribe = navigation.addListener('focus', async () => {
      try {
        setCurrentChainId(undefined);
        setCurrentRemote(undefined);
        setLoading('Loading network...');
        const chainId = await gno.getChainID();
        const remote = await gno.getRemote();
        setCurrentChainId(chainId);
        setCurrentRemote(remote);
        setLoading(undefined);
      } catch (error: unknown | Error) {
        setLoading(error?.toString());
        console.log(error);
      }
    });
    return unsubscribe;
  }, [navigation]);

  const onNetworkChange = async (networkMetainfo: NetworkMetainfo) => {
    try {
      setLoading('Changing network...');
      await gno.setChainID(networkMetainfo.chainId);
      await gno.setRemote(networkMetainfo.gnoUrl);
      setLoading(undefined);
      navigation.navigate(RoutePath.Home);
    } catch (error: unknown | Error) {
      setLoading(error?.toString());
      console.log(error);
    }
  };

  if (loading) return <Loading message={loading} />;

  return (
    <>
      <Layout.Container>
        <Layout.Header />
        <Layout.Body>
          <Text.Title>Change Network</Text.Title>
          <Text.Subheadline>Current Network: {currentChainId}</Text.Subheadline>
          <Text.Subheadline>{currentRemote}</Text.Subheadline>
          <NetworkList currentChainId={currentChainId} networkMetainfos={chains} onNetworkChange={onNetworkChange} />
        </Layout.Body>
      </Layout.Container>
    </>
  );
};

export default ChangeNetwork;
