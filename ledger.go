package main

import (
	"errors"
	"fmt"
)

// id - int
// timestamp - int
// basic validation -

type Timestamp int
type AccountId int
type MoneyAmount int

type Ledger struct {
	accounts map[AccountId]MoneyAmount
}

func NewLedger() Ledger {
	return Ledger{make(map[AccountId]MoneyAmount)}
}

func (l *Ledger) CreateAccount(timestamp Timestamp, id AccountId) error {
	_, exists := l.accounts[id]
	if exists {
		return errors.New(fmt.Sprintf("Account with id %d already exists", id))
	}
	l.accounts[id] = 0
	return nil
}

func (l *Ledger) Deposit(timestamp Timestamp, account AccountId, amount MoneyAmount) error {
	_, exists := l.accounts[account]
	if !exists {
		return errors.New(fmt.Sprintf("Account with id %d does not exist", account))
	}
	return nil
}

func (l *Ledger) Transfer(timestamp Timestamp, fromAccount AccountId, toAccount AccountId, amount MoneyAmount) error {
	return nil
}

func (l *Ledger) GetCurrentAmount(account AccountId) (error, MoneyAmount) {
	return nil, 0
}
