package main

import (
	"crypto/rsa"
	"encoding/json"
	"github.com/Stype0912/fe-learning/threshold_signature"
	"math/big"
)

type CredMaster struct {
	Content   *M     `json:"content"`
	Signature string `json:"signature"`
}

type M struct {
	PkU               *rsa.PublicKey `json:"pk_u"`
	Context           string         `json:"context"`
	Claim             *Claim         `json:"claim"`
	CredentialSubject struct {
		Check     string     `json:"check"`
		Attribute *Attribute `json:"attribute"`
	} `json:"credential_subject"`
}

type Attribute []struct {
	Value    string `json:"value"`
	Provider string `json:"provider"`
}

func CredentialIssue(pc PreCredential) *CredMaster {
	claim := pc.Claim
	m := &M{
		PkU:     Pk,
		Context: "Master",
		Claim:   claim,
	}
	m.CredentialSubject.Check = "dedupOver"
	m.CredentialSubject.Attribute = &Attribute{
		{
			Value:    claim.Cv.String(),
			Provider: "gov.cn",
		},
	}
	mByte, _ := json.Marshal(m)
	//mStr := string(mByte)
	sigma := threshold_signature.Sign(new(big.Int).SetBytes(mByte))
	signature := threshold_signature.Combine(new(big.Int).SetBytes(mByte), sigma)
	credMaster := &CredMaster{
		Content:   m,
		Signature: signature.String(),
	}
	return credMaster
}
