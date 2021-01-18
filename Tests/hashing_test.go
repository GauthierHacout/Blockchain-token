package Tests

import (
	"Blockchain-token/Blockchain"
	"Blockchain-token/Token"
	"testing"
)

func TestCalculateHash(t *testing.T) {

	tables := []struct {
		input    Blockchain.Block
		expected string
	}{
		{Blockchain.Block{
			Index:        0,
			Timestamp:    "2021-01-13 17:06:33.7330141 +0100 CET m=+0.004031801",
			Transaction:      Token.Transaction{},
			PreviousHash: "03c28c828cc2b2558d975399118363de6dbf96a7ac82dfa53621c524319349a1",
			Nonce:        "34",
		},
			"bb1003dc92a0a64f5b68feabdc54e9dc479c601f697e12300084a5471d43d831"},
	}

	for _, table := range tables {
		// GIVEN (table.input)
		// WHEN (apply function to test)
		result := Blockchain.CalculateBlockHash(table.input)
		// THEN (check if result is the expected)
		if result != table.expected {
			t.Errorf("Incorrect hashing output got : %v, expected : %v", result, table.expected)
		} else {
			t.Logf("Correct hashing output %v", result)
		}
	}
}

func TestIsHashValid(t *testing.T) {

	tables := []struct {
		inputHash    	string
		inputDifficulty int
		expected 		bool
	}{
		{"0027698468396c34fr3", 2, true},
		{"027698468396c34fr3", 2, false},
		{"00027698468396c34fr3", 2, true},
	}

	for _, table := range tables {
		// GIVEN (table.input)
		// WHEN (apply function to test)
		result := Blockchain.IsHashValid(table.inputHash, table.inputDifficulty)
		// THEN (check if result is the expected)
		if result != table.expected {
			t.Errorf("Incorrect validity of hash (%v), got : %v, expected : %v", table.inputHash, result, table.expected)
		} else {
			t.Logf("Correct validity of hash (%v)", table.inputHash)
		}
	}
}


