package threshold_signature

import (
	"math/big"
	"math/rand"
	"testing"
	"time"
)

// 节点数
var l int = 15

// 门限值
var k = 7

var a map[int]*big.Int

func TestSign(t *testing.T) {
	t.Log(p, q, p_hat, q_hat, m, n, e, d)
	//t.Log(new(big.Int).Exp(big.NewInt(0), big.NewInt(5), nil))
	a = make(map[int]*big.Int)
	a[0] = d
	for i := 1; i <= k-1; i++ {
		rand.Seed(time.Now().Unix())
		a[i] = new(big.Int).Rand(rand.New(rand.NewSource(time.Now().Unix())), new(big.Int).Sub(m, big.NewInt(1)))
	}

	s := make(map[int]*big.Int)
	for i := 0; i <= l; i++ {
		s[i] = f(i)
		t.Log(s[i])
	}
}

func f(x int) *big.Int {
	res := big.NewInt(0)
	for i := 0; i <= k-1; i++ {
		res = new(big.Int).Add(res, new(big.Int).Mod(new(big.Int).Mul(a[i], new(big.Int).Exp(big.NewInt(int64(x)), big.NewInt(int64(i)), m)), m))
	}
	return res
}
