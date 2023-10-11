package hdwallet

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

type HDWallet struct {
	RootKey  *bip32.Key
	Seed     []byte
	FromRoot bool
}

func (hd *HDWallet) Derive(path []uint32) (*bip32.Key, error) {
	var currKey *bip32.Key = hd.RootKey
	for _, childIdX := range path {
		k, err := hd._deriveOne(currKey, childIdX)
		if err != nil {
			return nil, err
		}
		currKey = k
	}

	return currKey, nil
}

func (hd *HDWallet) _deriveOne(key *bip32.Key, childIdX uint32) (*bip32.Key, error) {
	childKey, err := key.NewChildKey(childIdX)
	if err != nil {
		return nil, err
	}

	return childKey, nil
}

func (hd *HDWallet) DeriveFromPath(path string) (*bip32.Key, error) {
	derivedPath, err := ParseDerivationPath(path)
	if err != nil {
		return &bip32.Key{}, err
	}

	return hd.Derive(derivedPath)
}

func ParseDerivationPath(path string) (accounts.DerivationPath, error) {
	parsed, err := accounts.ParseDerivationPath(path)
	return parsed, err
}

// NewEntropy will create random entropy bytes so long as the requested size bitSize is an appropriate size.
//
// bitSize has to be a multiple 32 and be within the inclusive range of {128, 256}
func NewEntropy(bitSize int) ([]byte, error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return nil, err
	}

	return entropy, nil
}

// NewMnemonic will return a string consisting of the mnemonic words for the given entropy. If the provide entropy is invalid, an error will be returned.
func NewMnemonic() (string, error) {
	entropy, err := NewEntropy(128)
	if err != nil {
		return "", err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}
	return mnemonic, nil
}

// NewMnemonicFromEntropy generates a mnemonic representation of the provided entropy.
// The entropy needs to be a byte array. If the entropy is invalid, it returns an error.
func NewMnemonicFromEntropy(entropy []byte) (string, error) {
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}
	return mnemonic, nil
}

// New creates a new HDWallet given a passphrase. It will first generate a new mnemonic and then, using that mnemonic
// and the provided passphrase, it creates a seed. This seed is then used to generate a master key, which is
// the root key of the wallet. If there are any errors in the process, it will return the error.
func New(passphrase string) (*HDWallet, error) {
	mnemonic, err := NewMnemonic()
	if err != nil {
		return nil, err
	}

	seed := bip39.NewSeed(mnemonic, passphrase)

	rootKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, err
	}

	return &HDWallet{
		RootKey: rootKey,
		Seed:    seed,
	}, nil

}

// NewFromMnemonic creates a new HDWallet given a mnemonic and a passphrase. It first generates a seed using the mnemonic
// and the provided passphrase. This seed is then used to generate a master key, which is the root key of the wallet.
// If there are any errors in the process, it will return the error.
func NewFromMnemonic(mnemonic string, passphrase string) (*HDWallet, error) {
	seed := bip39.NewSeed(mnemonic, passphrase)

	rootKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, err
	}

	return &HDWallet{
		RootKey: rootKey,
		Seed:    seed,
	}, nil
}

// NewFromSeed creates a new HDWallet given a seed. It uses this seed to generate a master key,
// which is the root key of the wallet. If there are any errors in the process of generating the master key,
// it will return the error.
func NewFromSeed(seed []byte) (*HDWallet, error) {
	rootKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, err
	}

	return &HDWallet{
		RootKey: rootKey,
		Seed:    seed,
	}, nil
}

func NewFromRootKey(root string) (*HDWallet, error) {
	rootKey, err := bip32.B58Deserialize(root)
	if err != nil {
		return nil, err
	}

	return &HDWallet{
		RootKey:  rootKey,
		FromRoot: true,
	}, nil
}

// the purpose of the wallet is to Generate a new masterkey
// 1. From mnemonic
// 2. From seed
// 3. From base58 private key
// and to be able to derive child keys from a derivation path
// also able to derive multiple wallets from the masterkey
func Main() {

}
