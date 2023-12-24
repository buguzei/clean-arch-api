package repo

type UserRepo interface {
	InsertUser() error
	DeleteUser() error
	FindUser() error
}

type WalletRepo interface {
	IncreaseWallet() error
	DecreaseWallet() error
}

type Repo interface {
	UserRepo
	WalletRepo
}
