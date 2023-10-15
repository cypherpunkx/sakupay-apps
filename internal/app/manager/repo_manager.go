package manager

import (
	"github.com/sakupay-apps/internal/app/repository"
)

type RepoManager interface {
	UserRepo() repository.UserRepository
	TransactionRepo() repository.TransactionRepository
	BillRepo() repository.BillRepository
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
