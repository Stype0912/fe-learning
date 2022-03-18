package main

import (
	crand "crypto/rand"
	"crypto/rsa"
	"fmt"
	"github.com/Stype0912/fe-learning/commitment"
	"github.com/Stype0912/fe-learning/threshold_signature"
	"math/big"
	"math/rand"
	"time"
)

type Claim struct {
	a  string
	Cv *big.Int
}

var Pk *rsa.PublicKey
var sk *rsa.PrivateKey

func init() {
	var err error
	sk, err = rsa.GenerateKey(crand.Reader, 1024)
	if err != nil {
		fmt.Println(err)
		return
	}
	Pk = &sk.PublicKey
}

func PreCredentialGen() {
	a := "id"
	v, _ := new(big.Int).SetString("320282200009128411", 10)
	r := new(big.Int).Rand(rand.New(rand.NewSource(time.Now().Unix())), commitment.Q)
	Cv := commitment.Commit(v, r)
	claim := &Claim{
		a:  a,
		Cv: Cv,
	}

	PC := struct {
		claim    *Claim
		pkU      *rsa.PublicKey
		piOracle *big.Int
	}{
		claim:    claim,
		pkU:      Pk,
		piOracle: threshold_signature.Combine(claim.Cv, threshold_signature.Sign(claim.Cv)),
	}
	fmt.Printf("%v", PC.piOracle)
}
