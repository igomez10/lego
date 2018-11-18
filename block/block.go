package block

import (
	"crypto/sha512"
	"strings"
)

type Block struct {
	Transactions  []string
	Hash          string
	PreviousBlock *Block
}

func NewBlock(transactions []string, previousBlock *Block) *Block {

	currentBlock := Block{}
	currentBlock.PreviousBlock = previousBlock
	currentBlock.Transactions = transactions

	transactionsContent := strings.Join(currentBlock.Transactions, "")

	harsher := sha512.New()
	harsher.Write([]byte(transactionsContent))
	currentBlock.Hash = string(harsher.Sum([]byte{}))

	return &currentBlock
}

