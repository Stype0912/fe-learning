package util

import "math/big"

func IsPrime(n *big.Int, num int) bool {
	return n.ProbablyPrime(num)
}
