package main

func main() {
    blockchain := NewBlockChain()
    defer blockchain.db.Close()

    cli := CLI{blockchain}
    cli.Run()
}
