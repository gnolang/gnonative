import styled from 'styled-components/native';

//styleName: title/xs/medium;
const Title = styled.Text`
  font-size: 20px;
  font-weight: 500;
  letter-spacing: 0;
  text-align: left;
  padding: 0 0 8px 0;
`;

/** styleName: body/base/regular; */
const Body = styled.Text`
  font-size: 16px;
  font-weight: 400;
  line-height: 21px;
  letter-spacing: 0;
  text-align: left;
`;

const Text = { Title, Body };

export default Text;
