package manager

import (
	"github.com/sakupay-apps/config"
	"github.com/sakupay-apps/internal/app/repository"
)

type RepoManager interface {
	UserRepo() repository.UserRepository
	TransactionRepo() repository.TransactionRepository
	BillRepo() repository.BillRepository
	ContactRepo() repository.ContactRepository
	FileRepo() repository.FileRepository
	UserPictureRepo() repository.UserPictureRepository
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

func (m *repoManager) FileRepo() repository.FileRepository {
	return repository.NewFileRepository(config.Cfg.File.UserPicturePath)
}

func (m *repoManager) UserPictureRepo() repository.UserPictureRepository {
	return repository.NewUserPictureRepository(m.infraManager.Conn())
}
