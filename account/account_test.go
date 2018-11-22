package account

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"fmt"
	"os"
	"testing"
)

func TestAccount_NewAccount(t *testing.T) {

	firstAccount := NewAccount()

	// EncryptOAEP
	msg := []byte("The secret message!")
	label := []byte("")
	sha1hash := sha1.New()
	encryptedmsg, err := rsa.EncryptOAEP(sha1hash, rand.Reader, &firstAccount.privateKey.PublicKey, msg, label)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("OAEP encrypted [%s] to \n[%x]\n", string(msg), encryptedmsg)
	fmt.Println()
	// DecryptOAEP
	decryptedmsg, err := rsa.DecryptOAEP(sha1hash, rand.Reader, firstAccount.privateKey, encryptedmsg, label)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if !bytes.Equal(msg, decryptedmsg) {
		t.Error("Encription/Decription with account prublic/private keys not working")
	}

}
