package main



type AccountStore struct {
  balance uint64;
  nonce uint64;
}

var Accounts  map[*Address]AccountStore;



func (a AccountStore) AddAccount(address Address, as AccountStore) {
   if(Accounts == nil){
	Accounts = make(map[*Address]AccountStore);
   }
   Accounts[&address] = as;
}
func (a AccountStore) UpdateStateTrieAll() bool{
 return true;
}