package usecase

import "toh-echo/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (i *UserInteractor) Add(u domain.User) {
	i.UserRepository.Store(u)
}

func (i *UserInteractor) GetInfo() []domain.User {
	return i.UserRepository.Select()
}

func (i *UserInteractor) Delete(id string) {
	i.UserRepository.Delete(id)
}
