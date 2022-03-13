package util

import "math/big"

func Inverse(a, n *big.Int) *big.Int {
	d, x, _ := GcdExpand(a, n)
	if d.Cmp(big.NewInt(1)) == 1 {
		tmp1 := x.Mod(x, n)
		tmp2 := tmp1.Add(tmp1, n)
		return tmp2.Mod(tmp2, n)
	} else {
		return big.NewInt(-1)
	}
}
