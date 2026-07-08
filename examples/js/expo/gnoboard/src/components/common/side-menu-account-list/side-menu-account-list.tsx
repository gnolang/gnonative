import { KeyInfoJson } from '@gnolang/gnonative';
import SideMenuAccountItem from '../side-menu-account-item/side-menu-account-item';

interface SideMenuAccountListProps {
  accounts: KeyInfoJson[];
  changeAccount: (account: KeyInfoJson) => void;
}

const SideMenuAccountList: React.FC<SideMenuAccountListProps> = ({ accounts, changeAccount }) => {
  return accounts.map((account, index) => <SideMenuAccountItem key={index} account={account} changeAccount={changeAccount} />);
};

export default SideMenuAccountList;
