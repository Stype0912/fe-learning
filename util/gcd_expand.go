package util

import "math/big"

func GcdExpand(a, b *big.Int) (*big.Int, *big.Int, *big.Int) {
	if b.Cmp(big.NewInt(0)) == 0 {
		return a, big.NewInt(1), big.NewInt(0)
	} else {
		tmp := new(big.Int)
		d, x1, y1 := GcdExpand(b, tmp.Mod(a, b))
		tmp1 := tmp.Div(a, b)
		tmp2 := tmp1.Mul(tmp1, y1)
		x0, y0 := y1, x1.Sub(x1, tmp2)
		return d, x0, y0
	}
}
