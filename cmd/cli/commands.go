package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/Awesome-Sauces/Validium/crypto/ethereum"
	"github.com/spf13/cobra"
)

var message string

var boot = &cobra.Command{
	Use:   "boot",
	Short: "Boots the node with the defined flags. (Defaults to \"<Refresh: True><Port: 3003>\")",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Flags: %s\n", message)
	},
}

var refresh = &cobra.Command{
	Use:   "refresh",
	Short: "Updates node to latest version",
	Long:  "Updates the node to the latest approved version. (Updates to Helix after v1.0.0 are made purely through the evm, unless they are updates relating to efficiency or speed.)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Node Booted")
	},
}

var badger = &cobra.Command{
	Use:   "hbbft",
	Short: "Test the Honey Badger Consensus",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World")
	},
}

var ethkeys = &cobra.Command{
	Use:   "ethkey",
	Short: "Test Ethereum Keys",
	Run: func(cmd *cobra.Command, args []string) {
		privatekey, err := ethereum.NewPrivateKey()

		privatekey2, err := ethereum.NewPrivateKey()

		fmt.Println(ethereum.PublicKeyToAddress(privatekey.PublicKey()))
		fmt.Println(ethereum.PublicKeyToAddress(privatekey2.PublicKey()))

		if err != nil {
			return
		}

		// 40-bit value
		value := uint64(0xFFFFFFFFFF)

		// Create a byte slice of size 8 (since uint64 is 8 bytes)
		bytes := make([]byte, 8)

		// Put the value into the byte slice using BigEndian
		binary.BigEndian.PutUint64(bytes, value)

		// Slice the last 5 bytes since those represent the 40-bit value
		bytes = bytes[3:]

		// Assuming ethereum.PublicKeyToAddress(privatekey2.PublicKey()) returns a byte slice
		address := []byte(ethereum.PublicKeyToAddress(privatekey2.PublicKey()))
		value = uint64(0xFFFFFFFFFFFFFFFF)
		tx_sequence := make([]byte, 8)

		binary.BigEndian.PutUint64(tx_sequence, value)

		// Append the address bytes to the bytes slice
		bytes = append(bytes, tx_sequence...)
		bytes = append(bytes, address...)

		signature, err := privatekey.Sign(bytes)

		fmt.Println(hex.EncodeToString(signature))

		fmt.Println(signature)
	},
}

func AddCommands() {
	rootCmd.AddCommand(refresh)

	rootCmd.AddCommand(boot)

	rootCmd.AddCommand(badger)

	rootCmd.AddCommand(ethkeys)

	boot.Flags().StringVarP(&message, "flags", "f", "<FlagName: Parameter>...", "Custom message to print")
}
