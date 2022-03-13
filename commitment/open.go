package commitment

import (
	"math/big"
)

func open(v, r, commitment *big.Int) bool {
	if new(big.Int).Mod(new(big.Int).Mul(new(big.Int).Exp(g, v, q), new(big.Int).Exp(h, r, q)), q).Cmp(commitment) == 0 {
		return true
	} else {
		return false
	}
}
