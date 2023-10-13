import styled from 'styled-components/native';

export interface Props {
  error?: boolean | string;
}

export const TextInput = styled.TextInput<Props>`
  height: 48px;
  margin-top: 4px;
  margin-bottom: 4px;
  border-width: 1px;
  padding: 10px;
  border-radius: 5px;
  width: 100%;
  border-color: ${(props) => (props.error ? 'red' : 'black')};
`;

export default TextInput;
