package manager

import (
	"github.com/sakupay-apps/internal/app/repository"
)

type RepoManager interface {
	UserRepo() repository.UserRepository
	TransactionRepo() repository.TransactionRepository
	BillRepo() repository.BillRepository
	ContactRepo() repository.ContactRepository
	CardRepo() repository.CardRepository
}

type repoManager struct {
	infraManager InfraManager
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infraManager: infra,
	}
}

func (m *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(m.infraManager.Conn())
}

func (m *repoManager) TransactionRepo() repository.TransactionRepository {
	return repository.NewTransactionRepository(m.infraManager.Conn())
}

func (m *repoManager) BillRepo() repository.BillRepository {
	return repository.NewBillRepository(m.infraManager.Conn())
}

func (m *repoManager) ContactRepo() repository.ContactRepository {
	return repository.NewContactRepository(m.infraManager.Conn())
}

func (m *repoManager) CardRepo() repository.CardRepository {
	return repository.NewCardRepository(m.infraManager.Conn())
}
