import Button from "@gno/components/buttons";
import { GnoAccount } from "@gno/native_modules/types";

interface SideMenuAccountItemProps {
  account: GnoAccount;
  changeAccount: (account: GnoAccount) => void;
}

const SideMenuAccountItem = (props: SideMenuAccountItemProps) => {
  const { account, changeAccount } = props;
  return <Button title={account.name} onPress={() => changeAccount(account)} />;
};

export default SideMenuAccountItem;
