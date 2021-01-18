package Network

import (
	"Blockchain-token/Blockchain"
	"Blockchain-token/Token"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/libp2p/go-libp2p-core/network"
	"os"
	"strings"
	"time"
)
// handleStream will set the behaviour a peer has when another peer connect to it
func handleStream(s network.Stream) {
	logger.Info("Got a new stream")
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))
	go readData(rw)
	go writeData(rw)
}

// readData is an infinite loop so we can always read new Blocks & Blockchain sent by other peers
func readData(rw *bufio.ReadWriter) {
	for {
		str, err := rw.ReadString('\n')
		if err != nil {
			logger.Warn(err)
		}

		if str == "" {
			return
		}

		if str != "\n" {
			newChain := make(Blockchain.Blockchain, 0)
			if err := json.Unmarshal([]byte(str), &newChain); err != nil {
				logger.Warn(err)
			}

			Blockchain.Mutex.Lock()
			Blockchain.Blckch = Blockchain.ReplaceChain(Blockchain.Blckch, newChain)
			logger.Infof("Blockchain : %v", Blockchain.Blckch)
			Blockchain.Mutex.Unlock()
		}
	}
}

// writeData is an infinite loop so we can always write new transactions,
// add them to our Blockchain & send it to other peers
func writeData(rw *bufio.ReadWriter) {

	// Broadcast the state of our Blockchain to peers every 5s
	go func() {
		for {
			time.Sleep(time.Second * 5)
			Blockchain.Mutex.Lock()
			bytes, err := json.Marshal(Blockchain.Blckch)
			if err != nil {
				logger.Warn(err)
			}

			_, _ = rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))
			_ = rw.Flush()
			Blockchain.Mutex.Unlock()
		}
	}()

	// Read input from user to add new Block inside our Blockchain
	stdReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := stdReader.ReadString('\n')
		if err != nil {
			logger.Warn(err)
			continue
		}
		cleanInput := strings.Replace(input, "\n", "", -1)
		transaction, err := Token.JsonToTransaction(cleanInput)
		if err != nil {
			logger.Warn(err)
			continue
		}

		// Check if new transaction if applicable
		if err = Token.GlbBalances.ApplyTransaction(transaction); err != nil {
			logger.Warn(err)
			continue
		}

		// Create a new Block with the input (Transaction) & add it to the Blockchain
		lastBlock := Blockchain.Blckch[len(Blockchain.Blckch)-1]
		newBlock := Blockchain.GenerateBlock(lastBlock, transaction)
		if Blockchain.IsBlockValid(newBlock, lastBlock) {
			Blockchain.Mutex.Lock()
			Blockchain.Blckch = append(Blockchain.Blckch, newBlock)
			Blockchain.Mutex.Unlock()
		}

		// Output the updated Blockchain
		bytes, err := json.Marshal(Blockchain.Blckch)
		if err != nil {
			logger.Warn(err)
		}
		spew.Dump(Blockchain.Blckch)

		Blockchain.Mutex.Lock()
		_, _ = rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))
		_ = rw.Flush()
		Blockchain.Mutex.Unlock()
	}
}