import { requireNativeViewManager } from 'expo-modules-core';
import * as React from 'react';

import { GnonativeViewProps } from './Gnonative.types';

const NativeView: React.ComponentType<GnonativeViewProps> =
  requireNativeViewManager('Gnonative');

export default function GnonativeView(props: GnonativeViewProps) {
  return <NativeView {...props} />;
}
