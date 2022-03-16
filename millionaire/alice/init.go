package alice

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

var Pk *rsa.PublicKey
var sk *rsa.PrivateKey

func init() {
	var err error
	sk, err = rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		fmt.Println(err)
		return
	}
	Pk = &sk.PublicKey
}
