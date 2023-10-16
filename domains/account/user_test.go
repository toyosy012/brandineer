package account

import (
	"fmt"
	"net/mail"
	"strings"
	"testing"

	"github.com/google/uuid"
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

func TestUsername(t *testing.T) {
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

const (
	validEmail = "user1@example.com"
	validUUID  = "f311a818-5d48-4603-8543-368976dcca6b"
)

func TestUserAccount(t *testing.T) {
	suite.Run(t, NewUserAccountTestSuite())
}

type TestUserAccountSuite struct {
	suite.Suite

	validEmailAddress mail.Address
	passwordValidator Validator
	validPassword     Password
	validUUID         uuid.UUID
}

func (s *TestUserAccountSuite) SetupSuite() {
	validEmailAddress, err := mail.ParseAddress(validEmail)
	if err != nil {
		s.Fail(err.Error())
		return
	}
	s.validEmailAddress = *validEmailAddress

	s.passwordValidator = NewNSPVValidator()
	validPassword, err := NewPassword(s.passwordValidator, validLowerLengthPassword)
	if err != nil {
		s.Fail(err.Error())
		return
	}
	s.validPassword = *validPassword

	id, err := uuid.Parse(validUUID)
	if err != nil {
		s.Fail(err.Error())
		return
	}
	s.validUUID = id
}

func NewUserAccountTestSuite() *TestUserAccountSuite { return &TestUserAccountSuite{} }

func (s *TestUserAccountSuite) TestValidCreateUserAccount() {
	s.T().Log(
		fmt.Sprintf(
			"\nテストケース: ユーザーアカウント作成処理正常系テスト\nテストデータ: \nID:%s\nユーザー名:%s\nメールアドレス:%s\nパスワード:%s",
			validUUID,
			moreLowerLimitLenUsername,
			validEmail,
			validLowerLengthPassword,
		),
	)

	name := NewUsername(moreLowerLimitLenUsername)
	if err := name.Valid(); err != nil {
		s.Fail(err.Error())
		return
	}

	account := NewUserAccount(s.validUUID, name, s.validEmailAddress, s.validPassword)
	s.Equal(
		UserAccount{
			id:       s.validUUID,
			username: moreLowerLimitLenUsername,
			email:    s.validEmailAddress,
			password: s.validPassword,
		},
		account,
	)
}
