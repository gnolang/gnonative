import { ConsoleView } from '@gno/components/consoleview';
import TextInput from '@gno/components/textinput';
import { screenStyleSheet as styles } from '@gno/styles';
import { useState } from 'react';
import { Linking, ScrollView, StyleSheet, Text, View } from 'react-native';
import Button from '@gno/components/buttons';
import Layout from '@gno/components/pages';
import { useGno } from '@gno/hooks/use-gno';
import { Buffer } from 'buffer';
import ReenterPassword from '../switch-accounts/ReenterPassword';

function DevMode() {
  const [postContent, setPostContent] = useState('');
  const [appConsole, setAppConsole] = useState<string>('');
  const [loading, setLoading] = useState<string | undefined>(undefined);
  const [reenterPassword, setReenterPassword] = useState<string | undefined>(undefined);

  const gno = useGno();

  const onPostPress = async () => {
    setLoading('Replying to a post...');
    setAppConsole('replying to a post...');
    const gasFee = '1000000ugnot';
    const gasWanted = 2000000;
    const args: Array<string> = ['2', '1', '1', postContent];
    try {
      const response = await gno.call('gno.land/r/demo/boards', 'CreateReply', args, gasFee, gasWanted);
      console.log('response: ', response);
      setAppConsole(Buffer.from(response.result).toString());
    } catch (error) {
      if (error?.rawMessage === 'ErrDecryptionFailed(#109)') {
        const account = await gno.getActiveAccount()
        setReenterPassword(account.key?.name);
        return;
      }
      console.log(error);
      setAppConsole('error' + JSON.stringify(error));
    } finally {
      setLoading(undefined);
    }
  };

  const loadInBrowser = () => {
    Linking.openURL('http://testnet.gno.berty.io/r/demo/boards:gnomobile/1').catch((err) => console.error("Couldn't load page", err));
  };

  const onCloseReenterPassword = async () => {
    setReenterPassword(undefined);
  };

  return (
    <>
      <Layout.Container>
        <Layout.Header />
        <Layout.Body>
          <ScrollView contentContainerStyle={styles.scrollViewContent}>
            <Text>Content:</Text>
            <View style={customStyles.sendGroupLikeWhatsapp}>
              <TextInput style={customStyles.inputMsg} value={postContent} onChangeText={setPostContent} />
              <Button title='Send' onPress={onPostPress} variant='primary' style={{ width: '30%' }} />
            </View>
            <ConsoleView text={appConsole} />
            <Button title='Open http://testnet.gno.berty.io/r/demo/boards:gnomobile/1' onPress={loadInBrowser} variant='primary' />
          </ScrollView>
        </Layout.Body>
      </Layout.Container>
      {reenterPassword ? (
        <ReenterPassword visible={Boolean(reenterPassword)} accountName={reenterPassword} onClose={onCloseReenterPassword} />
      ) : null}
    </>
  );
}

const customStyles = StyleSheet.create({
  sendGroupLikeWhatsapp: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
  },
  inputMsg: { width: '70%' },
});

export default DevMode;
