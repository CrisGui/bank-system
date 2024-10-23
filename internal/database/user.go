package database

import (
	"errors"
	"fmt"

	"github.com/asaskevich/govalidator"
)

// User representation
type User struct {
	id       int64  `json:"id" yaml:"id" valid:"id,required" db:"id"`
	username string `json:"username" yaml:"username" valid:"username,optional" db:"username"`
	fullname string `json:"fullname" yaml:"fullname" valid:"fullname" db:"fullname"`
	phone    string `json:"phoneNumeber" yaml:"phoneNumber" valid:"regex=^\\+{0,1}0{0,1}62[0-9]+$,optional" db:"phoneNumber"`
	email    string `json:"email" yaml:"email" valid:"email,optional" db:"email"`
	account  int64  `json:"accountRef" yaml:"accountRef" valid:"accountRef" db:"accountRef"`
}

func (u *User) getId() int64 {
	return u.id
}
func (u *User) getUsername() string {
	return u.username
}
func (u *User) getFullname() string {
	return u.fullname
}
func (u *User) getPhone() string {
	return u.phone
}
func (u *User) getEmail() string {
	return u.email
}
func (u *User) getAccount() int64 {
	return u.account
}

func (u *User) GetProperties(props []UserAttribute) ([]any, error) {
	var r []any
	for _, p := range props {
		switch p {
		case UserId:
			r = append(r, u.getId())
		case UserUsername:
			r = append(r, u.getUsername())
		case UserFullname:
			r = append(r, u.getFullname())
		case UserPhone:
			r = append(r, u.getPhone())
		case UserEmail:
			r = append(r, u.getEmail())
		case UserAccount:
			r = append(r, u.getAccount())
		default:
			return r, errors.New(fmt.Sprintf("property '%s' not fount", p))
		}
	}
	return r, nil
}

// Mimic structure for Builder Pattern
type UserBuilder struct {
	id       int64
	username string
	fullname string
	phone    string
	email    string
	account  int64
}

// Generate new user
func newUser() *UserBuilder {
	return &UserBuilder{}
}

// Set user id
func (u *UserBuilder) Id(identifier int64) *UserBuilder {
	if !govalidator.IsNatural(float64(identifier)) {
		panic(fmt.Sprintf("invalid identifier: %d", identifier))
	}
	u.id = identifier
	return u
}

// Set user username
func (u *UserBuilder) Username(username string) *UserBuilder {
	if !govalidator.IsAlphanumeric(username) {
		panic(fmt.Sprintf("invalid username: %s", username))
	}
	u.username = username
	return u
}

// Set user fullname
func (u *UserBuilder) Fullname(fullname string) *UserBuilder {
	if !govalidator.IsAlpha(fullname) {
		panic(fmt.Sprintf("invalid name: %s", fullname))
	}
	u.fullname = fullname
	return u
}

// Set user phone
func (u *UserBuilder) Phone(number string) *UserBuilder {
	u.phone = number
	return u
}

// Set user account
func (u *UserBuilder) Email(email string) *UserBuilder {
	if !govalidator.IsEmail(email) {
		panic(fmt.Sprintf("invalid email: %s", email))
	}
	u.email = email
	return u
}

// Set user account
func (u *UserBuilder) Account(identifier int64) *UserBuilder {
	if !govalidator.IsNatural(float64(identifier)) {
		panic(fmt.Sprintf("invalid identifier: %d", identifier))
	}
	u.account = identifier
	return u
}

// Build user structure
func (u *UserBuilder) Build() User {
	return User{
		id:       u.id,
		username: u.username,
		fullname: u.fullname,
		phone:    u.phone,
		email:    u.email,
		account:  u.account,
	}
}

type UserAttribute int8

const (
	UserId       UserAttribute = 0
	UserUsername UserAttribute = 1
	UserFullname UserAttribute = 2
	UserPhone    UserAttribute = 3
	UserEmail    UserAttribute = 4
	UserAccount  UserAttribute = 5
)

func (ua UserAttribute) String() string {
	switch ua {
	case UserId:
		return "UserId"
	case UserUsername:
		return "UserUsername"
	case UserFullname:
		return "UserFullname"
	case UserPhone:
		return "UserPhone"
	case UserEmail:
		return "UserEmail"
	case UserAccount:
		return "UserAccount"
	default:
		return ""
	}
}
