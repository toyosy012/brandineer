package account

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestUserAccountService(t *testing.T) {
	s := NewUserAccountServiceTestSuite()
	suite.Run(t, &s)
}

type UserAccountServiceTest struct {
	suite.Suite
}

func NewUserAccountServiceTestSuite() UserAccountServiceTest { return UserAccountServiceTest{} }
