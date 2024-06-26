package repository

import (
	"sync"
)

// 세션정의

var (
	repositoryInit     sync.Once
	repositoryInstance *Repository
)

type Repository struct {
	User *UserRepository
}

func NewRepository() *Repository {
	repositoryInit.Do(func() {
		repositoryInstance = &Repository{
			User: newUserRepository(),
		}
	})

	return repositoryInstance
}
