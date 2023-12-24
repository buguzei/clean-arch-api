package usecase

import (
	"clean-arch-api/internal/repo"
)

type UCUser interface {
	SignUp() error
	SignIn() error
	DeleteAccount() error
}

type UCWallet interface {
	FillUpWallet() error
	WithDrawWallet() error
}

type UC interface {
	UCUser
	UCWallet
}

type UseCase struct {
	db repo.Repo
}

func NewUseCase(r repo.Repo) UseCase {
	return UseCase{
		db: r,
	}
}
