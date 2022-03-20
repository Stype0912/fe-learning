package main

import (
	"fmt"
	"github.com/Stype0912/fe-learning/util"
	"math/big"
)

func DeduplicationUser(v *big.Int) *big.Int {
	fmt.Println(new(big.Int).Mod(v, p))
	bSum := big.NewInt(0)
	for i, _ := range b {
		bSum = new(big.Int).Add(bSum, new(big.Int).Mul(lambda[i], b[i]))
	}
	return new(big.Int).Add(bSum, v)
}

func DeduplicationCommittee(v_hat *big.Int) {
	var v []*big.Int
	for i := int64(0); i < N; i++ {
		tmp1 := new(big.Int).Mul(big.NewInt(N), lambda[i])
		tmp2 := util.Inverse(new(big.Int).Mod(tmp1, p), p)
		tmp3 := new(big.Int).Mul(v_hat, tmp2)
		v = append(v, new(big.Int).Sub(tmp3, b[i]))
	}
	vSum := big.NewInt(0)
	for i, _ := range v {
		vSum = new(big.Int).Add(vSum, new(big.Int).Mul(lambda[i], v[i]))
		vSum = new(big.Int).Mod(vSum, p)
	}
	fmt.Println(vSum)
}
