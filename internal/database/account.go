package database

import (
	"errors"
	"fmt"

	"github.com/asaskevich/govalidator"
)

// Account representation
type Account struct {
	id      int64  `json:"id" yaml:"id" valid:"id,required" db:"id"`
	owner   string `json:"owner" yaml:"owner" valid:"owner" db:"owner"`
	balance int64  `json:"balance" yaml:"balance" valid:"balance,optional" db:"balance"`
}

func (a *Account) getId() int64 {
	return a.id
}
func (a *Account) getOwner() string {
	return a.owner
}
func (a *Account) getBalance() int64 {
	return a.balance
}

func (a *Account) GetProperties(props []AccountAttribute) ([]any, error) {
	var r []any
	for _, p := range props {
		switch p {
		case AccountId:
			r = append(r, a.getId())
		case AccountOwner:
			r = append(r, a.getOwner())
		case AccountBalance:
			r = append(r, a.getBalance())
		default:
			return r, errors.New(fmt.Sprintf("property '%s' not fount", p))
		}
	}
	return r, nil
}

// Mimic structure for Builder Pattern
type AccountBuilder struct {
	id      int64
	owner   string
	balance int64
}

// Generate new account
func newAccount() *AccountBuilder {
	return &AccountBuilder{}
}

// Set account id
func (a *AccountBuilder) Id(identifier int64) *AccountBuilder {
	if !govalidator.IsNatural(float64(identifier)) {
		panic(fmt.Sprintf("invalid identifier: %d", identifier))
	}
	a.id = identifier
	return a
}

// Set account owner
func (a *AccountBuilder) Owner(owner string) *AccountBuilder {
	if !govalidator.IsAlpha(owner) {
		panic(fmt.Sprintf("invalid name: %s", owner))
	}
	a.owner = owner
	return a
}

// Set account balance
func (a *AccountBuilder) Balance(balance int64) *AccountBuilder {
	if !govalidator.IsNumeric(string(balance)) {
		panic(fmt.Sprintf("invalid identifier: %d", balance))
	}
	a.balance = balance
	return a
}

// Build account structure
func (u *AccountBuilder) Build() Account {
	return Account{
		id:      u.id,
		owner:   u.owner,
		balance: u.balance,
	}
}

type AccountAttribute int8

const (
	AccountId      AccountAttribute = 0
	AccountOwner   AccountAttribute = 1
	AccountBalance AccountAttribute = 2
)

func (aa AccountAttribute) String() string {
	switch aa {
	case AccountId:
		return "AccountId"
	case AccountOwner:
		return "AccountOwner"
	case AccountBalance:
		return "AccountBalance"
	default:
		return ""
	}
}
