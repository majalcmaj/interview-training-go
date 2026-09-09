package main

import (
	"container/heap"
	"errors"
	"fmt"
)

// id - int
// timestamp - int
// basic validation -

type Timestamp int
type AccountId int
type MoneyAmount uint

type AccountData struct {
	balance   MoneyAmount
	spendings MoneyAmount
}

func (ad *AccountData) Deposit(amount MoneyAmount) {
	ad.balance += amount
}

func (ad *AccountData) Withdraw(amount MoneyAmount) {
	ad.balance -= amount
	ad.spendings += amount
}

type Ledger struct {
	accounts map[AccountId]*AccountData
}

func NewLedger() Ledger {
	return Ledger{make(map[AccountId]*AccountData)}
}

func (l *Ledger) CreateAccount(timestamp Timestamp, id AccountId) error {
	_, exists := l.accounts[id]
	if exists {
		return errors.New(fmt.Sprintf("Account with id %d already exists", id))
	}
	l.accounts[id] = &AccountData{0, 0}
	return nil
}

func (l *Ledger) Deposit(timestamp Timestamp, account AccountId, amount MoneyAmount) error {
	_, exists := l.accounts[account]
	if !exists {
		return errors.New(fmt.Sprintf("Account with id %d does not exist", account))
	}
	l.accounts[account].Deposit(amount)
	return nil
}

func (l *Ledger) Transfer(timestamp Timestamp, fromAccountId AccountId, toAccountId AccountId, amount MoneyAmount) error {
	fromAccount, fromExists := l.accounts[fromAccountId]
	_, toExists := l.accounts[toAccountId]
	if !fromExists || !toExists {
		return errors.New(fmt.Sprintf("Accounts with ids %d and %d need to exist", fromAccount, toAccountId))
	}

	if fromAccountId == toAccountId {
		return errors.New(fmt.Sprintf("Cannot transfer to same account %d", fromAccount))
	}

	if fromAccount.balance < amount {
		return errors.New(fmt.Sprintf("Insufficient funds on account %d", fromAccount))
	}

	l.accounts[fromAccountId].Withdraw(amount)
	l.accounts[toAccountId].Deposit(amount)

	return nil
}

func (l *Ledger) GetCurrentBalance(accountId AccountId) (error, MoneyAmount) {
	account, exists := l.accounts[accountId]

	if !exists {
		return errors.New(fmt.Sprintf("Account with id %d does not exist", accountId)), 1
	}

	return nil, account.balance
}

type idAndSpending struct {
	id       AccountId
	spending MoneyAmount
}

type spendingHeap []idAndSpending

func (s spendingHeap) Len() int           { return len(s) }
func (s spendingHeap) Less(i, j int) bool { return s[i].spending < s[j].spending }
func (s spendingHeap) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s *spendingHeap) Push(el any) {
	*s = append(*s, el.(idAndSpending))
}

func (s *spendingHeap) Pop() any {
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}

func (l *Ledger) GetTopSpenders(number int) []AccountId {
	if len(l.accounts) <= number {
		return l.getAllAccountIds()
	}

	h := spendingHeap(make([]idAndSpending, 0, number))

	i := 0
	for id, acc := range l.accounts {
		if i < number {
			heap.Push(&h, idAndSpending{id, acc.spendings})
		} else if h[0].spending < acc.spendings {
			h[0] = idAndSpending{id, acc.spendings}
			heap.Fix(&h, 0)
		}
		i++
	}

	result := make([]AccountId, number)
	for idx, idSp := range h {
		result[idx] = idSp.id
	}
	return result
}

func (l *Ledger) getAllAccountIds() []AccountId {
	result := make([]AccountId, len(l.accounts))
	i := 0
	for accountId, _ := range l.accounts {
		result[i] = accountId
		i++
	}
	return result
}
