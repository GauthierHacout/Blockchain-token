package Tests

import (
	"Blockchain-token/Blockchain"
	"Blockchain-token/Token"
	"testing"
)

func TestIsBlockValid (t *testing.T) {

	tables := []struct {
		inputNewBlock   Blockchain.Block
		inputOldBlock   Blockchain.Block
		expected 		bool
	}{
		{Blockchain.Block{Hash: "953d10d3185bfc1659fa71fa9c7ccf8fc52a4ba9e819cfc6f67427dc09af577b", PreviousHash: "12345", Index: 1},
			Blockchain.Block{Hash: "12345", Index: 0} , true},

		{Blockchain.Block{Hash: "953d10d3185bfc1659fa71fa9c7ccf8fc52a4ba9e819cfc6f67427dc09af577b", PreviousHash: "12345", Index: 1},
			Blockchain.Block{Hash: "12345", Index: 1} , false},

		{Blockchain.Block{Hash: "953d10d3185bfc1659fa71fa9c7ccf8fc52a4ba9e819cfc6f67427dc09af577b", PreviousHash: "12345", Index: 1},
			Blockchain.Block{Hash: "23456", Index: 0} , false},

		{Blockchain.Block{Hash: "ABC4356", PreviousHash: "12345", Index: 1},
			Blockchain.Block{Hash: "12345", Index: 0} , false},
	}

	for _, table := range tables {
		// GIVEN (table.input)
		// WHEN (apply function to test)
		result := Blockchain.IsBlockValid(table.inputNewBlock, table.inputOldBlock)
		// THEN (check if result is the expected)
		if result != table.expected {
			t.Errorf("Incorrect output of IsBlockValid (%v) expected %v, got : %v", table.inputNewBlock, table.expected, result)
		} else {
			t.Logf("Correct output of IsBlockValid (%v) ", table.inputNewBlock)
		}
	}
}

func TestGenerateBlock (t *testing.T) {

	tables := []struct {
		inputTransa    	Token.Transaction
		inputOldBlock   Blockchain.Block
		expected 		Blockchain.Block
	}{
		{Token.Transaction{}, Blockchain.Block{Hash: "12345", Index: 0} ,
			Blockchain.Block{Index:1, Transaction: Token.Transaction{}, PreviousHash: "12345"}},
	}

	for _, table := range tables {
		// GIVEN (table.input)
		// WHEN (apply function to test)
		result := Blockchain.GenerateBlock(table.inputOldBlock, table.inputTransa)
		// THEN (check if result is the expected)
		if result.Transaction != table.expected.Transaction ||
			result.PreviousHash != table.expected.PreviousHash ||
			result.Index != table.expected.Index {
			t.Errorf("Incorrect generation of Block expected %v, got : %v", table.expected, result)
		} else {
			t.Log("Correct generation of Block ")
		}
	}
}

