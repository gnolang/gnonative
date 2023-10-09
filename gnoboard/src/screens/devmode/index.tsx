import { ConsoleView } from '@gno/components/consoleview';
import TextInput from '@gno/components/textinput';
import { GoBridge } from '@gno/native_modules';
import { screenStyleSheet as styles } from '@gno/styles';
import { useState } from 'react';
import { Linking, ScrollView, StyleSheet, Text, View } from 'react-native';
import { getTcpPort } from '@gno/utils/bridge';
import Button from '@gno/components/buttons';
import Layout from '@gno/components/pages';
import { createClient } from '@gno/grpc/client';
import { ListKeyInfoRequest } from '@gno/api/rpc_pb';

function DevMode() {
  const [postContent, setPostContent] = useState('');
  const [appConsole, setAppConsole] = useState<string>('');
  const [loading, setLoading] = useState<string | undefined>(undefined);

  const onPostPress = async () => {
    setLoading('Replying to a post...');
    setAppConsole('replying to a post...');
    const gasFee = '1000000ugnot';
    const gasWanted = 2000000;
    const args: Array<string> = ['2', '1', '1', postContent];
    GoBridge.call('gno.land/r/demo/boards', 'CreateReply', args, gasFee, gasWanted)
      .then((data) => {
        setAppConsole(data);
        setPostContent('');
      })
      .catch((err) => {
        setAppConsole(err);
      })
      .finally(() => setLoading(undefined));
  };

  const onLoadAccountPress = async () => {
    setLoading('Loading account...');
    setAppConsole('Loading account...');

    try {
      const port = await getTcpPort();
      console.log('port: ', port);
      const client = await createClient(port);

      const response = await client.listKeyInfo(new ListKeyInfoRequest())
      console.log('response: ', response);
      setAppConsole(JSON.stringify(response));
      console.log(response);
    } catch (error) {
      console.log(error)
      setAppConsole('error' + JSON.stringify(error));
    } finally {
      setAppConsole('done');
    }
  };

  const loadInBrowser = () => {
    Linking.openURL('http://testnet.gno.berty.io/r/demo/boards:gnomobile/1').catch((err) => console.error("Couldn't load page", err));
  };

  return (
    <Layout.Container>
      <Layout.Header />
      <Layout.Body>
        <ScrollView contentContainerStyle={styles.scrollViewContent}>
          <Button title='Load Accounts' onPress={onLoadAccountPress} />
          <Text>Content:</Text>
          <View style={customStyles.sendGroupLikeWhatsapp}>
            <TextInput style={customStyles.inputMsg} value={postContent} onChangeText={setPostContent} />
            <Button title='Send' onPress={onPostPress} />
          </View>
          <ConsoleView text={appConsole} />
          <Button title='Open http://testnet.gno.berty.io/r/demo/boards:gnomobile/1' onPress={loadInBrowser} />
        </ScrollView>
      </Layout.Body>
    </Layout.Container>
  );
}

const customStyles = StyleSheet.create({
  sendGroupLikeWhatsapp: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
  },
  inputMsg: { width: '80%' },
});

export default DevMode;
