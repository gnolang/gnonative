import Button from '@gno/components/buttons';
import { Spacer } from '@gno/components/row';
import { KeyInfo } from '@gnolang/gnonative';

interface SideMenuAccountItemProps {
  account: KeyInfo;
  changeAccount: (account: KeyInfo) => void;
}

const SideMenuAccountItem = (props: SideMenuAccountItemProps) => {
  const { account, changeAccount } = props;
  return (
    <>
      <Spacer />
      <Button title={account.name} onPress={() => changeAccount(account)} variant='primary' />
    </>
  );
};

export default SideMenuAccountItem;
