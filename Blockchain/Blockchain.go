package Blockchain

import (
	"Blockchain-token/Token"
	"time"
)

type Blockchain []Block

var Blckch Blockchain

// ReplaceChain will check if the provided Blockchain is longer than the current one
// so it can be replaced. On replacements update the balances with the new transactions
func ReplaceChain(current Blockchain, new Blockchain) Blockchain{
	if len(new) > len(current) {
		for i := len(current)-1; i<len(new); i++{
			_ = Token.GlbBalances.ApplyTransaction(new[i].Transaction)
		}

		return new
	}
	return current
}

// Initiate will create a genesis Block inside the Blockchain to start it
func (Blckch *Blockchain)Initiate() {
	genesisBlock := Block{
		Index:        0,
		Timestamp:    time.Now().String(),
		Transaction:  Token.Transaction{},
		Hash:         CalculateBlockHash(Block{}),
		PreviousHash: "",
		Difficulty:   Difficulty,
		Nonce:        "",
	}

	Mutex.Lock()
	*Blckch = append(*Blckch, genesisBlock)
	Mutex.Unlock()
}