package tests

import (
	"fmt"
	"lego/account"
	"lego/transaction"
	"testing"
	"time"
)

func TestTransaction_CalculateHash(t *testing.T) {

	originAccount := account.NewAccount("aloha")
	targetAccount := account.NewAccount("hello")
	otherTransaction := transaction.Transaction{}
	testingTransaction := transaction.Transaction{
		Origin:    originAccount,
		Target:    targetAccount,
		Amount:    1,
		Timestamp: time.Unix(1541964066, 0),
	}
	fmt.Print(otherTransaction)
	fmt.Println(testingTransaction.CalculateHash())

}
