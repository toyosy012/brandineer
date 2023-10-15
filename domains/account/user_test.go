package account

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	lessLowerLimitLenUsername = "a1"
	moreLowerLimitLenUsername = "a1@"
	invalidFirstCharUsername  = "@1a"
	firstBlankUsername        = " 1a@"
)

var (
	lessUpperLimitLenUsername = strings.Repeat("a", 100)
	moreUpperLimitLenUsername = strings.Repeat("a", 101)
)

func TestUser(t *testing.T) {
	suite.Run(t, NewUsernameTestSuite())
}

type UsernameTestSuite struct {
	suite.Suite
}

func NewUsernameTestSuite() *UsernameTestSuite {
	return &UsernameTestSuite{}
}

func (s *UsernameTestSuite) TestMoreLowerLen() {
	username := NewUsername(moreLowerLimitLenUsername)
	s.T().Log(fmt.Sprintf("\nテストケース: ユーザー名最短値正常系テスト\nテストデータ: %s\nデータ長: %d", username, len(username)))
	s.Equal(nil, username.Valid())
}

func (s *UsernameTestSuite) TestLessUpperLen() {
	username := NewUsername(lessUpperLimitLenUsername)
	s.T().Log(fmt.Sprintf("\nテストケース: ユーザー名最長値境界値正常系テスト\nテストデータ: %s\nデータ長: %d", username, len(username)))
	s.Equal(nil, username.Valid())
}

func (s *UsernameTestSuite) TestLessLowerLen() {
	username := NewUsername(lessLowerLimitLenUsername)
	s.T().Log(fmt.Sprintf("\nテストケース: ユーザー名最短値境界値異常系テスト\nテストデータ: %s\nデータ長: %d", username, len(username)))
	s.ErrorIs(errMinUsernameLength, username.Valid())
}

func (s *UsernameTestSuite) TestMoreUpperLen() {
	username := NewUsername(moreUpperLimitLenUsername)
	s.T().Log(fmt.Sprintf("\nテストケース: ユーザー名最長値境界値異常系テスト\nテストデータ: %s\nデータ長: %d", username, len(username)))
	s.ErrorIs(errMaxUsernameLength, username.Valid())
}

func (s *UsernameTestSuite) TestInvalidFirstChar() {
	username := NewUsername(invalidFirstCharUsername)
	s.T().Log(fmt.Sprintf("\nテストケース: ユーザー名先頭使用不可能文字使用テスト\nテストデータ: %s", username))
	s.ErrorIs(errUsernameFormat, username.Valid())
}

func (s *UsernameTestSuite) TestFirstBlank() {
	username := NewUsername(firstBlankUsername)
	s.T().Log(fmt.Sprintf("\nテストケース: ユーザー名空文字使用テスト\nテストデータ: %s", username))
	s.ErrorIs(errUsernameFormat, username.Valid())
}
