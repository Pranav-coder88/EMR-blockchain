package model

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"time"
)

type hashcode []byte
type noncecode []byte

// A Block is a collection of Transactions, a timestamp, a hash, a previous hash, and a nonce.
// @property {[]Transaction} Transactions - An array of transactions that are included in the block.
// @property Timestamp - The time at which the block was created.
// @property {hashcode} Hash - The hash of the block.
// @property {hashcode} PrevHash - The hash of the previous block in the chain.
// @property {noncecode} Nonce - A random number that is used to mine the block.
type Block struct {
	Transaction Transaction `json:"transactions"`
	Timestamp    time.Time     `json:"timestamp"`
	Hash         hashcode      `json:"hash"`
	PrevHash     hashcode      `json:"prev_hash"`
	Nonce        noncecode     `json:"nonce"`
}

// GenesisBlock It creates a first block, sets the timestamp, previous hash, and nonce to empty values, and then
// generates the hash for the block
func GenesisBlock() *Block {
	block := new(Block)
	block.Transaction = Transaction{}
	block.Timestamp = time.Now()
	block.PrevHash = hashcode{}
	block.Nonce = make(noncecode, 32)

	block.Hash = block.generateHash()
	return block
}

// NewBlock creates a new block with the given transactions and previous block's hash
func NewBlock(transaction Transaction, previousHash hashcode) *Block {
	block := new(Block)
	block.Transaction = transaction
	block.Timestamp = time.Now()
	block.PrevHash = previousHash
	block.Hash = block.generateHash()
	block.Nonce = make(noncecode, 32)
	return block
}

// Generating a hash for the block.
func (block *Block) generateHash() hashcode {
	input := append(block.Nonce, block.PrevHash...)
	input = append(input, block.Timestamp.String()...)

	input = append(input, []byte(block.Transaction.Writable())...)

	hash := sha256.Sum256(input)
	return hash[:]
}

// If the first `depth` bytes of `solution` are all zero, then return true, otherwise return false
func isSolved(solution []byte, depth int) bool {
	for index, solutionByte := range solution {

		if index >= depth {
			break
		}

		if solutionByte != 0 {
			return false
		}
	}

	return true
}

// Mine , Generates a hashcode for the given difficulty
func (block *Block) Mine(difficulty int) {
	// This is the mining process. It is generating a random number and assigning it to the nonce when the hash remains unsolved.
	for !isSolved(block.Hash, difficulty) {
		_, err := rand.Read(block.Nonce)
		if err != nil {
			panic(err)
		}

		block.Hash = block.generateHash()
	}
}

func (block *Block) Print() {
	fmt.Printf("\ttime: %v\n", block.Timestamp.String())
	fmt.Printf("\tprevious hash: %x\n", block.PrevHash)
	fmt.Printf("\thash: %x\n", block.Hash)

	fmt.Printf("\ttransaction: %v\n", block.Transaction.Writable())
}

// IsValid Checks for altered contents within the block's transactions
func (block *Block) IsValid() bool {
	return bytes.Equal(block.Hash, block.generateHash())
}