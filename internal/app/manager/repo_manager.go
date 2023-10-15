package manager

import (
	"github.com/sakupay-apps/internal/app/repository"
)

type RepoManager interface {
	UserRepo() repository.UserRepository
	TransactionRepo() repository.TransactionRepository
}

type repoManager struct {
	infraManager InfraManager
}

func (m *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(m.infraManager.Conn())
}

func (m *repoManager) TransactionRepo() repository.TransactionRepository {
	return repository.NewTransactionRepository(m.infraManager.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infraManager: infra,
	}
}
