package commitment

import (
	"math/big"
)

func commit(v, r *big.Int) *big.Int {
	return new(big.Int).Mod(new(big.Int).Mul(new(big.Int).Exp(g, v, q), new(big.Int).Exp(h, r, q)), q)
}
