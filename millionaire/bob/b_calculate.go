package bob

import (
	"crypto/rsa"
	my_rsa "github.com/Stype0912/fe-learning/rsa"
	"math/big"
)

func BCalculateM(j *big.Int, pk *rsa.PublicKey) *big.Int {
	k := my_rsa.Encrypt(pk, X.Bytes())
	m := new(big.Int).Add(new(big.Int).Sub(k, j), big.NewInt(1))
	//fmt.Println("22:", m)
	return m
}
