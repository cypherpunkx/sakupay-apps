package manager

import (
	"github.com/sakupay-apps/internal/app/service"
)

type ServiceManager interface {
	UserService() service.UserService
	AuthService() service.AuthService
	BillService() service.BillService
}

type serviceManager struct {
	repoManager RepoManager
}

func NewServiceManager(repo RepoManager) ServiceManager {
	return &serviceManager{
		repoManager: repo,
	}
}

func (m *serviceManager) UserService() service.UserService {
	return service.NewUserService(m.repoManager.UserRepo())
}

func (m *serviceManager) AuthService() service.AuthService {
	return service.NewAuthService(m.UserService())
}

func (m *serviceManager) BillService() service.BillService {
	return service.NewBillService(m.repoManager.BillRepo(),m.repoManager.UserRepo())
}