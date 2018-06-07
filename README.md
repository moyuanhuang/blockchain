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

## Example
```bash
$ ./simple_blockchain createwallet
Your new wallet address: 1MfCxbN471GdK2FKfEShN1J1q6nT9kexdn

$ ./simple_blockchain createblockchain --address 1MfCxbN471GdK2FKfEShN1J1q6nT9kexdn
Mining a new block...
0000007e1610d899ae1e5b1600d904f392bebc35a47691b979f6907dfe28ef1c
2018/06/06 23:05:19 New block added. used 4.703125037s
Done!

$ ./simple_blockchain getbalance --address 1MfCxbN471GdK2FKfEShN1J1q6nT9kexdn
Balance of 1MfCxbN471GdK2FKfEShN1J1q6nT9kexdn: 10

$ ./simple_blockchain createwallet
Your new wallet address: 1CNSohgdWLyGY8ZTFp1ZvHuP6Yav8djmD

$ ./simple_blockchain send --from 1MfCxbN471GdK2FKfEShN1J1q6nT9kexdn --to 1CNSohgdWLyGY8ZTFp1ZvHuP6Yav8djmD --amount 1
Mining a new block...
000000ef2acce3828404b454c71e6faf80b248d1036c2d30f3127acb76fe6e75
2018/06/06 23:06:00 New block added. used 2.743028811s
Success!

$ ./simple_blockchain getbalance --address 1MfCxbN471GdK2FKfEShN1J1q6nT9kexdn
Balance of 1MfCxbN471GdK2FKfEShN1J1q6nT9kexdn: 9

$ ./simple_blockchain getbalance --address 1CNSohgdWLyGY8ZTFp1ZvHuP6Yav8djmD
Balance of 1CNSohgdWLyGY8ZTFp1ZvHuP6Yav8djmD: 1

$ ./simple_blockchain listaddresses
1MfCxbN471GdK2FKfEShN1J1q6nT9kexdn
1CNSohgdWLyGY8ZTFp1ZvHuP6Yav8djmD
```

## some minor issues
### part 3
1. `NewBlockChain()` doesn't have to take in address as parameter
2. it is worth notifying that when iterate through txs, the order is from back to front
3. `FindUnspentTransactions()` logic can be optimized for the sake of time complexity\
