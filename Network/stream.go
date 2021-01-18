package Network

import (
	"Blockchain-token/Blockchain"
	"Blockchain-token/Token"
	"bufio"
	"github.com/libp2p/go-libp2p-core/network"
	"os"
	"time"
)


// handleStream will set the behaviour a peer has when another peer connect to it
func handleStream(s network.Stream) {
	logger.Info("Got a new stream")
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))
	go readData(rw)
	go writeData(rw)
}

// readData is an infinite loop so we can always read new Transactions, Blocks & Blockchain sent by other peers
func readData(rw *bufio.ReadWriter) {
	for {
		str, err := rw.ReadString('\n')
		if err != nil {
			logger.Warnf("Read String from I/O : %v", err)
		}

		if str == "" {
			return
		}

		if str != "\n" {
			data, err := handleReceivedString(str)
			if err != nil {
				logger.Warnf("Read handle Received String : %v",err)
				continue
			}

			switch v := data.(type) {
			case Blockchain.Blockchain:
				replaced := Blockchain.ReplaceChain(Blockchain.Blckch, v)
				if replaced {
					logger.Infof("Updated Blockchain : %v", Blockchain.Blckch)
				}
			case Token.Transaction:
				go func() {
					if err := handleTransaction(v); err != nil {
						logger.Warnf("Read handleTransaction : %v", err)
					} else {
						broadcastBlockchain(rw)
					}
				}()
			default:
				logger.Warnf("Unable to identify the data received from stream")
			}

		}
	}
}

// writeData is an infinite loop so we can always write new transactions,
// add them to our Blockchain & send it to other peers
func writeData(rw *bufio.ReadWriter) {

	go func() {
		for {
			time.Sleep(time.Second*30)
			broadcastBlockchain(rw)
		}
	}()

	stdReader := bufio.NewReader(os.Stdin)
	for {
		transaction, err := handleInput(stdReader)
		if err != nil {
			logger.Warnf("Write handleInput : %v", err)
			continue
		}

		logger.Infof("New transaction written : %v",transaction)
		broadcastTransaction(rw, transaction)

		go func() {
			if err := handleTransaction(transaction); err != nil {
				logger.Warnf("Write handleTransaction : %v", err)
			} else {
				broadcastBlockchain(rw)
			}
		}()
	}
}
