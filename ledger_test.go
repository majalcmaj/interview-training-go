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

func TestTransferFromOrToNotExistingAccountYieldsError(t *testing.T) {
	ledger := NewLedger()
	ledger.CreateAccount(1, 1)
	_ = ledger.Deposit(2, 1, 100)

	for _, err := range []error{
		ledger.Transfer(3, 2, 3, 10),
		ledger.Transfer(4, 1, 2, 10),
		ledger.Transfer(5, 2, 1, 10),
	} {
		if err == nil {
			t.Error(`tansfers from/to non-existing accounts should fail`)
		}
	}
}

func TestTransferToSameAccountYieldsError(t *testing.T) {
	ledger := NewLedger()
	ledger.CreateAccount(1, 1)
	_ = ledger.Deposit(2, 1, 100)

	err := ledger.Transfer(3, 1, 1, 10)
	if err == nil {
		t.Error(`Cannot transfer to same account`)
	}
}

func TestTransferWithInsufficientFundsYieldsError(t *testing.T) {
	ledger := NewLedger()
	ledger.CreateAccount(1, 1)
	ledger.CreateAccount(2, 2)
	_ = ledger.Deposit(3, 1, 100)

	err := ledger.Transfer(4, 1, 2, 150)

	if err == nil {
		t.Error(`Cannot transfer when insufficient balance`)
	}
}

func TestSuccessfulTransfer(t *testing.T) {
	ledger := NewLedger()
	ledger.CreateAccount(1, 1)
	ledger.CreateAccount(2, 2)
	_ = ledger.Deposit(3, 1, 100)

	err := ledger.Transfer(4, 1, 2, 80)

	if err != nil {
		t.Errorf(`Transfer failure: %v`, err)
	}

	_, fromBalance := ledger.GetCurrentBalance(1)
	if fromBalance != 20 {
		t.Errorf(`Expected account 1 balance to be 20 but was: %d`, fromBalance)
	}
	_, toBalance := ledger.GetCurrentBalance(2)
	if toBalance != 80 {
		t.Errorf(`Expected account 2 balance to be 80 but was: %d`, toBalance)
	}
}
