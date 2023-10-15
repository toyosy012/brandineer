package account

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	MinUsernameLength = 3
	MaxUsernameLength = 100
	UsernameFormat    = "^(\\w|\\d)+[\\w\\d-_.@+*^~¥|]{2,99}"
)

var (
	usernameFormatRegexp = regexp.MustCompile(UsernameFormat)
	errMinUsernameLength = errors.New(fmt.Sprintf("ユーザー名は%d文字以上で入力してください", MinUsernameLength))
	errMaxUsernameLength = errors.New(fmt.Sprintf("ユーザー名は%d文字未満で入力してください", MaxUsernameLength))
	errUsernameFormat    = errors.New("ユーザー名に使用できない文字が含まれています")
)

type Username string

func NewUsername(value string) Username {
	return Username(value)
}
func (u Username) Valid() error {
	if len(u) < MinUsernameLength {
		return errMinUsernameLength
	}

	if MaxUsernameLength < len(u) {
		return errMaxUsernameLength
	}

	if !usernameFormatRegexp.MatchString(string(u)) {
		return errUsernameFormat
	}

	return nil
}

type UserAccount struct {
	username Username
}

func NewUserAccount(name Username) UserAccount {
	return UserAccount{
		username: name,
	}
}

func (u UserAccount) Username() Username { return u.username }
