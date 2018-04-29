# blockchain
learning blockchain

## Installation
```bash
go build
```

## Usage
```bash
> ./simple_blockchain
> Usage:
>   createblockchain -data BLOCK_DATA - add a block to the blockchain
>   printchain - print all the blocks of the blockchain
>   send -from FROM -to TO -amount AMOUNT - send AMOUNT of coins from FROM to TO
>   getbalance -address ADDRESS - get the balance of ADDRESS
```

## some minor issues
### part 3
1. `NewBlockChain()` doesn't have to take in address as parameter
2. it is worth notifying that when iterate through txs, the order is from back to front
3. `FindUnspentTransactions()` logic can be optimized for the sake of time complexity\
