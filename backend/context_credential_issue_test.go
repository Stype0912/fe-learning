package main

import (
	crand "crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/ing-bank/zkrp/bulletproofs"
	"math/big"
	"testing"
)

func TestContext(t *testing.T) {
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
		t.Log(DeduplicationCommitteeMPC(rawData, i, vi))
	}

	credMaster := &Credential{}
	credMaster.MasterCredentialIssue(PC)

	params, errSetup := bulletproofs.SetupGeneric(18, 200)
	if errSetup != nil {
		t.Errorf(errSetup.Error())
		t.FailNow()
	}
	bigSecret := new(big.Int).SetInt64(int64(20))
	proof, _ := bulletproofs.ProveGeneric(bigSecret, params)

	skNew, err := rsa.GenerateKey(crand.Reader, 1024)
	if err != nil {
		fmt.Println(err)
		return
	}
	PkNew := &skNew.PublicKey
	credContext := Credential{}
	credContext.ContextCredentialIssue(PkNew, credMaster, proof)
	credStr, _ := json.Marshal(credContext)
	t.Log(string(credStr))
}
