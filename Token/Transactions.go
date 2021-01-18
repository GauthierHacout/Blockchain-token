package Token

import (
	"encoding/json"
	"strconv"
)

type Transaction struct {
	From 	string 	`json:"from"` // Address of the giver
	To		string 	`json:"to"` // Address of the receiver
	Amount	int 	`json:"amount"` // Amount of token exchanged
}

// JsonToTransaction will try to transform a string in json format into a Transaction
// (e.g `{"from":"A","to":"B","amount":3}` => Transaction{From:"A", To:"B", Amount:3})
func JsonToTransaction(str string) (Transaction, error) {
	trans := Transaction{}
	if err := json.Unmarshal([]byte(str), &trans) ; err != nil {
		return trans, err
	}

	return trans, nil
}

// String => Transaction.toString()
func (t *Transaction) String() string {
	return t.From + t.To + strconv.Itoa(t.Amount)
}