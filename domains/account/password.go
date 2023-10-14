package account

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptMaxLength = 72
	bcryptMinLength = 8
)

var (
	errMinLen         = errors.New(fmt.Sprintf("パスワードの文字数は%d文字上にしてください", bcryptMinLength))
	errMaxLen         = errors.New(fmt.Sprintf("パスワードの文字数は%d文字上にしてください", bcryptMaxLength))
	errHashedPassword = errors.New("パスワード生成に失敗")
	errCompareHash    = errors.New("パスワード検証に失敗")
)

type Password struct {
	validator Validator
	value     string
	hash      []byte
}

func NewPassword(validator Validator, value string) (*Password, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), 10)
	if err != nil {
		if errors.Is(bcrypt.ErrPasswordTooLong, err) {
			return nil, errMaxLen
		}
		return nil, errHashedPassword
	}
	return &Password{
		validator: validator,
		value:     value,
		hash:      hash,
	}, nil
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

	// 初期化時点で検証されるが最大長のバリデーションルールが存在することを明示するために残す
	if p.morePasswordLen() {
		return errMaxLen
	}

	if err := p.validator.Valid(p); err != nil {
		return err
	}

	return nil
}

func (p Password) CompareHash(input Password) error {
	if err := bcrypt.CompareHashAndPassword(p.hash, []byte(input.Value())); err != nil {
		return errCompareHash
	}
	return nil
}

type Validator interface {
	Valid(Password) error
}
