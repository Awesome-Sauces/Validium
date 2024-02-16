package ozone

import (
	"encoding/json"
	"log"
	"math/big"
)

type t_Block struct {
	transactions []string
}

type t_Ledger struct {
	accounts map[string]big.Int
	blocks   []t_Block
}

func (ledger *t_Ledger) NewAccount(address string, balance big.Int) {
	ledger.accounts[address] = balance
}

func (ledger *t_Ledger) ToBytes() []byte {
	serializableAccounts := make(map[string]string)
	for address, balance := range ledger.accounts {
		serializableAccounts[address] = balance.String() // Convert big.Int to string
	}

	data, err := json.Marshal(serializableAccounts)
	if err != nil {
		log.Fatalf("Error marshaling ledger accounts: %v", err)
	}
	return data
}

func (ledger *t_Ledger) FromBytes(data []byte) {
	serializableAccounts := make(map[string]string)
	err := json.Unmarshal(data, &serializableAccounts)
	if err != nil {
		log.Fatalf("Error unmarshaling ledger accounts: %v", err)
	}

	ledger.accounts = make(map[string]big.Int)
	for address, balanceStr := range serializableAccounts {
		var balance big.Int
		balance.SetString(balanceStr, 10) // Assuming balance is in base 10
		ledger.accounts[address] = balance
	}
}
