import React, { useEffect, useState } from 'react';
import Layout from '@gno/components/pages';
import { useGnoNativeContext } from '@gnolang/gnonative';
import { RouterWelcomeStack, RouterWelcomeStackProp } from '@gno/router/custom-router';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { useNavigation } from '@react-navigation/native';
import Markdown from 'react-native-marked';
import Loading from '../loading';
import Text from '@gno/components/texts';

export type Props = NativeStackScreenProps<RouterWelcomeStack, 'Board'>;

const Board: React.FC<Props> = ({ route }) => {
  const { board, thread } = route.params;

  const { gnonative } = useGnoNativeContext();
  const navigation = useNavigation<RouterWelcomeStackProp>();
  const [loading, setLoading] = useState<string | undefined>(undefined);
  const [renderedBoard, setRenderedBoard] = useState<string | undefined>(undefined);

  useEffect(() => {
    const unsubscribe = navigation.addListener('focus', async () => {
      try {
        setLoading('Calling Gno Render Function...');

        const response = await gnonative.render('gno.land/r/demo/boards', 'testboard/1');
        setRenderedBoard(response);

        setLoading(undefined);
      } catch (error: unknown | Error) {
        setLoading(error?.toString());
        console.log(error);
      }
    });
    return unsubscribe;
  }, [navigation]);

  if (loading) return <Loading message={loading} />;

  return (
    <>
      <Layout.Container>
        <Layout.Header />
        <Layout.Body>
          <Text.Title>{`${board}`}</Text.Title>
          <Text.Body>{`${thread}`}</Text.Body>

          {renderedBoard ? (
            <Markdown
              value={renderedBoard}
              flatListProps={{
                initialNumToRender: 8,
              }}
            />
          ) : null}
        </Layout.Body>
      </Layout.Container>
    </>
  );
};

export default Board;
