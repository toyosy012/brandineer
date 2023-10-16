package account

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"net/mail"
	"testing"
)

const (
	validServiceUsername = "user1"
	validServiceEmail    = "user1@example.com"
	validServiceUUID     = "f311a818-5d48-4603-8543-368976dcca6b"
	validServicePassword = "Zj4GvQKy"
)

type UserAccountRepositoryStub struct {
	validUserAccount  UserAccount
	validUserAccounts []UserAccount
}

func NewUserAccountRepositoryStub(account UserAccount) UserAccountRepositoryStub {
	return UserAccountRepositoryStub{
		validUserAccount:  account,
		validUserAccounts: []UserAccount{account},
	}
}
func (s UserAccountRepositoryStub) Find(_ uuid.UUID) (*UserAccount, error) {
	return &s.validUserAccount, nil
}

func (s UserAccountRepositoryStub) List() ([]UserAccount, error) {
	return s.validUserAccounts, nil
}

func TestUserAccountService(t *testing.T) {
	s := NewUserAccountServiceTestSuite()
	suite.Run(t, &s)
}

type UserAccountServiceTest struct {
	suite.Suite

	validationID       uuid.UUID
	userAccountService UserAccountService
	validUserAccount   UserAccount
	validUserAccounts  []UserAccount
}

func NewUserAccountServiceTestSuite() UserAccountServiceTest { return UserAccountServiceTest{} }

func (s *UserAccountServiceTest) SetupSuite() {
	id, err := uuid.Parse(validServiceUUID)
	if err != nil {
		s.Fail(err.Error())
		return
	}
	s.validationID = id

	address, err := mail.ParseAddress(validServiceEmail)
	if err != nil {
		s.Fail(err.Error())
		return
	}

	validator := NewNSPVValidator()
	password, err := NewPassword(validator, validServicePassword)

	account := NewUserAccount(id, NewUsername(validServiceUsername), *address, *password)
	s.validUserAccount = account
	s.validUserAccounts = []UserAccount{account}
	dbStub := NewUserAccountRepositoryStub(account)
	s.userAccountService = NewUserAccountService(dbStub)
}

func (s *UserAccountServiceTest) TestFindUserAccountOutput() {
	account, err := s.userAccountService.Find(s.validationID)
	if err != nil {
		s.Fail(err.Error())
		return
	}

	s.Equal(
		NewUserAccountOutput(s.validUserAccount),
		*account,
	)
}

func (s *UserAccountServiceTest) TestFindUserAccountsOutput() {
	accounts, err := s.userAccountService.List()
	if err != nil {
		s.Fail(err.Error())
		return
	}

	var outputs []UserAccountOutput
	for _, a := range s.validUserAccounts {
		outputs = append(outputs, NewUserAccountOutput(a))
	}

	s.Equal(
		outputs,
		accounts,
	)
}
