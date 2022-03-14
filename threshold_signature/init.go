package threshold_signature

import (
	"math/big"
)

var p, q, p_hat, q_hat, n, m, e, d *big.Int

func init() {
	//file, err := os.Open("big_prime1.txt")
	//if err != nil {
	//	fmt.Printf("%v", err)
	//	return
	//}
	//defer file.Close()
	//content, err := ioutil.ReadAll(file)
	//big_primes := strings.Split(string(content), ",")
	//
	//for _, prime := range big_primes {
	//	p_hat, _ = new(big.Int).SetString(prime, 10)
	//	p = new(big.Int).Add(new(big.Int).Mul(big.NewInt(2), p_hat), big.NewInt(1))
	//	if !p.ProbablyPrime(10) {
	//		continue
	//	} else {
	//		break
	//	}
	//}
	//fmt.Printf("p:%v,q:%v,p':%v,q':%v,m:%v,n:%v", p, q, p_hat, q_hat, m, n)

	p, _ = new(big.Int).SetString("195035409467", 10)
	p_hat, _ = new(big.Int).SetString("97517704733", 10)
	q, _ = new(big.Int).SetString("195288414899", 10)
	q_hat, _ = new(big.Int).SetString("97644207449", 10)
	n, _ = new(big.Int).SetString("38088155963987848448833", 10)
	m, _ = new(big.Int).SetString("9522038990899381156117", 10)
	e = big.NewInt(17)
	d = new(big.Int).ModInverse(e, m)
	//fmt.Println(p.ProbablyPrime(10), q.ProbablyPrime(10), p_hat.ProbablyPrime(10), q_hat.ProbablyPrime(10))
	//n = new(big.Int).Mul(p, q)
	//m = new(big.Int).Mul(p_hat, q_hat)
}
