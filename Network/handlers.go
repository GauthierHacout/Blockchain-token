package Network

import (
	"Blockchain-token/Blockchain"
	"Blockchain-token/Token"
	"bufio"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

// handleInput will read input from the user and try to turn it into a Transaction
func handleInput(stdReader *bufio.Reader) (Token.Transaction, error) {
	fmt.Print(">>>>>>>>> ")
	input, err := stdReader.ReadString('\n')
	if err != nil {
		return Token.Transaction{}, err
	}
	cleanInput := strings.Replace(input, "\n", "", -1)
	transaction, err := Token.JsonToTransaction(cleanInput)
	if err != nil {
		return Token.Transaction{}, err
	}

	return transaction, nil
}

// handleTransaction will try to create a new Block with the given transaction if applicable
func handleTransaction(t Token.Transaction) error{
	// Check if the transaction can be applied to current balances
	if err := Token.GlbBalances.CanTransactionBeApplied(t); err != nil {
		return err
	}

	// Create a new Block with the input (Transaction) & add it to the Blockchain
	lastBlock := Blockchain.Blckch[len(Blockchain.Blckch)-1]
	newBlock := Blockchain.GenerateBlock(lastBlock, t)
	lastBlock = Blockchain.Blckch[len(Blockchain.Blckch)-1] // In case of update of the Blockchain while computing
	if Blockchain.IsBlockValid(newBlock, lastBlock) {
		Blockchain.Mutex.Lock()
		Blockchain.Blckch = append(Blockchain.Blckch, newBlock)
		Blockchain.Mutex.Unlock()
	}

	return nil
}

// handleReceivedString will try to transform the data received from stream into either a Transaction or a Blockchain
func handleReceivedString(str string) (interface{}, error){
	transactionRegex := "^\\s?{.*\"from\":\"[[:alnum:]]*\".*\"to\":\"[[:alnum:]]*\".*}\\n$"
	matched, _ := regexp.MatchString(transactionRegex, str)
	if matched {
		newTransaction := Token.Transaction{}
		if err := json.Unmarshal([]byte(str), &newTransaction); err != nil {
			return nil, err
		}

		logger.Infof("New Transaction detected : %v", newTransaction)
		return newTransaction, nil
	} else {
		newChain := make(Blockchain.Blockchain, 0)
		if err := json.Unmarshal([]byte(str), &newChain); err != nil {
			return nil, err
		}

		return newChain, nil
	}
}