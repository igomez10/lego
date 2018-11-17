package block

import "../transaction"

type block struct {
	transactions  []transaction.Transaction
	ID            string
	previousBlock *block
}

func NewBlock(transactions []transaction.Transaction) *block {

	// MerkleTreeImplementation
	//iteratingString := ""
	//for i, currentTransaction := range transactions {
	//
	//
	//}
	return &block{transactions: transactions}
}
