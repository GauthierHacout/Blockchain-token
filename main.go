package main

import (
	"Blockchain-token/Blockchain"
	"Blockchain-token/Network"
	"flag"
	"github.com/joho/godotenv"
	"log"
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

	// Start the peer 2 peer network
	Network.Start(*listenFlag)
}