import Button from '@gno/components/buttons';
import { Spacer } from '@gno/components/row';
import { KeyInfoJson } from '@gnolang/gnonative';

interface SideMenuAccountItemProps {
  account: KeyInfoJson;
  changeAccount: (account: KeyInfoJson) => void;
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
