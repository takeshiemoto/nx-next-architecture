package usecase

import "toh-echo/domain"

type UserRepository interface {
	Store(user domain.User)
	Select() []domain.User
	Delete(id string)
}
