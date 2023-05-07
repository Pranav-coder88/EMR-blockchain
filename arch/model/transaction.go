package model

import (
	"encoding/json"
)

type Transaction struct {
	Record EMR
}

func (transaction *Transaction) Read(filepath string) {
	transaction.Record = *JSONToEMR(filepath)
}

func TransactionFromFile(filepath string) *Transaction {
	return &Transaction{ Record: *JSONToEMR(filepath) }
}

func (transaction *Transaction) Writable() string {
	var onPaper string
	json.Unmarshal(transaction.Record.Bytable(), &onPaper)
	return onPaper
}