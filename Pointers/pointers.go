package pointers

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}
type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposite(val Bitcoin) {
	w.balance += val
}

var errorrrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(val Bitcoin) error {
	if w.balance < val {
		return errorrrInsufficientFunds
	}
	w.balance -= val
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
