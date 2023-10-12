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
	errMinLen     = errors.New(fmt.Sprintf("パスワードの文字数は%d文字上にしてください", bcryptMinLength))
	errMaxLen     = errors.New(fmt.Sprintf("パスワードの文字数は%d文字上にしてください", bcryptMaxLength))
	errValidation = errors.New("不正なパスワードです")
)

type Password struct {
	validator Validator
	value     string
}

func NewPassword(validator Validator, value string) (*Password, error) {
	if lessPasswordLen(value) {
		return nil, errMinLen
	}
	if morePasswordLen(value) {
		return nil, errMaxLen
	}

	return &Password{
		validator: validator,
		value:     value,
	}, nil
}

func (p Password) Value() string {
	return p.value
}

func lessPasswordLen(password string) bool {
	return len(password) < bcryptMinLength
}
func morePasswordLen(password string) bool {
	return bcryptMaxLength < len(password)
}

func (p Password) Valid() error {
	if err := p.validator.Valid(p); err != nil {
		return errValidation
	}

	return nil
}

type Validator interface {
	Valid(Password) error
}
