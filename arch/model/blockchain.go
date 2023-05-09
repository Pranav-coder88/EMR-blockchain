package model

import (
	"bytes"
	"errors"
)

// A Blockchain facilitates processing of transactions through consensus of blocks.
// @property Chain - This is the chain of blocks that make up the blockchain.
// @property {int} difficulty - The number of leading zeros that the hash of a block must have.
// @property {[]Transaction} transactionBuffer - This is a slice of transactions that are waiting to be
// added to the blockchain.
// @property {CurrencyType} reward - The amount of currency that is rewarded to the miner of a block.
type Blockchain struct {
	Chain             *ChainLinks
	difficulty        int
	transactionBuffer []Transaction
	reward            CurrencyType
}

// NewBlockchain creates a new blockchain with a genesis block and returns a pointer to it
func NewBlockchain(difficulty int) *Blockchain {
	blockchain := &Blockchain{
		Chain:             NewChain(),
		difficulty:        difficulty,
		transactionBuffer: make([]Transaction, 0),
		reward:            CurrencyType(1),
	}

	return blockchain
}

// Adding a block to the blockchain.
func (blockchain *Blockchain) addBlock(blockData []Transaction) {
	for _, block := range blockData {
		statement := NewBlock(block, blockchain.Chain.Latest.Hash)
		statement.Mine(blockchain.difficulty)
		blockchain.Chain.Add(statement)
	}
}

// AddTransactions Adds transactions to the blockchain.
func (blockchain *Blockchain) AddTransactions(transactionData []Transaction) error {
	if len(transactionData) == 100 {
		return errors.New("transactionBuffer is full. Try processing pending transactions")
	}

	blockchain.transactionBuffer = append(blockchain.transactionBuffer, transactionData...)

	return nil
}

// ProcessPendingTransactions  , Processes the pending transactions. (Currently only supports PoW)
func (blockchain *Blockchain) ProcessPendingTransactions(rewardAddress walletType) {
	blockchain.addBlock(blockchain.transactionBuffer)
	blockchain.transactionBuffer = blockchain.transactionBuffer[:0]
}

// IsChainValid Checking if the blockchain is valid.
func (blockchain *Blockchain) IsChainValid() bool {
	var prevHash []byte

	for _, block := range blockchain.Chain.Links() {
		if !block.IsValid() || !bytes.Equal(prevHash, block.PrevHash) {
			return false
		}

		prevHash = block.Hash
	}

	return true
}
