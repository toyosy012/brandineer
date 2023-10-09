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
	value string
}

func NewPassword(value string) (*Password, error) {
	if lessPasswordLen(value) {
		return nil, errMinLen
	}
	if morePasswordLen(value) {
		return nil, errMaxLen
	}

	return &Password{
		value: value,
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
