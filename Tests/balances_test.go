package Tests

import (
	"Blockchain-token/Token"
	"testing"
)

func TestApplyTransaction(t *testing.T) {
	//GIVEN
	var balances Token.Balances = map[string]int{"Alice": 30}
	tables := []struct{
		input Token.Transaction
		expected bool
	}{
		{Token.Transaction{From: "Alice", To: "Bob", Amount: 30}, false},
		{Token.Transaction{From: "Alice", To: "Bob", Amount: 40}, true},
		{Token.Transaction{From: "XFSR", To: "Bob", Amount: 30}, true},
	}

	for _, table := range tables{
		//WHEN
		err := balances.ApplyTransaction(table.input)
		//THEN
		result := err!=nil
		if result != table.expected {
			t.Errorf("Unexpected error handling for transaction(%v), got : %v", table.input, err)
		} else {
			t.Log("Correct error handling for transaction")
		}
	}
}
