package ethereum

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"

	"github.com/ethereum/go-ethereum/crypto" // Ethereum crypto library
	"github.com/tyler-smith/go-bip39"        // BIP39 library
)

// PrivateKey struct represents a private key
type PrivateKey struct {
	key *ecdsa.PrivateKey
}

// NewPrivateKey generates a new secp256k1 ECDSA private key.
func NewPrivateKey() (*PrivateKey, error) {
	privKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return &PrivateKey{key: privKey}, nil
}

// Sign signs data using the private key.
func (pk *PrivateKey) Sign(data []byte) ([]byte, error) {
	hash := crypto.Keccak256Hash(data)
	signature, err := crypto.Sign(hash.Bytes(), pk.key)
	if err != nil {
		return nil, err
	}
	return signature, nil
}

// VerifySignature verifies a signature given the data and the public key.
func VerifySignature(publicKey *ecdsa.PublicKey, data, signature []byte) bool {
	hash := crypto.Keccak256Hash(data)
	return crypto.VerifySignature(crypto.FromECDSAPub(publicKey), hash.Bytes(), signature[:len(signature)-1])
}

// PublicKey returns the public key associated with the private key.
func (pk *PrivateKey) PublicKey() *ecdsa.PublicKey {
	return &pk.key.PublicKey
}

// PublicKeyToAddress converts a public key to an Ethereum address.
func PublicKeyToAddress(pubKey *ecdsa.PublicKey) string {
	address := crypto.PubkeyToAddress(*pubKey)
	return address.Hex()
}

func PublicKeyToAddressBytes(pubKey *ecdsa.PublicKey) []byte {
	return crypto.PubkeyToAddress(*pubKey).Bytes()
}

// GenerateMnemonic generates a new mnemonic phrase.
// The wordCount should be one of the following: 12, 15, 18, 21, 24.
func GenerateMnemonic(wordCount int) (string, error) {

	bitSize := int(float64(wordCount) * 10.6666666667)
	entropy, err := bip39.NewEntropy(bitSize)
	if err != nil {
		return "", err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}
	return mnemonic, nil
}

// MnemonicToPrivateKey converts a mnemonic to a private key.
func MnemonicToPrivateKey(mnemonic string) (*PrivateKey, error) {
	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, errors.New("invalid mnemonic")
	}
	seed := bip39.NewSeed(mnemonic, "") // Generates a 512-bit seed

	privKey, err := crypto.ToECDSA(seed[:32]) // Use only the first 32 bytes of the seed
	if err != nil {
		return nil, err
	}
	return &PrivateKey{key: privKey}, nil
}

// ToString returns the hexadecimal representation of the private key.
func (pk *PrivateKey) ToString() string {
	return hex.EncodeToString(crypto.FromECDSA(pk.key))
}

// VerifySignatureFromAddress verifies a signature given an Ethereum address, data, and a signature.
func VerifySignatureFromAddress(address string, data, signature []byte) (bool, error) {
	hash := crypto.Keccak256Hash(data)

	publicKey, err := crypto.SigToPub(hash.Bytes(), signature) // Recover the public key from the signature
	if err != nil {
		return false, err
	}

	recoveredAddr := crypto.PubkeyToAddress(*publicKey) // Convert the recovered public key to an Ethereum address
	return recoveredAddr.Hex() == address, nil
}
