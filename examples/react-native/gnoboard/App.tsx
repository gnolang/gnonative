// order of imports is important
import 'react-native-polyfill-globals/auto';

import { GnoNativeProvider } from '@gno/provider/gnonative/gnonative-provider';
import CustomRouter from '@gno/router/custom-router';

// Polyfill async.Iterator. For some reason, the Babel presets and plugins are not doing the trick.
// Code from here: https://www.typescriptlang.org/docs/handbook/release-notes/typescript-2-3.html#caveats
(Symbol as any).asyncIterator = Symbol.asyncIterator || Symbol.for('Symbol.asyncIterator');

function App() {
  return (
    <GnoNativeProvider>
      <CustomRouter />
    </GnoNativeProvider>
  );
}

const AppEntryPoint = App;
// const  AppEntryPoint = require('./.storybook').default;

export default AppEntryPoint;
