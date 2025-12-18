package validate

import "github.com/go-playground/validator/v10"

type Validator struct {
	v *validator.Validate
}

func New() *Validator {
	return &Validator{v: validator.New()}
}

func (v *Validator) Struct(s any) error {
	return v.v.Struct(s)
}
