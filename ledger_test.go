package main

import (
	"slices"
	"sort"
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

func TestTopNSpendersWhenNoAccountsYieldsEmptySlice(t *testing.T) {
	ledger := NewLedger()

	topSpenders := ledger.GetTopSpenders(3)

	if len(topSpenders) != 0 {
		t.Fatalf("Should get empty slice when no accounts added")
	}
}

func TestTopNSpendersWhenSomeAccountsPresent(t *testing.T) {
	ledger := NewLedger()
	ledger.CreateAccount(1, 1)
	ledger.CreateAccount(2, 2)
	ledger.Deposit(3, 1, 100)
	ledger.Transfer(4, 1, 2, 10)

	topSpenders := ledger.GetTopSpenders(3)

	assertTopSpenderIds(t, []AccountId{1, 2}, topSpenders)
}

func TestTopNSpendersWhenMoreThanNAccountsPresent(t *testing.T) {
	ledger := NewLedger()
	ledger.CreateAccount(1, 1)
	ledger.CreateAccount(2, 2)
	ledger.CreateAccount(3, 3)
	ledger.CreateAccount(4, 4)
	ledger.CreateAccount(5, 5)
	ledger.CreateAccount(6, 6)
	ledger.Deposit(7, 1, 100)
	ledger.Deposit(8, 2, 100)
	ledger.Deposit(9, 3, 100)
	ledger.Deposit(10, 4, 100)
	ledger.Deposit(11, 5, 100)
	ledger.Deposit(12, 6, 100)

	ledger.Transfer(13, 1, 2, 10)
	ledger.Transfer(14, 6, 1, 100) // 6 has spendings of 100 -> 1st
	ledger.Transfer(15, 1, 2, 15)  // 1 has spendings of 25 -> 3rd
	ledger.Transfer(16, 3, 6, 20)  // 3 has spendings of 20
	ledger.Transfer(17, 2, 3, 30)  // 2 has spendings of 30 -> 2nd
	ledger.Transfer(18, 4, 5, 5)   // 4 has spendings of 5

	topSpenders := ledger.GetTopSpenders(3)

	assertTopSpenderIds(t, []AccountId{1, 2, 6}, topSpenders)
}

func assertTopSpenderIds(t *testing.T, expected, actual []AccountId) {
	sort.Slice(actual, func(i, j int) bool { return actual[i] < actual[j] })
	if !slices.Equal(expected, actual) {
		t.Fatalf("Expected top spenders with ids %v but got %v", expected, actual)
	}
}
