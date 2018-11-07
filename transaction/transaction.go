package transaction

import "time"
import "../account"

type Address account.Account

type Transaction struct {
	origin    *account.Account
	target    *account.Account
	amount    int64
	timestamp time.Time
	Hash      string
}

func (t Transaction) ValidateTransaction() bool {
	a := new(account.Account)
	a.Initialize("hola")


	return true
}
