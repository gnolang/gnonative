// order of imports is important
import 'react-native-polyfill-globals/auto';

import { GnoNativeProvider } from '@gnolang/gnonative';
import { GnoboardProvider } from '@gno/provider/gnoboard-provider';
import CustomRouter from '@gno/router/custom-router';

// Polyfill async.Iterator. For some reason, the Babel presets and plugins are not doing the trick.
// Code from here: https://www.typescriptlang.org/docs/handbook/release-notes/typescript-2-3.html#caveats
(Symbol as any).asyncIterator = Symbol.asyncIterator || Symbol.for('Symbol.asyncIterator');

function App() {
  const defaultConfig = {
    remote: 'gno.land:26657',
    chain_id: 'portal-loop',
  };

  return (
    <GnoNativeProvider config={defaultConfig}>
      <GnoboardProvider>
        <CustomRouter/>
      </GnoboardProvider>
    </GnoNativeProvider>
  );
}

const AppEntryPoint = App;
// const  AppEntryPoint = require('./.storybook').default;

export default AppEntryPoint;
