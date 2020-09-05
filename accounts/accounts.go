package accounts

import (
	"errors"
	"fmt"
)

var errNoMoney = errors.New("Can't withdraw you are poor")

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount create new Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	fmt.Println("Gonna deposit ", amount)
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil //error의 리턴 타입은 error 와 nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

// String
func (a Account) String() string {

	return fmt.Sprint(a.Owner(), "'s account. \nHas: ", a.Balance())
}
