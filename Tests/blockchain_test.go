package Tests

import (
	"Blockchain-token/Blockchain"
	"Blockchain-token/Token"
	"reflect"
	"testing"
	"time"
)

func TestReplaceChain(t *testing.T) {

	tables := []struct {
		inputCurrent   	Blockchain.Blockchain
		inputNew 		Blockchain.Blockchain
		expected 		Blockchain.Blockchain
	}{
		{
			inputCurrent: []Blockchain.Block{{Index: 0}, {Index: 1}},
			inputNew: []Blockchain.Block{{Index: 0}},
			expected: []Blockchain.Block{{Index: 0}, {Index: 1}},
		},
		{
			inputCurrent: []Blockchain.Block{{Index: 0}, {Index: 1}},
			inputNew: []Blockchain.Block{{Index: 0}, {Index: 1}, {Index: 2}},
			expected: []Blockchain.Block{{Index: 0}, {Index: 1}, {Index: 2}},
		},
	}

	for _, table := range tables {
		// GIVEN (table.input)
		// WHEN (apply function to test)
		result := Blockchain.ReplaceChain(table.inputCurrent, table.inputNew)
		// THEN (check if result is the expected)
		if !reflect.DeepEqual(result, table.expected) {
			t.Errorf("Incorrect current should be %v, is : %v", table.expected, result)
		} else {
			t.Logf("Correct replacement of Blockchain")
		}
	}
}

func TestInitiate(t *testing.T) {
	//GIVEN
	blckch := Blockchain.Blockchain{}
	genesisBlock := Blockchain.Block{
		Index:        0,
		Timestamp:    time.Now().String(),
		Transaction:  Token.Transaction{},
		Hash:         Blockchain.CalculateBlockHash(Blockchain.Block{}),
		PreviousHash: "",
		Difficulty:   Blockchain.Difficulty,
		Nonce:        "",
	}

	//WHEN
	blckch.Initiate()
	//THEN
	if len(blckch) == 0 ||
		blckch[0].Index != genesisBlock.Index ||
		blckch[0].Transaction != genesisBlock.Transaction ||
		blckch[0].Hash != genesisBlock.Hash ||
		blckch[0].Difficulty != genesisBlock.Difficulty ||
		blckch[0].Nonce != genesisBlock.Nonce {
		t.Errorf("Incorrect initiation of Blockchain : %v", blckch)
	} else {
		t.Log("Correct initiation of Blockchain")
	}
}
