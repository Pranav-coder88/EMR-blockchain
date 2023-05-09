package main

import (
	"fmt"
	"model"
)

func printChain(bc *model.Blockchain) {
	for index, block := range bc.Chain.Links() {
		fmt.Println("\nindex:", index)
		fmt.Printf("Hashcode: %x\n", block.Hash)
		fmt.Println("block created on", block.Timestamp)
		fmt.Println("record created on", block.Transaction.Record.PersonalData.CreatedDate)
	}
}

func main() {
	bc := model.NewBlockchain(2)
	wallet := model.NewWallet()

	bc.AddTransactions([]model.Transaction{ *model.TransactionFromFile("model/dummy.json") })
	
	bc.ProcessPendingTransactions(wallet.GetWalletID())

	printChain(bc)
}
