package Token

import "errors"

type Balances map[string]int

var GlbBalances Balances

// ApplyTransaction will update the Balances with the transaction if possible
func (b Balances) ApplyTransaction(t Transaction) error{

	if err := b.CanTransactionBeApplied(t); err != nil {
		return err
	}

	b[t.From] -= t.Amount
	b[t.To] += t.Amount
	return nil
}


func (b Balances) CanTransactionBeApplied(t Transaction) error {
	v, ok := b[t.From]
	if !ok {
		return errors.New("account non existent")
	}

	if v<t.Amount {
		return errors.New("funds are not sufficient")
	}
	return nil
}