package millionaire

import (
	"github.com/fe-learning/millionaire/alice"
	"github.com/fe-learning/millionaire/bob"
	"math/big"
	"testing"
)

func TestMillionaire(t *testing.T) {
	//a := "123123"
	//b := hex.EncodeToString([]byte(a))
	//c, err := hex.DecodeString(b)
	//t.Log(string(c), err)

	i, j := big.NewInt(6), big.NewInt(4)
	pk := alice.Pk
	m := bob.BCalculateM(j, pk)
	z := alice.ACalculateZ(m, i)
	t.Log(bob.BCheck(j, z))
}
