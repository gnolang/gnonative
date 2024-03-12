import styled from 'styled-components/native';

const SeedBox = styled.TextInput.attrs({
  multiline: true,
  numberOfLines: 4,
})`
  border: 1px solid #000;
  border-radius: 8px;
  padding: 8px;
  height: 80px;
  margin-top: 8px;
  margin-bottom: 8px;
  color: ${(props) => ('editable' in props && !props.editable ? '#545454dd' : 'black')};
  background-color: ${(props) => ('editable' in props && !props.editable ? '#e3e3e3' : 'white')};
`;

export default SeedBox;
