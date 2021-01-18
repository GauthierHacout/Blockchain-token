package Blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
)

// CalculateBlockHash will create a Hash of the information in the given Block concatenated in a string.
// (Using SHA256 algorithm)
func CalculateBlockHash(block Block) string {
	record := strconv.Itoa(block.Index) +
		block.Timestamp +
		block.Transaction.String() +
		block.PreviousHash +
		block.Nonce

	hash := sha256.New()
	hash.Write([]byte(record))
	hashed := hash.Sum(nil)

	return hex.EncodeToString(hashed)
}

// IsHashValid will check if a Hash as the right number of leading 0s, depending on difficulty
// (e.g if difficulty=3 hash should start by "000")
func IsHashValid(hash string, difficulty int) bool{
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}

