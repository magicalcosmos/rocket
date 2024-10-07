package infrastructure

import (
	"errors"
	"rocket/domain/account"
)

type InMemoryAccountRepository struct {
	accounts map[string]*account.Account
}

func NewInMemoryAccountRepository() *InMemoryAccountRepository {
	return &InMemoryAccountRepository{
		accounts: make(map[string]*account.Account),
	}
}

func (repo *InMemoryAccountRepository) Save(a *account.Account) error {
	repo.accounts[a.ID] = a
	return nil
}

func (repo *InMemoryAccountRepository) FindById(id string) (*account.Account, error) {
	if account, exists := repo.accounts[id]; exists {
		return account, nil
	}
	return nil, errors.New("account not found")
}

func (repo *InMemoryAccountRepository) DeleteById(id string) error {
	delete(repo.accounts, id)
	return nil
}
