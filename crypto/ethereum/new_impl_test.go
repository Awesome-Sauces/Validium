package ethereum

import (
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

// TestNewPrivateKey tests the NewPrivateKey function for successful key generation
func TestNewPrivateKey(t *testing.T) {
	privKey, err := NewPrivateKey()
	assert.Nil(t, err)
	assert.NotNil(t, privKey)
}

// TestSignAndVerify tests the signing and verification functions
func TestSignAndVerify(t *testing.T) {
	privKey, _ := NewPrivateKey()
	pubKey := privKey.PublicKey()

	data := []byte("test message")
	hash := crypto.Keccak256Hash(data)
	signature, err := privKey.Sign(hash.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, signature)

	isValid := VerifySignature(pubKey, hash.Bytes(), signature)
	assert.True(t, isValid)
}

// TestPublicKeyToAddress tests the PublicKeyToAddress function
func TestPublicKeyToAddress(t *testing.T) {
	privKey, _ := NewPrivateKey()
	pubKey := privKey.PublicKey()
	address := PublicKeyToAddress(pubKey)

	// Basic validation of Ethereum address
	assert.Equal(t, 42, len(address))
	assert.Equal(t, "0x", address[:2])

	log.Println(address)

}

// TestMnemonicToPrivateKey tests mnemonic to private key conversion
func TestMnemonicToPrivateKey(t *testing.T) {
	mnemonic, _ := GenerateMnemonic(12)
	privKey, err := MnemonicToPrivateKey(mnemonic)
	assert.Nil(t, err)
	assert.NotNil(t, privKey)

	log.Println(privKey.ToString())
	log.Println(PublicKeyToAddress(privKey.PublicKey()))
}

// You can add more tests as required...
