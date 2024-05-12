package repository

import (
	"CRUD-SERVER/types"
	"CRUD-SERVER/types/errors"
)

type UserRepository struct {
	userMap []*types.User
}

func newUserRepository() *UserRepository {
	return &UserRepository{
		userMap: []*types.User{},
	}
}

func (u *UserRepository) Create(newUser *types.User) error {
	u.userMap = append(u.userMap, newUser)
	return nil
}

func (u *UserRepository) Update(name string, age int) error {

	var isExisted bool
	for _, user := range u.userMap {
		if user.Name == name {
			user.Age = age
			isExisted = true
			continue
		}
	}

	if !isExisted {
		return errors.ErrorOf(errors.NotFoundUser, nil)
	}

	return nil
}

func (u *UserRepository) Delete(userName string) error {
	var isExisted bool
	for idx, user := range u.userMap {
		if user.Name == userName {
			u.userMap = append(u.userMap[:idx], u.userMap[idx+1:]...)

			isExisted = true
			continue
		}
	}

	if !isExisted {
		return errors.ErrorOf(errors.NotFoundUser, nil)
	}

	return nil
}

func (u *UserRepository) Get() []*types.User {
	return u.userMap
}
