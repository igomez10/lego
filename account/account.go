package account

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"log"

	"golang.org/x/crypto/ssh"
)

// Account holds data regarding the public-private key and amount in tokens
type Account struct {
	Address    string
	privateKey *rsa.PrivateKey
	Amount     int64
}

// NewAccount creates a new account and returns a pointer to the
// new account. The optional string is intented to add a customization
// level. based on that string, genrate a new account.
func NewAccount() *Account {

	createdAccount := Account{}

	newPrivateKey, err := generatePrivateKey(4096)
	if err != nil {
		log.Fatal("Error generating rsa private key ", err)
	}
	newPublicKey, err := generatePublicKey(&newPrivateKey.PublicKey)
	if err != nil {
		log.Fatal("Error generating rsa public key ", err)
	}
	createdAccount.privateKey = newPrivateKey
	createdAccount.Address = string(newPublicKey)
	createdAccount.Amount = 0

	return &createdAccount
}

// generatePrivateKey creates a RSA Private Key of specified byte size
func generatePrivateKey(bitSize int) (*rsa.PrivateKey, error) {
	// Private Key generation
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	// Validate Private Key
	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// generatePublicKey take a rsa.PublicKey and return bytes suitable for writing to .pub file
// returns in the format "ssh-rsa ..."
func generatePublicKey(privatekey *rsa.PublicKey) ([]byte, error) {
	publicRsaKey, err := ssh.NewPublicKey(privatekey)
	if err != nil {
		return nil, err
	}

	pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)
	return pubKeyBytes, nil

}

// Sign signs the given transaction with the publicKey
func (ac *Account) Sign(transaction []byte) []byte {
	label := []byte("")
	sha1hash := sha1.New()
	encryptedmsg, err := rsa.EncryptOAEP(sha1hash, rand.Reader, ac.privateKey.Public().(*rsa.PublicKey), transaction, label)
	if err != nil {
		log.Println(err)
	}
	return encryptedmsg
}
