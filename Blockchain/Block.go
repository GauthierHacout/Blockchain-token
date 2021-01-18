package Blockchain

import (
	"Blockchain-token/Token"
	"strconv"
"sync"
"time"
)

const Difficulty = 1

var Mutex = &sync.Mutex{}

// A Block is an element of the Blockchain
type Block struct {
	Index 			int 	// Position of the Block in the Blockchain
	Timestamp 		string	// Exact Timestamp of when the Block was created
	Transaction 	Token.Transaction 	// The transaction validated, the essence of the Block
	Hash 			string	// Hash of the Block using SHA256 algorithm
	PreviousHash	string 	// Hash of the previous Block in the Blockchain
	Difficulty 		int		// The number of leading 0s the Hash of this Block should have
	Nonce			string	// Value assigned randomly for the Hash to be of the right difficulty
}

// IsBlockValid will check for 3 conditions of validity :
// Index, PreviousHash & CurrentHash
func IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index ||
		oldBlock.Hash != newBlock.PreviousHash ||
		CalculateBlockHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

// GenerateBlock will create a new Block using a custom Message, the current time
// & the last Block in the Blockchain
// params are meant to be used for Block.Message (1st) & Block.Validator (2nd) if a Validator is needed
func GenerateBlock(oldBlock Block, transaction Token.Transaction) Block {
	newBlock := Block{
		Index:        oldBlock.Index+1,
		Timestamp:    time.Now().String(),
		Transaction:  transaction,
		PreviousHash: oldBlock.Hash,
		Difficulty:   Difficulty,
	}

	// While loop until the Hash of the Block is valid
	for i:=0; ; i++ {
		newBlock.Nonce = strconv.Itoa(i) // We modify Nonce to change the calculated Hash
		calculatedHash := CalculateBlockHash(newBlock)
		if !IsHashValid(calculatedHash, newBlock.Difficulty) {
			continue
		} else {
			newBlock.Hash = calculatedHash
			break
		}
	}

	return newBlock
}

