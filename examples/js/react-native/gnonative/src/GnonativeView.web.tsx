import * as React from 'react';

import { GnonativeViewProps } from './Gnonative.types';

export default function GnonativeView(props: GnonativeViewProps) {
  return (
    <div>
      <span>{props.name}</span>
    </div>
  );
}
