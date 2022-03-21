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
	A  string   `json:"a"`
	Cv *big.Int `json:"cv"`
}

type PreCredential struct {
	Claim    *Claim         `json:"Claim"`
	PkU      *rsa.PublicKey `json:"pk_u"`
	PiOracle *big.Int       `json:"pi_oracle"`
}

var Pk *rsa.PublicKey
var sk *rsa.PrivateKey
var V *big.Int

func init() {
	var err error
	sk, err = rsa.GenerateKey(crand.Reader, 1024)
	if err != nil {
		fmt.Println(err)
		return
	}
	Pk = &sk.PublicKey
}

func PreCredentialGen() PreCredential {
	a := "id"
	V, _ = new(big.Int).SetString("320282200009128411", 10)
	r := new(big.Int).Rand(rand.New(rand.NewSource(time.Now().Unix())), commitment.Q)
	Cv := commitment.Commit(V, r)
	claim := &Claim{
		A:  a,
		Cv: Cv,
	}

	PC := PreCredential{
		Claim:    claim,
		PkU:      Pk,
		PiOracle: threshold_signature.Combine(claim.Cv, threshold_signature.Sign(claim.Cv)),
	}
	return PC
	//fmt.Printf("%V", PC)
}
