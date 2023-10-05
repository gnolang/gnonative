import * as React from 'react';
import Svg, { G, Path, SvgProps } from 'react-native-svg';

export default function Close(props: SvgProps) {
  return (
    <Svg width={props.width || 24} height={props.height || 24} viewBox='0 0 24 24' fill='none'>
      <G id='Basic/Close'>
        <Path
          id='Path'
          fillRule='evenodd'
          clipRule='evenodd'
          d='M17.6833 4.57021C18.1627 4.09078 18.94 4.09078 19.4195 4.57021C19.8989 5.04964 19.8989 5.82696 19.4195 6.30639L13.7482 11.9776L19.4246 17.654C19.9136 18.143 19.9136 18.9357 19.4246 19.4246C18.9357 19.9136 18.143 19.9136 17.6541 19.4246L11.9776 13.7482L6.30639 19.4195C5.82696 19.8989 5.04965 19.8989 4.57021 19.4195C4.09078 18.94 4.09078 18.1627 4.57021 17.6833L10.2415 12.012L4.57538 6.34595C4.08645 5.85702 4.08645 5.06431 4.57538 4.57538C5.06431 4.08645 5.85702 4.08645 6.34595 4.57538L12.012 10.2415L17.6833 4.57021Z'
          fill={props.color ? props.color : 'black'}
        />
      </G>
    </Svg>
  );
}
