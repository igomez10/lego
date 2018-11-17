package account

import (
	"log"
	"testing"
)

// import "lego/account"

func TestAccount_NewAccount(t *testing.T) {

	firstAccount := NewAccount()

	log.Println("Address:", firstAccount.Address)

}
