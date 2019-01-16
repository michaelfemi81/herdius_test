package main
import (
	log "github.com/sirupsen/logrus"
	"math"
	"encoding/hex"


)

//pre initialized values
var GENESIS_WALLETS = [...] string{"01","02"};
const TOTAL_COINBASE_REWARD = 0;
var acoountstore AccountStore


func getInstance() AccountStore {
	return acoountstore;
}
func HexStrToUint8Vec(hex_input string) bool{

	_, err := hex.DecodeString(hex_input);
	if(err != nil){
		log.Warn("Failed HexStrToUint8Vec conversion");
		return false;
	} else{
		return true;
	}


}
func addBalanceToGenesisAccount(){

	const bal uint64 = math.MaxUint64 ;
	const  nonce uint64 = 0 ;
	addrBytes:= make([]byte, 4)
	for _, walletHexStr := range GENESIS_WALLETS {
		if (HexStrToUint8Vec(walletHexStr)) {
		continue;
		}
		addr  := Address{account_addr: addrBytes};
		accstore := getInstance();
		accstore.AddAccount(addr, AccountStore{bal, nonce});
	}
	// Init account for issuing coinbase rewards
	accstore := getInstance();
	accstore.AddAccount(NewAddress(),
		AccountStore{TOTAL_COINBASE_REWARD, nonce});
        accstore.UpdateStateTrieAll();



}
