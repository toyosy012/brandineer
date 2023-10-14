package account

import (
	"errors"
	"fmt"
)

const (
	bcryptMaxLength = 72
	bcryptMinLength = 8
)

var (
	errMinLen = errors.New(fmt.Sprintf("パスワードの文字数は%d文字上にしてください", bcryptMinLength))
	errMaxLen = errors.New(fmt.Sprintf("パスワードの文字数は%d文字上にしてください", bcryptMaxLength))
)

type Password struct {
	validator Validator
	value     string
}

func NewPassword(validator Validator, value string) Password {
	return Password{
		validator: validator,
		value:     value,
	}
}

func (p Password) Value() string {
	return p.value
}

func (p Password) lessPasswordLen() bool {
	return len(p.Value()) < bcryptMinLength
}
func (p Password) morePasswordLen() bool {
	return bcryptMaxLength < len(p.Value())
}

func (p Password) Valid() error {
	if p.lessPasswordLen() {
		return errMinLen
	}

	if p.morePasswordLen() {
		return errMaxLen
	}

	if err := p.validator.Valid(p); err != nil {
		return err
	}

	return nil
}

type Validator interface {
	Valid(Password) error
}
