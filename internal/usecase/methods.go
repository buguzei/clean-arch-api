package usecase

import "fmt"

func (uc UseCase) SignUp() error {
	err := uc.db.InsertUser()
	if err != nil {
		return err
	}

	fmt.Println("from usecase")

	return nil
}

func (uc UseCase) SignIn() error {
	return nil
}

func (uc UseCase) DeleteAccount() error {
	return nil
}

func (uc UseCase) FillUpWallet() error {
	return nil
}

func (uc UseCase) WithDrawWallet() error {
	return nil
}
