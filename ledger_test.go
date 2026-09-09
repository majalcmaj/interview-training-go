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

func TestDepositToNonExistentAccountYieldsError(t *testing.T) {
	ledger := NewLedger()

	err := ledger.Deposit(2, 1, 100)

	if err == nil {
		t.Error(`Account does not exist - expected error`)
	}
}

func TestGetCurrentBalanceReturnsErrorWhenNoAccount(t *testing.T) {
	ledger := NewLedger()

	err, _ := ledger.GetCurrentBalance(1)

	if err == nil {
		t.Error(`Expected error when getting amount for non existent account`)
	}
}

func TestGetCurrentBalanceReturns0ForFreshAccount(t *testing.T) {
	ledger := NewLedger()
	_ = ledger.CreateAccount(1, 1)

	err, balance := ledger.GetCurrentBalance(1)

	if err != nil {
		t.Error(`Unexpected error`)
	}
	if balance != 0 {
		t.Errorf(`Expected balance 0, got %d`, balance)
	}
}

func TestDepositToAccountChangesBalance(t *testing.T) {
	ledger := NewLedger()
	ledger.CreateAccount(1, 1)

	_ = ledger.Deposit(2, 1, 100)
	_ = ledger.Deposit(2, 1, 50)

	_, amount := ledger.GetCurrentBalance(1)
	if amount != 150 {
		t.Errorf(`Expected value of 150 on the account, got %d`, amount)
	}
}
