package millionaire

import (
	"github.com/Stype0912/fe-learning/millionaire/alice"
	"github.com/Stype0912/fe-learning/millionaire/bob"
	"math/big"
	"testing"
)

func TestMillionaire(t *testing.T) {
	//a := "123123"
	//b := hex.EncodeToString([]byte(a))
	//c, err := hex.DecodeString(b)
	//t.Log(string(c), err)

	i, j := big.NewInt(2), big.NewInt(4)
	pk := alice.Pk
	m := bob.BCalculateM(j, pk)
	z := alice.ACalculateZ(m, i)
	//t.Log(bob.BCheck(j, z))
	if !bob.BCheck(j, z) {
		t.Log("i >= j")
	} else {
		t.Log("i < j")
	}
}
