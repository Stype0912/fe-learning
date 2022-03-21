package bob

import (
	"github.com/Stype0912/fe-learning/millionaire/constant"
	"math/big"
)

func BCheck(j *big.Int, z []*big.Int) bool {
	//fmt.Println(z)
	//fmt.Println(new(big.Int).Mod(X, constant.P))
	if z[j.Int64()].Cmp(new(big.Int).Mod(X, constant.P)) == 0 {
		return false
	}
	return true
}
