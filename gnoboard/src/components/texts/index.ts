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

/** styleName: body/base/medium; */
const BodyMedium = styled.Text`
  font-size: 16px;
  font-weight: 500;
  line-height: 21px;
  letter-spacing: 0;
  text-align: left;
`;

const Caption1 = styled.Text`
  font-size: 11px;
  font-weight: 400;
  line-height: 13px;
  letter-spacing: 0.5px;
  text-align: center;
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

export const HeaderTitleText = styled.Text`
  font-size: 17px;
  font-weight: 500;
  line-height: 22px;
  letter-spacing: 0;
  text-align: center;
`;

export const HeaderSubtitleText = styled.Text`
  padding-top: 4px;
  color: #667386;
  font-size: 13px;
  font-weight: 400;
  line-height: 18px;
  letter-spacing: 0.25px;
  text-align: center;
`;

//styleName: body/subheadline/medium;
const Subheadline = styled.Text`
  color: #8b949e;
  font-size: 15px;
  font-weight: 500;
  line-height: 20px;
  letter-spacing: 0;
  text-align: left;
`;

const Text = { Title, Body, HeaderSubtitle, HeaderTitle, HeaderTitleText, HeaderSubtitleText, Subheadline, BodyMedium, Caption1 };

export default Text;
