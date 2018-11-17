package transaction

import (
	"crypto/sha256"
	"strconv"
	"time"
)
import "lego/account"

type Address account.Account

type Transaction struct {
	Origin    *account.Account
	Target    *account.Account
	Amount    int64
	Timestamp time.Time
}

func (t *Transaction) CalculateHash() string {

	var transactionContent string = t.Origin.Address.N.String() +
		t.Target.Address.N.String() +
		strconv.Itoa(int(t.Amount)) +
		t.Timestamp.String()

	hasher := sha256.New()
	hasher.Write([]byte(transactionContent))
	finalString := string(hasher.Sum(nil))
	return finalString
}
