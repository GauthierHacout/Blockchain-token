package Network

import (
	"Blockchain-token/Blockchain"
	"Blockchain-token/Token"
	"bufio"
	"encoding/json"
	"fmt"
)

// broadcastBlockchain will send the state of our Blockchain to peers on the network
func broadcastBlockchain(rw *bufio.ReadWriter) {
	Blockchain.Mutex.Lock()
	bytes, err := json.Marshal(Blockchain.Blckch)
	if err != nil {
		logger.Warnf("Broadcast Blockchain : %v", err)
	}

	_, _ = rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))
	_ = rw.Flush()
	Blockchain.Mutex.Unlock()
}

// broadcastTransaction will send the given transaction to peers on the network
func broadcastTransaction(rw *bufio.ReadWriter, t Token.Transaction) {
	bytes, err := json.Marshal(t)
	if err != nil {
		logger.Warnf("Broadcast Transaction : %v", err)
	}

	_, _ = rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))
	_ = rw.Flush()
}
