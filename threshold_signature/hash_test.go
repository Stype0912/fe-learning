package threshold_signature

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"testing"
)

func TestHash(t *testing.T) {
	b := sha256.Sum256([]byte("123"))
	s := fmt.Sprintf("%x", b)
	x, err := new(big.Int).SetString(s, 16)
	t.Log(x, err)
}
