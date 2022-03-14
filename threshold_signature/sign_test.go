package threshold_signature

import (
	"crypto/sha256"
	"fmt"
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
	rand.Seed(time.Now().Unix())
	//t.Log(rand.Intn(2))
	//t.Log(new(big.Int).Exp(big.NewInt(0), big.NewInt(5), nil))
	a = make(map[int]*big.Int)
	a[0] = d
	for i := 1; i <= k-1; i++ {
		rand.Seed(time.Now().Unix())
		a[i] = new(big.Int).Rand(rand.New(rand.NewSource(time.Now().Unix())), new(big.Int).Sub(m, big.NewInt(1)))
	}

	s := make(map[int]*big.Int)
	for i := 0; i <= l; i++ {
		s[i] = new(big.Int).Mod(f(i), m)
		t.Logf("s[%v]: %v", i, s[i])
	}

	sqrtN := new(big.Int).Sub(new(big.Int).Sqrt(n), big.NewInt(5))
	v := new(big.Int).Mul(sqrtN, sqrtN)
	VK := make(map[int]*big.Int)
	for i := 1; i <= l; i++ {
		VK[i] = new(big.Int).Exp(v, s[i], n)
	}

	Delta := factorial(big.NewInt(int64(l)))

	counter := 0

	var lTemp []int
	for i := 0; i <= l; i++ {
		lTemp = append(lTemp, i)
	}
	var S, SOut []int
	rand.Shuffle(len(lTemp), func(i, j int) {
		lTemp[i], lTemp[j] = lTemp[j], lTemp[i]
	})
	for i := 0; i <= l; i++ {
		if counter < k {
			counter++
			S = append(S, lTemp[i])
		} else {
			SOut = append(SOut, lTemp[i])
		}
	}
	//t.Log(S, "\n", SOut)

	lambda := make(map[int]map[int]*big.Int)
	for _, i := range SOut {
		subLambda := make(map[int]*big.Int)
		for _, j := range S {
			fenzi := func() *big.Int {
				res := big.NewInt(1)
				for _, j_hat := range S {
					if j_hat == j {
						continue
					}
					res = new(big.Int).Mul(res, big.NewInt(int64(i-j_hat)))
				}
				return res
			}()
			fenmu := func() *big.Int {
				res := big.NewInt(1)
				for _, j_hat := range S {
					if j_hat == j {
						continue
					}
					res = new(big.Int).Mul(res, big.NewInt(int64(j-j_hat)))
				}
				return res
			}()
			subLambda[j] = new(big.Int).Div(new(big.Int).Mul(Delta, fenzi), fenmu)
			lambda[i] = subLambda
			//t.Logf("lambda[%v][%v]: %v", i, j, lambda[i][j])
		}
	}

	test1 := new(big.Int).Mul(Delta, f(SOut[1]))
	test2 := func() *big.Int {
		res := big.NewInt(0)
		for _, j := range S {
			res = new(big.Int).Add(res, new(big.Int).Mul(lambda[SOut[1]][j], f(j)))
		}
		return res
	}()
	t.Log(new(big.Int).Mod(test1, m))
	t.Log(new(big.Int).Mod(test2, m))

	M := "320282200009128411"
	x, _ := new(big.Int).SetString(fmt.Sprintf("%x", sha256.Sum256([]byte(M))), 16)
	t.Log(x)

	X := make(map[int]*big.Int)
	for i := 1; i <= l; i++ {
		X[i] = new(big.Int).Exp(x, new(big.Int).Mul(new(big.Int).Mul(big.NewInt(2), Delta), s[i]), n)
		//t.Log(X[i])
	}

	x_wave := new(big.Int).Exp(x, new(big.Int).Mul(big.NewInt(4), Delta), n)
	Ln := n.BitLen()
	L1 := 128
	r := new(big.Int).Rand(rand.New(rand.NewSource(time.Now().Unix())), new(big.Int).Exp(big.NewInt(2), new(big.Int).Add(big.NewInt(int64(Ln)), big.NewInt(int64(2*L1))), n))
	v_hat := new(big.Int).Exp(v, r, n)
	x_hat := new(big.Int).Exp(x_wave, r, n)
	for i := 1; i <= l; i++ {
		vStr := v.String()
		xWaveStr := x_wave.String()
		viStr := VK[i].String()
		xi2 := new(big.Int).Mul(X[i], X[i]).String()
		vHatStr := v_hat.String()
		xHatStr := x_hat.String()
		c, _ := new(big.Int).SetString(fmt.Sprintf("%x", sha256.Sum256([]byte(vStr+xWaveStr+viStr+xi2+vHatStr+xHatStr))), 16)
		z := new(big.Int).Add(new(big.Int).Mul(s[i], c), r)
		vzvic := new(big.Int).Mod(new(big.Int).Mul(new(big.Int).Exp(v, z, n), new(big.Int).Exp(VK[i], new(big.Int).Mul(big.NewInt(-1), c), n)), n).String()
		xzxi2c := new(big.Int).Mod(new(big.Int).Mul(new(big.Int).Exp(x_wave, z, n), new(big.Int).Exp(X[i], new(big.Int).Mul(big.NewInt(-2), c), n)), n).String()
		cOther, _ := new(big.Int).SetString(fmt.Sprintf("%x", sha256.Sum256([]byte(vStr+xWaveStr+viStr+xi2+vzvic+xzxi2c))), 16)
		if c.Cmp(cOther) != 0 {
			t.Log(c, cOther)
			t.FailNow()
		}
	}

}

func f(x int) *big.Int {
	res := big.NewInt(0)
	for i := 0; i <= k-1; i++ {
		res = new(big.Int).Add(res, new(big.Int).Mul(a[i], new(big.Int).Exp(big.NewInt(int64(x)), big.NewInt(int64(i)), nil)))
	}
	return res
}

func factorial(n *big.Int) *big.Int {
	if n.Cmp(big.NewInt(0)) == 1 {
		return new(big.Int).Mul(n, factorial(new(big.Int).Sub(n, big.NewInt(1))))
	}
	return big.NewInt(1)
}
