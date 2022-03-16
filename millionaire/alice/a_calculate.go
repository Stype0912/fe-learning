package alice

import (
	"fmt"
	"github.com/fe-learning/millionaire/constant"
	my_rsa "github.com/fe-learning/rsa"
	"math/big"
)

func ACalculateZ(m, i *big.Int) (z []*big.Int) {
	z = append(z, big.NewInt(0))
	counter := big.NewInt(0)
	for {
		fmt.Println(counter)
		if counter.Cmp(constant.Limit) >= 0 {
			break
		}

		yu := my_rsa.Decrypt(sk, new(big.Int).Add(m, counter).Bytes())
		if counter.Cmp(i) <= 0 {
			z = append(z, new(big.Int).Mod(yu, constant.P))
		} else {
			z = append(z, new(big.Int).Add(new(big.Int).Mod(yu, constant.P), big.NewInt(1)))
		}
		counter = new(big.Int).Add(counter, big.NewInt(1))
	}
	return
}
