package ozone

import (
	"fmt"
	"log"
	"testing"
)

func TestBytesToHexArray(t *testing.T) {
	data := []byte{123, 34, 48, 120, 65, 49, 115, 97, 117, 99, 101, 34, 58, 34, 49, 54, 49, 34, 44, 34, 48, 120, 74, 111, 104, 110, 110, 121, 34, 58, 34, 50, 53, 53, 34, 125}

	hexArray := make([]string, len(data))
	for i, b := range data {
		hexArray[i] = fmt.Sprintf("0x%s", fmt.Sprintf("%02X", b))
	}

	log.Println(hexArray)
}

func TestNewOzone(t *testing.T) {
	db, err := NewOzone("ledger-test")

	if err != nil {
		log.Fatal(err)
		return
	}

	if db == nil {
		log.Println("FATAL: DB FAILED TO OPEN")
		return
	}

	defer db.Close()

	// Iterate over all items in the database
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()

		// Deserialize value into t_Ledger
		ledger := &t_Ledger{}
		ledger.FromBytes(value)

		// Print ledger contents
		log.Printf("Ledger for key %s: %+v\n", key, ledger.accounts)
	}
	iter.Release()
	if err := iter.Error(); err != nil {
		t.Fatalf("Iterator error: %v", err)
	}

}
