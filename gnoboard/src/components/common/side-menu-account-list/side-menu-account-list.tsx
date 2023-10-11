import { GnoAccount } from "@gno/native_modules/types";
import SideMenuAccountItem from "../side-menu-account-item/side-menu-account-item";

interface SideMenuAccountListProps {
  accounts: GnoAccount[];
  changeAccount: (account: GnoAccount) => void;
}

const SideMenuAccountList: React.FC<SideMenuAccountListProps> = ({ accounts, changeAccount }) => {
  return accounts.map((account, index) => <SideMenuAccountItem key={index} account={account} changeAccount={changeAccount} />);
};

export default SideMenuAccountList;
