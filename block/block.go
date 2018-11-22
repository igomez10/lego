package block

import (
	"crypto/sha512"
)

// Transaction contains the information about a transaction int eh blockchain
// TODO inmplement actual transaction
type Transaction string

// Transaction contains information about a transaction in the blockchain
type mempool []Transaction

// Block is the basic unit in the blockchain, it contains a number of transactions
// and its corresponding hash, previous block
type Block struct {
	Transactions  []Transaction // a transaction is a string in v.alpha
	Hash          string
	PreviousBlock *Block
}

// NewBlock creates a new block with the set of given transactions and the
// given previous block
func NewBlock(transactionPool *mempool, previousBlock *Block) *Block {

	currentBlock := Block{}
	currentBlock.PreviousBlock = previousBlock

	// First, select which transactions the block will contain
	selectedTransactions := selectTransactions(transactionPool)
	currentBlock.Transactions = selectedTransactions

	// Second, calculate the hash of the selected transactions
	hashedTransaction := string(processTransactions(selectedTransactions))
	hashedBlockData := hashedTransaction + currentBlock.Hash
	currentBlock.Hash = hashedBlockData
	return &currentBlock
}

// selectTransactions chooses the transactions from the mempool that
// will be added into the blockchain, expected to be  judged by the transaction fee
func selectTransactions(transactionPool *mempool) []Transaction {
	return (*transactionPool)[:10]
}

// resolveTransaction hashes a specific transaction
// and returns the new value in the form of bytes
func resolveTransaction(transaction Transaction) []byte {
	transactionData := []byte(transaction)
	return transactionData
}

// processTransactions resolves a new block from the given mempool
// it is in charge of selecting which trasactions to process
func processTransactions(selectedTransactions []Transaction) []byte {
	harsher := sha512.New()
	for _, v := range selectedTransactions {
		transformedTransaction := resolveTransaction(v)
		harsher.Write(transformedTransaction)
	}
	return harsher.Sum([]byte{})
}
