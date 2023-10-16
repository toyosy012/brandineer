package account

import (
	"errors"
	"fmt"
	"testing"

	"github.com/s-tajima/nspv"
	"github.com/stretchr/testify/suite"
)

const (
	invalidLowerLengthPassword = "Sx2MHrN"
	validLowerLengthPassword   = "Zj4GvQKy"
	validUpperLengthPassword   = "EEErp7FhWXiL4fYEsKJieBWFzW55WpKPkNmAkUT9PL6vmkzEwPwD25AtMVWXMkF3jTkjN3VL"
	invalidUpperLengthPassword = "3Z7WX4AgZY4z79HPnBdJtg2eYHL4Y9TvxScDMrvCFuFitiN679nWk8fq4zetexzPdVzLnkr4i"
	pwnedPassword              = "12345678"
)

var (
	errPwnedHibp              = errors.New("推測されやすいパスワード")
	errHashedComparedPassword = errors.New("比較対象データの生成に失敗")
	errHashedTestPassword     = errors.New("テストデータの生成に失敗")
	errFailHashed             = errors.New("テストデータのハッシュ値計算に失敗")
)

func TestPassword(t *testing.T) {
	suite.Run(t, newPasswordTestSuite())
}

type PasswordTestSuite struct {
	suite.Suite

	validator Validator
}

type NSPVValidator struct {
	validate *nspv.Validator
}

func NewNSPVValidator() NSPVValidator {
	v := nspv.NewValidator()
	v.SetHibpThreshold(1)
	return NSPVValidator{
		validate: v,
	}
}
func (v NSPVValidator) Valid(value Password) error {
	r, err := v.validate.Validate(value.Value())
	if err != nil {
		return err
	}

	if r == nspv.ViolateHibpCheck {
		return errPwnedHibp
	}
	return nil
}

func newPasswordTestSuite() *PasswordTestSuite {
	return &PasswordTestSuite{
		validator: NewNSPVValidator(),
	}
}

func (s *PasswordTestSuite) Validator() Validator {
	return s.validator
}

func (s *PasswordTestSuite) TestMoreLowerLimitLen() {
	s.T().Log(
		fmt.Sprintf("\nテストケース: パスワード下限境界値正常系テスト\nテストデータ: %s\nデータ長: %d",
			validLowerLengthPassword,
			len(validLowerLengthPassword),
		),
	)

	password, err := NewPassword(s.Validator(), validLowerLengthPassword)
	if err != nil {
		s.Fail(errFailHashed.Error())
		return
	}

	if err = password.Valid(); err != nil {
		s.Fail(err.Error())
		return
	}

	s.Equal(validLowerLengthPassword, password.value)
}

func (s *PasswordTestSuite) TestLessUpperLimitLen() {
	s.T().Log(
		fmt.Sprintf("\nテストケース: パスワード上限境界値正常系テスト\nテストデータ: %s\nデータ長: %d",
			validUpperLengthPassword,
			len(validUpperLengthPassword),
		),
	)

	password, err := NewPassword(s.Validator(), validUpperLengthPassword)
	if err != nil {
		s.Fail(errFailHashed.Error())
		return
	}

	if err = password.Valid(); err != nil {
		s.Fail(err.Error())
		return
	}

	s.Equal(validUpperLengthPassword, password.Value())
}

func (s *PasswordTestSuite) TestLessLowerLimitLen() {
	s.T().Log(
		fmt.Sprintf("\nテストケース: パスワード下限境界値異常系テスト\nテストデータ: %s\nデータ長: %d",
			invalidLowerLengthPassword,
			len(invalidLowerLengthPassword),
		),
	)

	password, err := NewPassword(s.Validator(), invalidLowerLengthPassword)
	if err != nil {
		s.Fail(errFailHashed.Error())
		return
	}
	s.ErrorIs(errMinLen, password.Valid())
}

func (s *PasswordTestSuite) TestMoreUpperLimitLen() {
	s.T().Log(
		fmt.Sprintf("\nテストケース: パスワード上限境界値異常系テスト\nテストデータ: %s\nデータ長: %d",
			invalidUpperLengthPassword,
			len(invalidUpperLengthPassword),
		),
	)

	_, err := NewPassword(s.Validator(), invalidUpperLengthPassword)
	s.ErrorIs(errMaxLen, err)
}

func (s *PasswordTestSuite) TestValid() {
	s.T().Log(
		fmt.Sprintf("\nテストケース: パスワードバリデーション正常テスト\nテストデータ: %s\n",
			validLowerLengthPassword,
		),
	)

	password, err := NewPassword(s.Validator(), validLowerLengthPassword)
	if err != nil {
		s.Fail(errFailHashed.Error())
		return
	}
	s.Equal(validLowerLengthPassword, password.Value())
}

func (s *PasswordTestSuite) TestErrorImpl() {
	s.T().Log(
		fmt.Sprintf("\nテストケース: 注入された実装から返されるエラーを返すかをテスト\nテストデータ: %s\n",
			pwnedPassword,
		),
	)

	password, err := NewPassword(s.Validator(), pwnedPassword)
	if err != nil {
		s.Fail(errFailHashed.Error())
		return
	}
	s.ErrorIs(errPwnedHibp, password.Valid())
}

func (s *PasswordTestSuite) TestCompareHash() {
	s.T().Log(
		fmt.Sprintf("\nテストケース: パスワードハッシュ値の比較正常テスト\nテストデータ: %s\n",
			validLowerLengthPassword,
		),
	)

	password, err := NewPassword(s.validator, validLowerLengthPassword)
	if err != nil {
		s.Fail(errHashedComparedPassword.Error())
	}
	test, err := NewPassword(s.validator, validLowerLengthPassword)
	if err != nil {
		s.Fail(errHashedTestPassword.Error())
	}
	s.NoError(password.CompareHash(*test))
}

func (s *PasswordTestSuite) TestErrCompareHash() {
	s.T().Log(
		fmt.Sprintf("\nテストケース: パスワードハッシュ値の比較異常テスト\nテストデータ: %s\n",
			validLowerLengthPassword,
		),
	)

	password, err := NewPassword(s.validator, validLowerLengthPassword)
	if err != nil {
		s.Fail(errHashedComparedPassword.Error())
	}
	test, err := NewPassword(s.validator, validUpperLengthPassword)
	if err != nil {
		s.Fail(errHashedTestPassword.Error())
	}
	s.ErrorIs(errCompareHash, password.CompareHash(*test))
}
