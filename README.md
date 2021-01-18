# Blockchain-token

### How it works

This project is a Blockchain implementation in Golang. It is similar to implementation used in cryptocurrencies

The Blockchain act as a distributed decentralized data storage for token transactions.

Features :
- Peer 2 Peer Network, with peer discovery
- Proof of Work consensus Algorithm
- SHA256 Hashing
- Addition of new transactions

In order to test this you can ```go build -o <name>``` in the project root directory

And then using multiples terminal you can : 
- ``./<name> -l 8000``
- ``./<name> -l 8002``
- ``./<name> -l 8004``

This will create multiples nodes (3 to be exact) serving different local port (8000, 8002, 8004). The nodes will connect to each other & start building & sharing the blockchain with new transactions you input.

To input a new transaction you can input a string (json format) in any terminal : 
```{"from": "Alice", "to": "Bob", "amount": 30}```

When you input a new transaction, every node in the network start creating a new Block & the first node to find the correct Hash for the Block will broadcast his updated Blockchain to the network. Effectively updating the Blockchain everywhere.

### Tests

Most of the code is tested ! => 
```go test -v ./Tests``` in the root directory
