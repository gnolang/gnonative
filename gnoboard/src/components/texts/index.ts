import styled from 'styled-components/native';

//styleName: title/xs/medium;
const Title = styled.Text`
  font-size: 20px;
  font-weight: 500;
  letter-spacing: 0;
`;

/** styleName: body/base/regular; */
const Body = styled.Text`
  font-size: 16px;
  font-weight: 400;
  line-height: 21px;
  letter-spacing: 0;
  text-align: left;
`;

const HeaderTitle = styled.Text`
  color: black;
  font-size: 17px;
  font-weight: 500;
  line-height: 22px;
  letter-spacing: 0;
  text-align: center;
`;

const HeaderSubtitle = styled.Text`
  padding-top: 4px;
  color: #667386;
  font-size: 13px;
  font-weight: 400;
  line-height: 18px;
  letter-spacing: 0.25px;
  text-align: center;
`;

const Text = { Title, Body, HeaderSubtitle, HeaderTitle };

export default Text;
