package constant

import (
	"math/big"
	"math/rand"
	"time"
)

var Limit, P *big.Int

func init() {
	Limit, _ = new(big.Int).SetString("10", 10)
	P = new(big.Int).Rand(rand.New(rand.NewSource(time.Now().Unix())), new(big.Int).Exp(big.NewInt(2), big.NewInt(20), nil))
}
