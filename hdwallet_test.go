package hdwallet_test

import (
	"encoding/hex"
	"testing"

	"github.com/e4coder/hdwallet"
)

const MASTER_BIP32_ROOT_KEY = "xprv9s21ZrQH143K2LBWUUQRFXhucrQqBpKdRRxNVq2zBqsx8HVqFk2uYo8kmbaLLHRdqtQpUm98uKfu3vca1LqdGhUtyoFnCNkfmXRyPXLjbKb"

func TestDerivation(t *testing.T) {
	t.Log("TestDerivation")
	wallet, err := hdwallet.NewFromRootKey(MASTER_BIP32_ROOT_KEY)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	childKey, err := wallet.DeriveFromPath("m/83696968'/0'/0'")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(hex.EncodeToString(childKey.Key))
}

func TestBip85(t *testing.T) {
	t.Log("TestBip85")
	wallet, err := hdwallet.NewFromRootKey(MASTER_BIP32_ROOT_KEY)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	childKey, err := wallet.DeriveFromPath("m/83696968'/0'/0'")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(hex.EncodeToString(childKey.Key))
}
