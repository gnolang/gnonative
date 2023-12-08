import * as React from 'react';
import Svg, { Path, SvgProps } from 'react-native-svg';

const CheckMark = (props: SvgProps) => (
  <Svg width={24} height={24} fill='none' {...props}>
    <Path
      fillRule='evenodd'
      clipRule='evenodd'
      d='M18.181 5.389a1.25 1.25 0 0 1 1.99 1.508l-.083.109-9.583 11.302a1.25 1.25 0 0 1-1.808.104l-.099-.103-4.712-5.555a1.25 1.25 0 0 1 1.812-1.716l.094.099 3.759 4.43 8.63-10.178Z'
      fill={props.color || '#fff'}
    />
  </Svg>
);

export default CheckMark;
