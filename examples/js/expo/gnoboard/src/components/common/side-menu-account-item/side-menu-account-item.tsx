import Button from '@gno/components/buttons';
import { Spacer } from '@gno/components/row';
import { GnoAccount } from '@gno/native_modules/types';

interface SideMenuAccountItemProps {
  account: GnoAccount;
  changeAccount: (account: GnoAccount) => void;
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
