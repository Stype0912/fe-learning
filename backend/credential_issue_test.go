package main

import (
	"encoding/json"
	"testing"
)

func TestMaster(t *testing.T) {
	PC := PreCredentialGen()
	credMaster := CredentialIssue(PC)
	credStr, _ := json.Marshal(credMaster)
	t.Log(string(credStr))
	//t.Log(threshold_signature.Verify())
}
