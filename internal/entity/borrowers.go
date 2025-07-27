package entity

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Borrower struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
}

func (v *Borrower) Validation() error {
	return validation.ValidateStruct(
		v,
		validation.Field(&v.Username, validation.Required),
		validation.Field(&v.Phone, validation.Required),
	)
}
