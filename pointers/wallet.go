package main

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("No sufficient money to withdraw")

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct{ balance Bitcoin }

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance-amount < 0 {
		return ErrInsufficientFunds
	}

	(*w).balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return (*w).balance
}

func main() {
}
