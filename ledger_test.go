package main

import (
	"testing"
)

func TestCreateAccountWithNonExistentIdYieldsNoError(t *testing.T) {
	ledger := NewLedger()

	err := ledger.CreateAccount(1, 1)

	if err != nil {
		t.Error(`Expected no error`)
	}
}

func TestCreateAccountWithExistingIdYieldsError(t *testing.T) {
	ledger := NewLedger()

	_ = ledger.CreateAccount(1, 1)
	err := ledger.CreateAccount(2, 1)

	if err == nil {
		t.Error(`Expected account creation error`)
	}
}
