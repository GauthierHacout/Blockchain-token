package main

import (
	"Blockchain-token/Blockchain"
	"Blockchain-token/Network"
	"Blockchain-token/Token"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"time"
)

func main() {

	// Loading the values wrote inside .env file into the os.Env
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// Parse CLI Flags
	listenFlag := flag.Int("l", 0, "Port to serve inside the network")
	flag.Parse()
	if *listenFlag == 0 { // Port to serve is mandatory so node creation is possible
		log.Fatal("Please provide a port to bind on with -l")
	}

	// Create & Start a new Blockchain
	Blckch := new(Blockchain.Blockchain)
	Blckch.Initiate()
	Blockchain.Blckch = *Blckch

	// Random Values for Initial Balances
	Token.GlbBalances = map[string]int{"Alice": 50}

	go func() {
		for {
			time.Sleep(time.Second*30)
			fmt.Println("\n\n\n\n-------MY BLOCKCHAIN-----------", Blockchain.Blckch)
			fmt.Println("-----------MY BALANCE------------", Token.GlbBalances)
			fmt.Println("-----END--------\n\n\n")
		}
	}()

	// Start the peer 2 peer network
	Network.Start(*listenFlag)
}