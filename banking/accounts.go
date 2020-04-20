package accounts

import "errors"

//Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("your poor")

//NewAccount create Account
func NewAccount(owner string) *Account {
	account := Account{
		owner: owner, balance: 0,
	}
	return &account
}

//Deposit x amount on your account
func (a *Account) Deposit(amount int) {

	a.balance += amount
}

//Balance of your account
func (a Account) Balance() int {
	return a.balance
}

//Withdraw x amount from you account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}

	a.balance -= amount

	return nil
}

//ChangeOwner of the ccount
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Onwer() string {
	return a.owner
}
