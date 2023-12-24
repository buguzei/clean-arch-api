package entities

type Wallet struct {
	ID       int    `json:"id"`
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}
