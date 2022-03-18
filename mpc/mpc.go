package mpc

import (
	"crypto/sha256"
	"fmt"
	"math/big"
)

type Node struct {
	hash  string
	left  *Node
	right *Node
}

func merkleTree(hashedLeaves []string) string {
	lowerLayer := hashedLeaves
	higherLayer := make([]string, 0)
	for {
		for k := 0; k < len(lowerLayer)/2; k++ {
			higherLayer = append(higherLayer, fmt.Sprintf("%x", sha256.Sum256([]byte(lowerLayer[2*k]+lowerLayer[2*k+1]))))
		}
		if len(lowerLayer)%2 == 1 {
			higherLayer = append(higherLayer, lowerLayer[len(lowerLayer)-1])
		}
		if len(higherLayer) == 1 {
			break
		}
		lowerLayer = higherLayer
		higherLayer = make([]string, 0)
	}
	return higherLayer[0]
}

func Mpc(data []*big.Int, i int) string {
	var leaves = data
	hashedLeaves := make([]string, 0)
	for j, leaf := range leaves {
		if j == i {
			hashedLeaves = append(hashedLeaves, fmt.Sprintf("%x", leaf.Bytes()))
		} else {
			hashedLeaf := sha256.Sum256(leaf.Bytes())
			hashedLeaves = append(hashedLeaves, fmt.Sprintf("%x", hashedLeaf))
		}
	}
	return merkleTree(hashedLeaves)
}
