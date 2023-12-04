import Button from '@gno/components/buttons';
import Layout from '@gno/components/pages';
import { Spacer } from '@gno/components/row';
import SeedBox from '@gno/components/seedbox';
import Text from '@gno/components/texts';
import { useGno } from '@gno/hooks/use-gno';
import { RouterWelcomeStackProp } from '@gno/router/custom-router';
import { RoutePath } from '@gno/router/path';
import { useNavigation } from '@react-navigation/native';
import { useState } from 'react';

const text = {
  title: 'Seed Phrase',
  desc: 'This phrase is the only way to recover this wallet. DO NOT share it with anyone.',
  termsA: 'This phrase will only be stored on this device. Adena can’t recover it for you.',
  termsB: 'I have saved my seed phrase.',
  blurScreenText: 'Make sure no one is watching your screen',
};

const YourSeedPhrase: React.FC = () => {
  const [phrase, setPhrase] = useState<string | undefined>(undefined);
  const gno = useGno();
  const navigation = useNavigation<RouterWelcomeStackProp>();

  const onReviewPressHandler = async () => {
    try {
      setPhrase(await gno.generateRecoveryPhrase());
    } catch (error) {
      console.log(error);
    }
  };

  const onNextPressHandler = async () => {
    if (!phrase) return;
    try {
      navigation.navigate(RoutePath.CreatePassword, { phrase });
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <Layout.Container>
      <Layout.Header />
      <Layout.Body>
        <Text.Title>{text.title}</Text.Title>
        <Text.Body>{text.desc}</Text.Body>
        <SeedBox placeholder='Your seed phrase' value={phrase} onChangeText={setPhrase} editable={false} />
        <Spacer />
        {!phrase ? <Button title='Generate' onPress={onReviewPressHandler} variant='primary' /> : null}
        {phrase ? <Button title='Create Account' onPress={onNextPressHandler} variant='primary' /> : null}
      </Layout.Body>
    </Layout.Container>
  );
};

export default YourSeedPhrase;
