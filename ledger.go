package main

import (
	"fmt"
)

// id - int
// timestamp - int
// basic validation -

type Timestamp int
type AccountId int
type MoneyAmount int

type Ledger interface {
	CreateAccount(timestamp Timestamp, id AccountId) error
	Deposit(timestamp Timestamp, account AccountId, amount MoneyAmount) error
	Transfer(timestamp Timestamp, fromAccount AccountId, toAccount AccountId, amount MoneyAmount) error
	GetCurrentAmount(account AccountId) (error, MoneyAmount)
}

func doSomething() int {
	fmt.Println("Hello world!")
	return 1
}
