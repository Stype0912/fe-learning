package main

import (
	"math/big"
	"math/rand"
)

var N = int64(15)

var lambda, b []*big.Int

var p *big.Int

func init() {
	p, _ = new(big.Int).SetString("114466057660826548352085136953911344157318943320700600451512433047942256725079", 10)
	//for {
	//	r := rand.New(rand.NewSource(time.Now().Unix()))
	//	p = new(big.Int).Rand(r, new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil))
	//	if !p.ProbablyPrime(10) {
	//		continue
	//	}
	//	break
	//}
	for i := int64(0); i < N; i++ {
		//r := time.Now().UnixNano()
		r := int64(1)
		lambda = append(lambda, new(big.Int).Rand(rand.New(rand.NewSource(r)), new(big.Int).Exp(big.NewInt(2), big.NewInt(512), nil)))
		b = append(b, new(big.Int).Rand(rand.New(rand.NewSource(r)), new(big.Int).Exp(big.NewInt(2), big.NewInt(512), nil)))
	}
}
