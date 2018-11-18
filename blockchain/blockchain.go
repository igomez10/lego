package main

import (
	"fmt"
	"lego/block"
)


func main() {
	genesisTransactions := []string{"Hello","hi"}
	genesisBlock := block.NewBlock( genesisTransactions, nil )
	fmt.Printf("%+v \n", genesisBlock.Hash)
	fmt.Println("FINISH")
}