package postgres

import "fmt"

func (p Postgres) InsertUser() error {
	fmt.Println("from postgres method")

	return nil
}

func (p Postgres) FindUser() error {
	return nil
}

func (p Postgres) DeleteUser() error {
	return nil
}

func (p Postgres) IncreaseWallet() error {
	return nil
}

func (p Postgres) DecreaseWallet() error {
	return nil
}
