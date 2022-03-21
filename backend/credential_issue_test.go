package main

import (
	"encoding/json"
	"github.com/Stype0912/fe-learning/threshold_signature"
	"math/big"
	"testing"
)

func TestMaster(t *testing.T) {
	PC := PreCredentialGen()
	t.Log(PC.PiOracle, PC.Claim.Cv)
	v_hat, rawData := DeduplicationUser(V)
	//rawData := func() []string {
	//	return mpc.CalculateHashedLeaves(v, -1)
	//}()
	t.Log(rawData)
	for i := 0; i <= 10; i++ {
		rawData[i] = ""
		vi := DeduplicationCommitteeVi(v_hat, int64(i))
		DeduplicationCommitteeMPC(rawData, i, vi)
	}

	credMaster := CredentialIssue(PC)
	credStr, _ := json.Marshal(credMaster)
	t.Log(string(credStr))
	mByte, _ := json.Marshal(credMaster.Content)
	X, _ := new(big.Int).SetString(credMaster.Signature, 10)
	t.Log(threshold_signature.Verify(new(big.Int).SetBytes(mByte), X))
}
