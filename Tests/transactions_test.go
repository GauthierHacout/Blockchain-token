package Tests

import (
	"Blockchain-token/Token"
	"reflect"
	"testing"
)

func TestJsonToTransaction(t *testing.T) {
	//GIVEN
	input := `{"from":"Alice","to":"B","amount":3}`
	expected := Token.Transaction{From:"Alice", To:"B", Amount:3}
	//WHEN
	result, _ := Token.JsonToTransaction(input)
	//THEN
	if !reflect.DeepEqual(result,expected) {
		t.Errorf("Unexpected transaction from given string (%v), expected : %v, got : %v", input, expected, result)
	} else {
		t.Log("Correct transaction from given string")
	}
}

func TestString(t *testing.T) {
	//GIVEN
	transa := Token.Transaction{
		From:   "A",
		To:     "B",
		Amount: 3,
	}
	//WHEN
	result := transa.String()
	//THEN
	if result != "AB3" {
		t.Errorf("Incorrect result of Transaction.String() should be : %v, got : %v", "AB3", result)
	} else {
		t.Log("Correct stringify of Transaction")
	}
}