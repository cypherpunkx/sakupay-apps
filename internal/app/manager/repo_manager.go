package manager

import (
	"github.com/sakupay-apps/internal/app/repository"
)

type RepoManager interface {
	UserRepo() repository.UserRepository
	BillRepo() repository.BillRepository
}

type repoManager struct {
	infraManager InfraManager
}

func (m *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(m.infraManager.Conn())
}

func (m *repoManager) BillRepo() repository.BillRepository {
	return repository.NewBillRepository(m.infraManager.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infraManager: infra,
	}
}
