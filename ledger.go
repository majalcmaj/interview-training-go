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
	balance, exists := l.accounts[account]
	if !exists {
		return errors.New(fmt.Sprintf("Account with id %d does not exist", account))
	}
	l.accounts[account] = balance + amount
	return nil
}

func (l *Ledger) Transfer(timestamp Timestamp, fromAccount AccountId, toAccount AccountId, amount MoneyAmount) error {
	fromBalance, fromExists := l.accounts[fromAccount]
	_, toExists := l.accounts[toAccount]
	if !fromExists || !toExists {
		return errors.New(fmt.Sprintf("Accounts with ids %d and %d need to exist", fromAccount, toAccount))
	}

	if fromAccount == toAccount {
		return errors.New(fmt.Sprintf("Cannot transfer to same account %d", fromAccount))
	}

	if fromBalance-amount < 0 {
		return errors.New(fmt.Sprintf("Insufficient funds on account %d", fromAccount))

	}

	l.accounts[fromAccount] -= amount
	l.accounts[toAccount] += amount

	return nil
}

func (l *Ledger) GetCurrentBalance(account AccountId) (error, MoneyAmount) {
	balance, exists := l.accounts[account]

	if !exists {
		return errors.New(fmt.Sprintf("Account with id %d does not exist", account)), -1
	}

	return nil, balance
}
