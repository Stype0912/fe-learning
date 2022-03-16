package bob

import (
	"math/big"
	"math/rand"
	"time"
)

var X *big.Int

func init() {
	X = new(big.Int).Rand(rand.New(rand.NewSource(time.Now().Unix())), new(big.Int).Exp(big.NewInt(2), big.NewInt(50), nil))
}
