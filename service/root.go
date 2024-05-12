package service

import (
	"CRUD-SERVER/repository"
	"sync"
)

//Network, repoistory의 다리 역활

var (
	serviceInit     sync.Once
	serviceInstance *Service
)

type Service struct {
	//repository
	repository *repository.Repository
	User       *User
}

func NewService(rep *repository.Repository) *Service {
	serviceInit.Do(func() {
		serviceInstance = &Service{
			repository: rep,
		}
		serviceInstance.User = newUserService(rep.User)
	})

	return serviceInstance
}
