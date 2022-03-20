package main

import (
	"crypto/rsa"
	"encoding/json"
	"github.com/Stype0912/fe-learning/threshold_signature"
	"math/big"
)

type CredMaster struct {
	PkU               *rsa.PublicKey `json:"pk_u"`
	Context           string         `json:"context"`
	Claim             *Claim         `json:"claim"`
	CredentialSubject struct {
		Check     string `json:"check"`
		Attribute []struct {
			Value    string `json:"value"`
			Provider string `json:"provider"`
		} `json:"attribute"`
	} `json:"credential_subject"`
	Signature string `json:"signature"`
}

func CredentialIssue(pc PreCredential) *CredMaster {
	claim := pc.claim
	m := struct {
		pkU               *rsa.PublicKey
		context           string
		claim             *Claim
		credentialSubject struct {
			check     string
			attribute []struct {
				Value    string
				Provider string
			}
		}
	}{
		pkU:     Pk,
		context: "Master",
		claim:   claim,
		credentialSubject: struct {
			check     string
			attribute []struct {
				Value    string
				Provider string
			}
		}{check: "dedupOver", attribute: []struct {
			Value    string
			Provider string
		}{
			{
				Value:    claim.Cv.String(),
				Provider: "gov.cn",
			},
		}},
	}
	mByte, _ := json.Marshal(m)
	//mStr := string(mByte)
	sigma := threshold_signature.Sign(new(big.Int).SetBytes(mByte))
	signature := threshold_signature.Combine(new(big.Int).SetBytes(mByte), sigma)
	credMaster := &CredMaster{
		PkU:     Pk,
		Context: "Master",
		Claim:   claim,
		CredentialSubject: struct {
			Check     string `json:"check"`
			Attribute []struct {
				Value    string `json:"value"`
				Provider string `json:"provider"`
			} `json:"attribute"`
		}(struct {
			Check     string
			Attribute []struct {
				Value    string
				Provider string
			}
		}{Check: "dedupOver", Attribute: []struct {
			Value    string
			Provider string
		}{
			{
				Value:    claim.Cv.String(),
				Provider: "gov.cn",
			},
		}}),
		Signature: signature.String(),
	}
	return credMaster
}
