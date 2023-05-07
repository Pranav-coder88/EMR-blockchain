package main

import (
	"fmt"
	"model"
)

func main() {
	bc := model.NewBlockchain(2)

	bc.AddTransactions([]model.Transaction{ *model.TransactionFromFile("model/dummy.json") })
	
	if bc.IsChainValid() {
		fmt.Println("valid") 
	} else {
		fmt.Println("invalid")
	}
}
