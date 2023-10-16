package account

import (
	"net/mail"

	"github.com/google/uuid"
)

type UserAccountRepository interface {
	Find(uuid.UUID) (*UserAccount, error)
}

type UserAccountOutput struct {
	id       uuid.UUID
	name     Username
	email    mail.Address
	password Password
}

func NewUserAccountOutput(account UserAccount) UserAccountOutput {
	return UserAccountOutput{
		id:       account.id,
		name:     account.username,
		email:    account.email,
		password: account.password,
	}
}

type UserAccountService struct {
	userAccountRepository UserAccountRepository
}

func (s UserAccountService) Find(id uuid.UUID) (*UserAccountOutput, error) {
	account, err := s.userAccountRepository.Find(id)
	if err != nil {
		return nil, err
	}

	output := NewUserAccountOutput(*account)
	return &output, nil
}

func NewUserAccountService(repository UserAccountRepository) UserAccountService {
	return UserAccountService{
		userAccountRepository: repository,
	}
}
