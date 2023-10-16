package account

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"testing"
)

const (
	validServiceUUID = "f311a818-5d48-4603-8543-368976dcca6b"
)

type UserAccountRepositoryStub struct{}

func NewUserAccountRepositoryStub() UserAccountRepositoryStub {
	return UserAccountRepositoryStub{}
}
func (s UserAccountRepositoryStub) Find(_ uuid.UUID) (*UserAccount, error) {
	return &s.validUserAccount, nil
}

func TestUserAccountService(t *testing.T) {
	s := NewUserAccountServiceTestSuite()
	suite.Run(t, &s)
}

type UserAccountServiceTest struct {
	suite.Suite

	validationID       uuid.UUID
	userAccountService UserAccountService
}

func NewUserAccountServiceTestSuite() UserAccountServiceTest { return UserAccountServiceTest{} }

func (s *UserAccountServiceTest) SetupSuite() {
	id, err := uuid.Parse(validServiceUUID)
	if err != nil {
		s.Fail(err.Error())
		return
	}
	s.id = id
	dbStub := NewUserAccountRepositoryStub()
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
