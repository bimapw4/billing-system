package entity

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Claim struct {
	UserID   string
	Username string
	IsAdmin  bool
	Exp      int
}

type Authorization struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (v *Authorization) Validation() error {
	return validation.ValidateStruct(
		v,
		validation.Field(&v.Username, validation.Required),
		validation.Field(&v.Password, validation.Required),
	)
}
