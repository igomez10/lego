package account

import (
	"crypto/rsa"
	"strings"
)

// Account holds data regarding the public-private key and amount in tokens
type Account struct {
	Address  rsa.PublicKey
	privateKey rsa.PrivateKey
	Amount     int64
}


// Initialize is the method used to create an account
func (newAccount Account) Initialize(passphrase string) (*Account, error){

	newReader := strings.NewReader(passphrase)
	privKey, err := rsa.GenerateKey(newReader,4096)
	if err != nil {
		return nil, err
	}

	newAccount.privateKey = *privKey
	newAccount.Address = newAccount.privateKey.PublicKey
	newAccount.Amount = 0

	return &newAccount, nil
}



