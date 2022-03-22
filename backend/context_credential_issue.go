package main

import (
	"crypto/rsa"
	"encoding/json"
	"github.com/Stype0912/fe-learning/threshold_signature"
	"github.com/ing-bank/zkrp/bulletproofs"
	"math/big"
)

func (c *Credential) ContextCredentialIssue(pkNewU *rsa.PublicKey, credMaster *Credential, proof bulletproofs.ProofBPRP) {
	// Encode the proof to JSON
	jsonEncoded, err := json.Marshal(proof)
	if err != nil {
		return
	}

	// Decode the proof from JSON
	var decodedProof bulletproofs.ProofBPRP
	err = json.Unmarshal(jsonEncoded, &decodedProof)
	if err != nil {
		return
	}

	// Verify the proof
	ok, errVerify := decodedProof.Verify()
	if errVerify != nil || !ok {
		return
	}

	m := &M{
		PkU:     pkNewU,
		Context: "age over 18",
	}
	m.CredentialSubject.Check = "dedupOver"

	mByte, _ := json.Marshal(m)
	//mStr := string(mByte)
	sigma := threshold_signature.Sign(new(big.Int).SetBytes(mByte))
	signature := threshold_signature.Combine(new(big.Int).SetBytes(mByte), sigma)
	c.Content = m
	c.Signature = signature.String()
	return
}
